package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/public-awesome/stargaze/x/staking/types"
	"github.com/spf13/cobra"
)

type OsmosisSnapshot struct {
	TotalOsmoAmount         sdk.Int                           `json:"total_osmo_amount"`
	TotalStarsAirdropAmount sdk.Int                           `json:"total_stars_amount"`
	NumberAccounts          uint64                            `json:"num_accounts"`
	Accounts                map[string]OsmosisSnapshotAccount `json:"accounts"`
	StargazeDelegators      map[string]sdk.Int                `json:"stargaze_delegators"`
	LiquidityProviderCount  uint64                            `json:"lp_count"`
}

// OsmosisSnapshotAccount provide fields of snapshot per account
type OsmosisSnapshotAccount struct {
	OsmoAddress string `json:"osmo_address"` // Atom Balance = AtomStakedBalance + AtomUnstakedBalance

	OsmoBalance          sdk.Int `json:"osmo_balance"`
	OsmoOwnershipPercent sdk.Dec `json:"osmo_ownership_percent"`

	OsmoStakedBalance   sdk.Int `json:"osmo_staked_balance"`
	OsmoUnstakedBalance sdk.Int `json:"osmo_unstaked_balance"` // AtomStakedPercent = AtomStakedBalance / AtomBalance
	OsmoStakedPercent   sdk.Dec `json:"osmo_staked_percent"`

	// StarsBalance = sqrt( AtomBalance ) * (1 + 1.5 * atom staked percent)
	StarsBalance sdk.Int `json:"stars_balance"`
	// StarsBalanceBase = sqrt(atom balance)
	StarsBalanceBase sdk.Int `json:"stars_balance_base"`
	// StarsBalanceBonus = StarsBalanceBase * (1.5 * atom staked percent)
	StarsBalanceBonus sdk.Int `json:"stars_balance_bonus"`
	// StarsPercent = OsmoNormalizedBalance / TotalStarsupply
	StarsPercent sdk.Dec `json:"stars_ownership_percent"`

	StargazeDelegator bool `json:"stargaze_delegator"`
	LiquidityProvider bool `json:"liquidity_provider"`
}

// ExportOsmosisSnapshotCmd generates a snapshot.json from a provided Cosmos Hub genesis export.
func ExportOsmosisSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export-osmosis-snapshot [airdrop-to-denom] [input-genesis-file] [output-snapshot-json]",
		Short: "Export snapshot from a provided Osmosis genesis export",
		Long: `Export snapshot from a provided Osmosis genesis export
Example:
	starsd export-osmosis-snapshot uosmo genesis.json snapshot.json
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			genesisFile := args[1]
			snapshotOutput := args[2]

			// Read genesis file
			genesisJSON, err := os.Open(genesisFile)
			if err != nil {
				return err
			}
			defer genesisJSON.Close()

			// setCosmosBech32Prefixes()

			// Produce the map of address to total atom balance, both staked and unstaked
			snapshotAccs := make(map[string]OsmosisSnapshotAccount)
			totalOsmoBalance := sdk.NewInt(0)

			cdc := clientCtx.Codec

			appState, _, error := genutiltypes.GenesisStateFromGenFile(genesisFile)
			if error != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			LPCount := 0

			bankGenState := banktypes.GetGenesisStateFromAppState(cdc, appState)
			for _, account := range bankGenState.Balances {
				isLP := false
				numCoins := len(account.Coins)
				for i := 0; i < numCoins; i++ {
					denom := account.Coins.GetDenomByIndex(i)
					if strings.HasPrefix(denom, "gamm") {
						isLP = true
						// fmt.Println(denom)
					}
				}
				balance := account.Coins.AmountOf("uosmo")
				totalOsmoBalance = totalOsmoBalance.Add(balance)

				if isLP {
					LPCount++
				}

				snapshotAccs[account.Address] = OsmosisSnapshotAccount{
					OsmoAddress:         account.Address,
					OsmoBalance:         balance,
					OsmoUnstakedBalance: balance,
					OsmoStakedBalance:   sdk.ZeroInt(),
					LiquidityProvider:   isLP,
				}
			}

			stakingGenState := stakingtypes.GetGenesisStateFromAppState(cdc, appState)

			// Make a map from validator operator address to the validator type
			validators := make(map[string]stakingtypes.Validator)
			for _, validator := range stakingGenState.Validators {
				validators[validator.OperatorAddress] = validator
			}

			stargazeDelegators := make(map[string]sdk.Int)

			for _, delegation := range stakingGenState.Delegations {
				address := delegation.DelegatorAddress

				acc, ok := snapshotAccs[address]
				if !ok {
					panic("no account found for delegation")
				}

				val := validators[delegation.ValidatorAddress]
				stakedAtoms := delegation.Shares.MulInt(val.Tokens).Quo(val.DelegatorShares).RoundInt()

				acc.OsmoBalance = acc.OsmoBalance.Add(stakedAtoms)
				acc.OsmoStakedBalance = acc.OsmoStakedBalance.Add(stakedAtoms)

				if delegation.ValidatorAddress == "cosmosvaloper1et77usu8q2hargvyusl4qzryev8x8t9wwqkxfs" {
					stargazeDelegators[address] = stakedAtoms
					acc.StargazeDelegator = true
				}

				snapshotAccs[address] = acc
			}

			totalStarsBalance := sdk.NewInt(0)
			onePointFive := sdk.MustNewDecFromStr("1.5")

			for address, acc := range snapshotAccs {
				allOsmos := acc.OsmoBalance.ToDec()

				// acc.AtomOwnershipPercent = allAtoms.QuoInt(totalAtomBalance)

				if allOsmos.IsZero() {
					acc.OsmoStakedPercent = sdk.ZeroDec()
					acc.StarsBalanceBase = sdk.ZeroInt()
					acc.StarsBalanceBonus = sdk.ZeroInt()
					acc.StarsBalance = sdk.ZeroInt()
					snapshotAccs[address] = acc
					continue
				}

				stakedOsmos := acc.OsmoStakedBalance.ToDec()
				stakedPercent := stakedOsmos.Quo(allOsmos)
				acc.OsmoStakedPercent = stakedPercent

				baseStars, error := allOsmos.ApproxSqrt()
				if error != nil {
					panic(fmt.Sprintf("failed to root atom balance: %s", err))
				}
				acc.StarsBalanceBase = baseStars.RoundInt()

				bonusStars := baseStars.Mul(onePointFive).Mul(stakedPercent)
				acc.StarsBalanceBonus = bonusStars.RoundInt()

				allStars := baseStars.Add(bonusStars)
				// StarsBalance = sqrt( all atoms) * (1 + 1.5) * (staked atom percent) =
				acc.StarsBalance = allStars.RoundInt()

				if allOsmos.LTE(sdk.NewDec(1000000)) {
					acc.StarsBalanceBase = sdk.ZeroInt()
					acc.StarsBalanceBonus = sdk.ZeroInt()
					acc.StarsBalance = sdk.ZeroInt()
				}

				totalStarsBalance = totalStarsBalance.Add(acc.StarsBalance)

				snapshotAccs[address] = acc
			}

			// iterate to find Stars ownership percentage per account
			for address, acc := range snapshotAccs {
				acc.StarsPercent = acc.StarsBalance.ToDec().Quo(totalStarsBalance.ToDec())
				snapshotAccs[address] = acc
			}

			snapshot := OsmosisSnapshot{
				TotalOsmoAmount:         totalOsmoBalance,
				TotalStarsAirdropAmount: totalStarsBalance,
				NumberAccounts:          uint64(len(snapshotAccs)),
				Accounts:                snapshotAccs,
				StargazeDelegators:      stargazeDelegators,
				LiquidityProviderCount:  uint64(LPCount),
			}

			fmt.Printf("num accounts: %d\n", len(snapshotAccs))
			fmt.Printf("osmoTotalSupply: %s\n", totalOsmoBalance.String())
			fmt.Printf("starsTotalSupply: %s\n", totalStarsBalance.String())
			fmt.Printf("num Stargaze delegators: %d\n", len(snapshot.StargazeDelegators))
			fmt.Printf("num LPs: %d\n", LPCount)

			// export snapshot json
			snapshotJSON, err := json.MarshalIndent(snapshot, "", "    ")
			if err != nil {
				return fmt.Errorf("failed to marshal snapshot: %w", err)
			}

			err = ioutil.WriteFile(snapshotOutput, snapshotJSON, 0600)
			return err
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
