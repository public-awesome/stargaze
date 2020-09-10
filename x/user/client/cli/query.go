package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/public-awesome/stakebird/x/user/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group stake queries under a subcommand
	userQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	userQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdQueryParams(queryRoute, cdc),
			// GetCmdQueryPost(queryRoute, cdc),
		)...,
	)

	return userQueryCmd
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current user parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as user parameters.
Example:
$ %s query user params
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

// GetCmdQueryPost implements the post query command.
// func GetCmdQueryPost(storeName string, cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "post [vendor-id] [post-id]",
// 		Args:  cobra.ExactArgs(2),
// 		Short: "Query for a post by vendor ID and post ID",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Query post by vendor ID and post ID.
// Example:
// $ %s query user posts 1 123
// `,
// 				version.ClientName,
// 			),
// 		),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)

// 			vendorID := args[0]
// 			postID := args[1]

// 			route := fmt.Sprintf("custom/%s/%s/%s/%s", storeName, types.QueryPost, vendorID, postID)
// 			bz, _, err := cliCtx.QueryWithData(route, nil)
// 			if err != nil {
// 				return err
// 			}

// 			var post types.Post
// 			cdc.MustUnmarshalJSON(bz, &post)
// 			return cliCtx.PrintOutput(post)
// 		},
// 	}
// }
