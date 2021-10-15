package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	lockuptypes "github.com/osmosis-labs/osmosis/x/lockup/types"
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
	OsmoStaker        bool   `json:"osmo_staker"`
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

			// Produce the map of address
			snapshotAccs := make(map[string]OsmosisSnapshotAccount)

			cdc := clientCtx.Codec

			appState, _, error := genutiltypes.GenesisStateFromGenFile(genesisFile)
			if error != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			LPCount := 0

			// case1: LP token holders
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

				if isLP {
					LPCount++
					snapshotAccs[account.Address] = OsmosisSnapshotAccount{
						OsmoAddress:       account.Address,
						LiquidityProvider: true,
					}
				}
			}

			// case 2: accounts with locked LP tokens
			lockupGenState := GetLockupGenesisStateFromAppState(cdc, appState)
			for _, lock := range lockupGenState.Locks {
				addr := lock.Owner
				acc, ok := snapshotAccs[addr]
				if !ok {
					// account does not exist
					snapshotAccs[addr] = OsmosisSnapshotAccount{
						OsmoAddress:       addr,
						LiquidityProvider: true,
						StargazeDelegator: false,
					}
					LPCount++
				} else {
					// account exists
					acc.LiquidityProvider = true
					snapshotAccs[addr] = acc
				}
			}

			stakerCount := 0
			delegatorCount := 0
			// case 3: stakers + delegators to stargaze
			stakingGenState := stakingtypes.GetGenesisStateFromAppState(cdc, appState)
			for _, delegation := range stakingGenState.Delegations {
				address := delegation.DelegatorAddress

				acc, ok := snapshotAccs[address]
				if !ok {
					// account does not exist
					snapshotAccs[address] = OsmosisSnapshotAccount{
						OsmoAddress:       address,
						OsmoStaker:        true,
						LiquidityProvider: false,
						StargazeDelegator: false,
					}
				} else {
					// account exists
					acc.OsmoStaker = true
					snapshotAccs[address] = acc
				}
				stakerCount++

				if delegation.ValidatorAddress == "osmovaloper1et77usu8q2hargvyusl4qzryev8x8t9weceqyk" {
					acc.StargazeDelegator = true
					snapshotAccs[address] = acc
					delegatorCount++
				}
			}

			snapshot := OsmosisSnapshot{
				Accounts:               snapshotAccs,
				LiquidityProviderCount: uint64(LPCount),
			}

			fmt.Printf("num accounts: %d\n", len(snapshotAccs))
			fmt.Printf("num LPs: %d\n", LPCount)
			fmt.Printf("num Stargaze delegators: %d\n", delegatorCount)
			fmt.Printf("num Osmo stakers: %d\n", stakerCount)

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

func GetLockupGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *lockuptypes.GenesisState {
	var genesisState lockuptypes.GenesisState

	if appState["lockup"] != nil {
		cdc.MustUnmarshalJSON(appState["lockup"], &genesisState)
	}

	return &genesisState
}
