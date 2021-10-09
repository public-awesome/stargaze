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
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cobra"
)

type Snapshot struct {
	TotalStarsAirdropAmount sdk.Int                    `json:"total_stars_amount"`
	Accounts                map[string]SnapshotAccount `json:"accounts"`
}

type SnapshotAccount struct {
	AtomAddress              string `json:"atom_address"`
	OsmoAddress              string `json:"osmo_address"`
	RegenAddress             string `json:"regen_address"`
	StargazeHubDelegator     bool   `json:"hub_delegator"`
	StargazeOsmosisDelegator bool   `json:"osmosis_delegator"`
	StargazeRegenDelegator   bool   `json:"regen_delegator"`
	OsmosisLiquidityProvider bool   `json:"osmosis_lp"`
}

func ExportSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export-snapshot [input-hub-snapshot] [input-osmo-snapshot] [input-regen-snapshot] [output-snapshot]",
		Short: "Export final snapshot from a provided snapshots",
		Long: `Export final snapshot from a provided snapshots
Example:
	starsd export-snapshot hub-snapshot.json osmo-snapshot.json regen-snapshot.json snapshot.json
`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			hubSnapshotFile := args[0]
			osmoSnapshotFile := args[1]
			regenSnapshotFile := args[2]
			snapshotOutput := args[3]

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
					AtomAddress:       address,
					AtomStakedBalance: sdk.ZeroInt(),
					StargazeDelegator: false,
				}

				acc, ok := snapshotAccs[address]
				if !ok {
					panic("no account found for delegation")
				}

				val := validators[delegation.ValidatorAddress]
				stakedAtoms := delegation.Shares.MulInt(val.Tokens).Quo(val.DelegatorShares).RoundInt()

				acc.AtomStakedBalance = acc.AtomStakedBalance.Add(stakedAtoms)
				acc.StarsBalance = sdk.NewInt(50)

				if delegation.ValidatorAddress == "cosmosvaloper1et77usu8q2hargvyusl4qzryev8x8t9wwqkxfs" {
					stargazeDelegators[address] = stakedAtoms
					acc.StargazeDelegator = true
					acc.StarsBalance = acc.StarsBalance.Add(sdk.NewInt(50))
				}

				snapshotAccs[address] = acc
			}

			totalStarsBalance := sdk.NewInt(0)
			for _, acc := range snapshotAccs {
				totalStarsBalance = totalStarsBalance.Add(acc.StarsBalance)
			}

			snapshot := HubSnapshot{
				TotalAtomAmount:         totalAtomBalance,
				TotalStarsAirdropAmount: totalStarsBalance,
				NumStakers:              uint64(len(snapshotAccs)),
				Stakers:                 snapshotAccs,
				StargazeDelegators:      stargazeDelegators,
			}

			fmt.Printf("num stakers: %d\n", len(snapshotAccs))
			// fmt.Printf("atomTotalSupply: %s\n", totalAtomBalance.String())
			// fmt.Printf("starsTotalSupply: %s\n", totalStarsBalance.String())
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
