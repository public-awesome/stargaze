package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	claimtypes "github.com/public-awesome/stargaze/x/claim/types"
	"github.com/spf13/cobra"
)

const Denom = "stars"

func AddAirdropCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-airdrop [airdrop-snapshot-file]",
		Short: "Add balances of accounts to claim module.",
		Args:  cobra.ExactArgs(1),
		Long: fmt.Sprintf(`Add balances of accounts to claim module.
Example:
$ %s add-airdrop /path/to/snapshot.json
`, version.AppName),

		RunE: func(cmd *cobra.Command, args []string) error {

			var clientCtx = client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			// read snapshot
			snapshotFile := args[0]
			snapshotJSON, _ := ioutil.ReadFile(snapshotFile)
			snapshot := Snapshot{}
			json.Unmarshal([]byte(snapshotJSON), &snapshot)

			genFile := config.GenesisFile()

			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			fmt.Printf("Accounts %d\n", len(snapshot.Accounts))

			claimGenState := claimtypes.GetGenesisStateFromAppState(cdc, appState)

			// [TODO] add claim genesis params
			// [TODO] change denom to stars
			// [TODO] module account balance?
			// [TODO] remove denom in claim?

			for address, acc := range snapshot.Accounts {
				// empty account check
				if acc.AirdropAmount.LTE(sdk.NewInt(0)) {
					panic("Empty account")
				}

				coin := sdk.NewCoin(Denom, acc.AirdropAmount)
				coins := sdk.NewCoins(coin)

				record := claimtypes.ClaimRecord{
					Address:                address,
					InitialClaimableAmount: coins,
					ActionCompleted:        []bool{false, false, false, false},
				}
				claimGenState.ClaimRecords = append(claimGenState.ClaimRecords, record)
			}
			// claimGenState.Params = genesisParams.ClaimParams
			claimGenStateBz, err := cdc.MarshalJSON(claimGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal claim genesis state: %w", err)
			}
			appState[claimtypes.ModuleName] = claimGenStateBz

			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON

			fmt.Printf("Saving genesis...")
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	return cmd
}
