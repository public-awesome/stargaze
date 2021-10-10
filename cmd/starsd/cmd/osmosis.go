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
	Accounts               map[string]OsmosisSnapshotAccount `json:"accounts"`
	LiquidityProviderCount uint64                            `json:"lp_count"`
}

// OsmosisSnapshotAccount provide fields of snapshot per account
type OsmosisSnapshotAccount struct {
	OsmoAddress       string `json:"osmo_address"`
	StargazeDelegator bool   `json:"stargaze_delegator"`
	LiquidityProvider bool   `json:"liquidity_provider"`
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
					}
				}
				balance := account.Coins.AmountOf("uosmo")
				totalOsmoBalance = totalOsmoBalance.Add(balance)

				if isLP {
					LPCount++
					snapshotAccs[account.Address] = OsmosisSnapshotAccount{
						OsmoAddress:       account.Address,
						LiquidityProvider: true,
					}
				}
			}

			stakingGenState := stakingtypes.GetGenesisStateFromAppState(cdc, appState)
			for _, delegation := range stakingGenState.Delegations {
				address := delegation.DelegatorAddress

				if delegation.ValidatorAddress == "osmovaloper1et77usu8q2hargvyusl4qzryev8x8t9weceqyk" {
					acc, ok := snapshotAccs[address]
					if !ok {
						// account does not exist
						snapshotAccs[address] = OsmosisSnapshotAccount{
							OsmoAddress:       address,
							LiquidityProvider: false,
							StargazeDelegator: true,
						}
					} else {
						// account exists
						acc.StargazeDelegator = true
						snapshotAccs[address] = acc
					}
				}
			}

			snapshot := OsmosisSnapshot{
				Accounts:               snapshotAccs,
				LiquidityProviderCount: uint64(LPCount),
			}

			fmt.Printf("num accounts: %d\n", len(snapshotAccs))
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
