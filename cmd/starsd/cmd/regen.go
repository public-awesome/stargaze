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

type RegenSnapshot struct {
	Accounts map[string]RegenSnapshotAccount `json:"accounts"`
}

type RegenSnapshotAccount struct {
	RegenAddress      string `json:"atom_address"`
	StargazeDelegator bool   `json:"stargaze_delegator"`
}

// ExportHubSnapshotCmd generates a snapshot.json from a provided Cosmos Hub genesis export.
func ExportRegenSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export-regen-snapshot [input-genesis-file] [output-snapshot-json]",
		Short: "Export snapshot from a provided Regen genesis export",
		Long: `Export snapshot from a provided Regen genesis export
Example:
	starsd export-regen-snapshot genesis.json regen-snapshot.json
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			genesisFile := args[0]
			snapshotOutput := args[1]

			// Read genesis file
			genesisJSON, err := os.Open(genesisFile)
			if err != nil {
				return err
			}
			defer genesisJSON.Close()

			// Produce the map of address to total atom balance, both staked and unstaked
			snapshotAccs := make(map[string]RegenSnapshotAccount)

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

			for _, delegation := range stakingGenState.Delegations {
				val, ok := validators[delegation.ValidatorAddress]
				if !ok {
					panic(fmt.Sprintf("missing validator %s ", delegation.GetValidatorAddr()))
				}

				address := delegation.DelegatorAddress
				delegationAmount := val.TokensFromShares(delegation.Shares).Quo(sdk.NewDec(1_000_000))
				// MIN 1REGEN
				if delegationAmount.LT(sdk.NewDec(0)) {
					continue
				}

				if delegation.ValidatorAddress == "regenvaloper1me6xrxj80cp4smrmy2mwaan67spnsen9n2sjxq" {
					acc, ok := snapshotAccs[address]
					if !ok {
						// account does not exist
						snapshotAccs[address] = RegenSnapshotAccount{
							RegenAddress:      address,
							StargazeDelegator: true,
						}
					} else {
						// account exists
						acc.StargazeDelegator = true
						snapshotAccs[address] = acc
					}
				}
			}

			snapshot := RegenSnapshot{
				Accounts: snapshotAccs,
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
