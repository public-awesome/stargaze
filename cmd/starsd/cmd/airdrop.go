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
	"github.com/spf13/cobra"
)

const Denom = "stars"

func AddAirdrop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-airdrop [airdrop-snapshot-file]",
		Short: "Add balances of accounts to claim module.",
		Args:  cobra.ExactArgs(3),
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
			snapshotFile := args[1]
			snapshotJSON, _ := ioutil.ReadFile(snapshotFile)
			snapshot := Snapshot{}
			json.Unmarshal([]byte(snapshotJSON), &snapshot)

			genFile := config.GenesisFile()

			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			fmt.Printf("Accounts %d\n", len(snapshot.Accounts))

			for address, acc := range snapshot.Accounts {

				// empty account check
				if acc.AirdropAmount.LTE(sdk.NewInt(0)) {
					panic("Empty account")
				}

				addr, _ := sdk.AccAddressFromBech32(address)

				coin := sdk.NewCoin(Denom, acc.AirdropAmount)
				coins := sdk.NewCoins(coin)
			}

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
