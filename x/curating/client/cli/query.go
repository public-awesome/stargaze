package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/public-awesome/stakebird/x/curating/types"
	"github.com/spf13/cobra"
)

// [TODO]
// https://github.com/public-awesome/stakebird/issues/57
// https://github.com/public-awesome/stakebird/issues/58

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group stake queries under a subcommand
	curatingQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	curatingQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdQueryParams(queryRoute, cdc),
		)...,
	)

	return curatingQueryCmd
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current curating parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as curating parameters.
Example:
$ %s query curating params
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", storeName, types.QueryParams)
			bz, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params types.Params
			cdc.MustUnmarshalJSON(bz, &params)
			return cliCtx.PrintOutput(params)
		},
	}
}
