package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/public-awesome/stargaze/x/staking/types"
	"github.com/spf13/cobra"
)

type OsmosisSnapshot struct {
	TotalOsmoAmount         sdk.Int                       `json:"total_osmo_amount"`
	TotalStarsAirdropAmount sdk.Int                       `json:"total_stars_amount"`
	NumberAccounts          uint64                        `json:"num_accounts"`
	Accounts                map[string]HubSnapshotAccount `json:"accounts"`
	StargazeDelegators      map[string]sdk.Int            `json:"stargaze_delegators"`
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
			snapshotAccs := make(map[string]HubSnapshotAccount)
			totalAtomBalance := sdk.NewInt(0)

			cdc := clientCtx.Codec

			appState, _, error := genutiltypes.GenesisStateFromGenFile(genesisFile)
			if error != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
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

				snapshotAccs[address] = HubSnapshotAccount{
					AtomAddress:         address,
					AtomBalance:         sdk.ZeroInt(),
					AtomUnstakedBalance: sdk.ZeroInt(),
					AtomStakedBalance:   sdk.ZeroInt(),
				}

				acc, ok := snapshotAccs[address]
				if !ok {
					panic("no account found for delegation")
				}

				val := validators[delegation.ValidatorAddress]
				stakedAtoms := delegation.Shares.MulInt(val.Tokens).Quo(val.DelegatorShares).RoundInt()

				acc.AtomBalance = acc.AtomBalance.Add(stakedAtoms)
				acc.AtomStakedBalance = acc.AtomStakedBalance.Add(stakedAtoms)

				if delegation.ValidatorAddress == "cosmosvaloper1et77usu8q2hargvyusl4qzryev8x8t9wwqkxfs" {
					stargazeDelegators[address] = stakedAtoms
					acc.StargazeDelegator = true
				}

				snapshotAccs[address] = acc
			}

			totalStarsBalance := sdk.NewInt(0)
			onePointFive := sdk.MustNewDecFromStr("1.5")

			for address, acc := range snapshotAccs {
				allAtoms := acc.AtomBalance.ToDec()

				// acc.AtomOwnershipPercent = allAtoms.QuoInt(totalAtomBalance)

				if allAtoms.IsZero() {
					acc.AtomStakedPercent = sdk.ZeroDec()
					acc.StarsBalanceBase = sdk.ZeroInt()
					acc.StarsBalanceBonus = sdk.ZeroInt()
					acc.StarsBalance = sdk.ZeroInt()
					snapshotAccs[address] = acc
					continue
				}

				stakedAtoms := acc.AtomStakedBalance.ToDec()
				stakedPercent := stakedAtoms.Quo(allAtoms)
				acc.AtomStakedPercent = stakedPercent

				baseStars, error := allAtoms.ApproxSqrt()
				if error != nil {
					panic(fmt.Sprintf("failed to root atom balance: %s", err))
				}
				acc.StarsBalanceBase = baseStars.RoundInt()

				bonusStars := baseStars.Mul(onePointFive).Mul(stakedPercent)
				acc.StarsBalanceBonus = bonusStars.RoundInt()

				allStars := baseStars.Add(bonusStars)
				// StarsBalance = sqrt( all atoms) * (1 + 1.5) * (staked atom percent) =
				acc.StarsBalance = allStars.RoundInt()

				if allAtoms.LTE(sdk.NewDec(1000000)) {
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

			snapshot := HubSnapshot{
				TotalAtomAmount:         totalAtomBalance,
				TotalStarsAirdropAmount: totalStarsBalance,
				NumberAccounts:          uint64(len(snapshotAccs)),
				Accounts:                snapshotAccs,
				StargazeDelegators:      stargazeDelegators,
			}

			fmt.Printf("num accounts: %d\n", len(snapshotAccs))
			fmt.Printf("atomTotalSupply: %s\n", totalAtomBalance.String())
			fmt.Printf("starsTotalSupply: %s\n", totalStarsBalance.String())
			fmt.Printf("num Stargaze delegators: %d\n", len(snapshot.StargazeDelegators))

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
