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

type HubSnapshot struct {
	TotalAtomAmount         sdk.Int                       `json:"total_atom_amount"`
	TotalStarsAirdropAmount sdk.Int                       `json:"total_stars_amount"`
	NumStakers              uint64                        `json:"num_stakers"`
	Stakers                 map[string]HubSnapshotAccount `json:"accounts"`
	StargazeDelegators      map[string]sdk.Int            `json:"stargaze_delegators"`
}

// HubSnapshotAccount provide fields of snapshot per account
type HubSnapshotAccount struct {
	AtomAddress       string  `json:"atom_address"`
	AtomStakedBalance sdk.Int `json:"atom_staked_balance"`
	StarsBalance      sdk.Int `json:"stars_balance"`
	StargazeDelegator bool    `json:"stargaze_delegator"`
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

// ExportHubSnapshotCmd generates a snapshot.json from a provided Cosmos Hub genesis export.
func ExportHubSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export-hub-snapshot [airdrop-to-denom] [input-genesis-file] [output-snapshot-json]",
		Short: "Export snapshot from a provided Cosmos Hub genesis export",
		Long: `Export snapshot from a provided Cosmos Hub genesis export
Example:
	starsd export-hub-snapshot uatom genesis.json snapshot.json
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			// denom := args[0]
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

				if delegation.ValidatorAddress == "cosmosvaloper1et77usu8q2hargvyusl4qzryev8x8t9wwqkxfs" {
					stargazeDelegators[address] = stakedAtoms
					acc.StargazeDelegator = true
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
