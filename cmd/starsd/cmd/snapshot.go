package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
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
	StargazeHubDelegator     bool   `json:"sg_hub_delegator"`
	StargazeOsmosisDelegator bool   `json:"sg_osmosis_delegator"`
	StargazeRegenDelegator   bool   `json:"sg_regen_delegator"`
	AtomStaker               bool   `json:"atom_staker"`
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
			// regenSnapshotFile := args[2]
			snapshotOutput := args[3]

			hubJSON, _ := ioutil.ReadFile(hubSnapshotFile)
			osmoJSON, _ := ioutil.ReadFile(osmoSnapshotFile)
			// regenJSON, _ := ioutil.ReadFile(regenSnapshotFile)

			snapshotAccs := make(map[string]SnapshotAccount)

			hubSnapshot := HubSnapshot{}
			json.Unmarshal([]byte(hubJSON), &hubSnapshot)
			for _, staker := range hubSnapshot.Accounts {
				starsAddr, _ := ConvertCosmosAddressToStargaze(staker.AtomAddress)
				snapshotAcc := SnapshotAccount{
					AtomAddress:              staker.AtomAddress,
					OsmoAddress:              "",
					RegenAddress:             "",
					StargazeHubDelegator:     staker.StargazeDelegator,
					StargazeOsmosisDelegator: false,
					StargazeRegenDelegator:   false,
					AtomStaker:               true,
					OsmosisLiquidityProvider: false,
				}
				snapshotAccs[starsAddr.String()] = snapshotAcc
			}

			osmosisSnapshot := OsmosisSnapshot{}
			json.Unmarshal([]byte(osmoJSON), &osmosisSnapshot)
			for _, acct := range osmosisSnapshot.Accounts {
				starsAddr, _ := ConvertCosmosAddressToStargaze(acct.OsmoAddress)

				if acc, ok := snapshotAccs[starsAddr.String()]; ok {
					acc.OsmoAddress = acct.OsmoAddress
					acc.StargazeOsmosisDelegator = acct.StargazeDelegator
					acc.OsmosisLiquidityProvider = acct.LiquidityProvider
					snapshotAccs[starsAddr.String()] = acc
				} else {
					snapshotAcc := SnapshotAccount{
						OsmoAddress:              acct.OsmoAddress,
						StargazeOsmosisDelegator: acct.StargazeDelegator,
						OsmosisLiquidityProvider: acct.LiquidityProvider,
					}
					snapshotAccs[starsAddr.String()] = snapshotAcc
				}
			}

			snapshot := Snapshot{
				TotalStarsAirdropAmount: sdk.Int{},
				Accounts:                snapshotAccs,
			}

			fmt.Printf("accounts: %d\n", len(snapshotAccs))

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

func ConvertCosmosAddressToStargaze(address string) (sdk.AccAddress, error) {
	config := sdk.GetConfig()
	starsPrefix := config.GetBech32AccountAddrPrefix()

	_, bytes, err := bech32.DecodeAndConvert(address)
	if err != nil {
		return nil, err
	}

	newAddr, err := bech32.ConvertAndEncode(starsPrefix, bytes)
	if err != nil {
		return nil, err
	}

	sdkAddr, err := sdk.AccAddressFromBech32(newAddr)
	if err != nil {
		return nil, err
	}

	return sdkAddr, nil
}
