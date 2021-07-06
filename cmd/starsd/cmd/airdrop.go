package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cobra"
)

type Snapshot struct {
	TotalAtomAmount         sdk.Int `json:"total_atom_amount"`
	TotalStarsAirdropAmount sdk.Int `json:"total_stars_amount"`
	NumberAccounts          uint64  `json:"num_accounts"`

	Accounts map[string]SnapshotAccount `json:"accounts"`
}

// SnapshotAccount provide fields of snapshot per account
type SnapshotAccount struct {
	AtomAddress string `json:"atom_address"` // Atom Balance = AtomStakedBalance + AtomUnstakedBalance

	AtomBalance          sdk.Int `json:"atom_balance"`
	AtomOwnershipPercent sdk.Dec `json:"atom_ownership_percent"`

	AtomStakedBalance   sdk.Int `json:"atom_staked_balance"`
	AtomUnstakedBalance sdk.Int `json:"atom_unstaked_balance"` // AtomStakedPercent = AtomStakedBalance / AtomBalance
	AtomStakedPercent   sdk.Dec `json:"atom_staked_percent"`

	// StarsBalance = sqrt( AtomBalance ) * (1 + 1.5 * atom staked percent)
	StarsBalance sdk.Int `json:"stars_balance"`
	// StarsBalanceBase = sqrt(atom balance)
	StarsBalanceBase sdk.Int `json:"stars_balance_base"`
	// StarsBalanceBonus = StarsBalanceBase * (1.5 * atom staked percent)
	StarsBalanceBonus sdk.Int `json:"stars_balance_bonus"`
	// StarsPercent = OsmoNormalizedBalance / TotalStarsupply
	StarsPercent sdk.Dec `json:"stars_ownership_percent"`
}

// setCosmosBech32Prefixes set config for cosmos address system
func setCosmosBech32Prefixes() {
	defaultConfig := sdk.NewConfig()
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(defaultConfig.GetBech32AccountAddrPrefix(), defaultConfig.GetBech32AccountPubPrefix())
	config.SetBech32PrefixForValidator(
		defaultConfig.GetBech32ValidatorAddrPrefix(),
		defaultConfig.GetBech32ValidatorPubPrefix(),
	)
	config.SetBech32PrefixForConsensusNode(
		defaultConfig.GetBech32ConsensusAddrPrefix(),
		defaultConfig.GetBech32ConsensusPubPrefix(),
	)
}

// ExportAirdropSnapshotCmd generates a snapshot.json from a provided cosmos-sdk v0.36 genesis export.
func ExportAirdropSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export-airdrop-snapshot [airdrop-to-denom] [input-genesis-file] [output-snapshot-json]",
		Short: "Export a quadratic fairdrop snapshot from a provided cosmos-sdk genesis export",
		Long: `Export a quadratic fairdrop snapshot from a provided cosmos-sdk genesis export
Sample genesis file:
	https://raw.githubusercontent.com/cephalopodequipment/cosmoshub-3/master/genesis.json
Example:
	starsd export-airdrop-snapshot uatom ~/.gaiad/config/genesis.json ../snapshot.json
	- Check input genesis:
		file is at ~/.gaiad/config/genesis.json
	- Snapshot
		file is at "../snapshot.json"
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			denom := args[0]
			genesisFile := args[1]
			snapshotOutput := args[2]

			// Read genesis file
			genesisJSON, err := os.Open(genesisFile)
			if err != nil {
				return err
			}
			defer genesisJSON.Close()

			setCosmosBech32Prefixes()

			// Produce the map of address to total atom balance, both staked and unstaked
			snapshotAccs := make(map[string]SnapshotAccount)
			totalAtomBalance := sdk.NewInt(0)

			depCdc := clientCtx.JSONMarshaler
			cdc := depCdc.(codec.Marshaler)

			appState, _, error := genutiltypes.GenesisStateFromGenFile(genesisFile)
			if error != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			bankGenState := banktypes.GetGenesisStateFromAppState(cdc, appState)
			for _, account := range bankGenState.Balances {
				balance := account.Coins.AmountOf(denom)
				totalAtomBalance = totalAtomBalance.Add(balance)

				snapshotAccs[account.Address] = SnapshotAccount{
					AtomAddress:         account.Address,
					AtomBalance:         balance,
					AtomUnstakedBalance: balance,
					AtomStakedBalance:   sdk.ZeroInt(),
				}
			}

			stakingGenState := stakingtypes.GetGenesisStateFromAppState(cdc, appState)
			for _, unbonding := range stakingGenState.UnbondingDelegations {
				address := unbonding.DelegatorAddress
				acc, ok := snapshotAccs[address]
				if !ok {
					panic("no account found for unbonding")
				}

				unbondingAtoms := sdk.NewInt(0)
				for _, entry := range unbonding.Entries {
					unbondingAtoms = unbondingAtoms.Add(entry.Balance)
				}

				acc.AtomBalance = acc.AtomBalance.Add(unbondingAtoms)
				acc.AtomUnstakedBalance = acc.AtomUnstakedBalance.Add(unbondingAtoms)

				snapshotAccs[address] = acc
			}

			// Make a map from validator operator address to the v036 validator type
			validators := make(map[string]stakingtypes.Validator)
			for _, validator := range stakingGenState.Validators {
				validators[validator.OperatorAddress] = validator
			}

			for _, delegation := range stakingGenState.Delegations {
				address := delegation.DelegatorAddress

				acc, ok := snapshotAccs[address]
				if !ok {
					panic("no account found for delegation")
				}

				val := validators[delegation.ValidatorAddress]
				stakedAtoms := delegation.Shares.MulInt(val.Tokens).Quo(val.DelegatorShares).RoundInt()

				acc.AtomBalance = acc.AtomBalance.Add(stakedAtoms)
				acc.AtomStakedBalance = acc.AtomStakedBalance.Add(stakedAtoms)

				snapshotAccs[address] = acc
			}

			totalStarsBalance := sdk.NewInt(0)
			onePointFive := sdk.MustNewDecFromStr("1.5")

			for address, acc := range snapshotAccs {
				allAtoms := acc.AtomBalance.ToDec()

				acc.AtomOwnershipPercent = allAtoms.QuoInt(totalAtomBalance)

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

			snapshot := Snapshot{
				TotalAtomAmount:         totalAtomBalance,
				TotalStarsAirdropAmount: totalStarsBalance,
				NumberAccounts:          uint64(len(snapshotAccs)),
				Accounts:                snapshotAccs,
			}

			fmt.Printf("# accounts: %d\n", len(snapshotAccs))
			fmt.Printf("atomTotalSupply: %s\n", totalAtomBalance.String())
			fmt.Printf("starsTotalSupply: %s\n", totalStarsBalance.String())

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
