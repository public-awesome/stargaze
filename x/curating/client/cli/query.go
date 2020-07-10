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
			GetCmdQueryPost(queryRoute, cdc),
			GetCmdQueryPosts(queryRoute, cdc),
			GetCmdQueryUpvotes(queryRoute, cdc),
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

// GetCmdQueryPost implements the post query command.
func GetCmdQueryPost(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "post [vendor-id] [post-id]",
		Args:  cobra.MinimumNArgs(2),
		Short: "Query for a post by vendor ID and post ID",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query post by vendor ID and post ID.
Example:
$ %s query curating posts 1 123
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			vendorID := args[0]
			postID := args[1]

			route := fmt.Sprintf("custom/%s/%s/%s/%s", storeName, types.QueryPost, vendorID, postID)
			bz, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var post types.Post
			cdc.MustUnmarshalJSON(bz, &post)
			return cliCtx.PrintOutput(post)
		},
	}
}

// GetCmdQueryPosts implements the posts query command.
func GetCmdQueryPosts(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "posts [vendor-id]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Query all posts for a given vendor ID",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query posts for a given vendor ID.
Example:
$ %s query curating posts 1
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			vendorID := args[0]

			route := fmt.Sprintf("custom/%s/%s/%s", storeName, types.QueryPosts, vendorID)
			bz, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var posts []types.Post
			cdc.MustUnmarshalJSON(bz, &posts)
			return cliCtx.PrintOutput(posts)
		},
	}
}

// GetCmdQueryUpvote implements the upvotes query command.
func GetCmdQueryUpvotes(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "upvote [vendor-id] [post-id]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Query for an upvote by vendor ID and post ID",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query upvote by vendor ID and optionally post ID.
Example:
$ %s query curating upvotes 1 "123"

or...

$ %s query curating upvotes 1
`,
				version.ClientName,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			vendorID := args[0]
			var route string
			if len(args) > 1 {
				postID := args[1]
				route = fmt.Sprintf("custom/%s/%s/%s/%s", storeName, types.QueryUpvotes, vendorID, postID)
			} else {
				route = fmt.Sprintf("custom/%s/%s/%s", storeName, types.QueryUpvotes, vendorID)
			}

			bz, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var upvote []types.Upvote
			cdc.MustUnmarshalJSON(bz, &upvote)
			return cliCtx.PrintOutput(upvote)
		},
	}
}
