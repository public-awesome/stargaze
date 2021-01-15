package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/public-awesome/stakebird/x/curating/types"
	"github.com/spf13/cobra"
)

// NewPostsQueryCmd defines the command to query posts from a vendor_id
func NewPostsQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "posts [vendor-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for posts by vendor ID",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query posts by vendor ID.
Example:
$ %s query curating posts 1
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			vendorID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryPostsRequest{
				VendorId: uint32(vendorID),
			}

			res, err := queryClient.Posts(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// NewPostQueryCmd defines the command to query a post from a vendor_id,post_id
func NewPostQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post [vendor-id] [post-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Query for a post by vendor ID and post ID",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query post by vendor ID and post ID.
Example:
$ %s query curating post 1 123
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			vendorID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}
			postID := strings.TrimSpace(args[1])

			if postID == "" {
				return fmt.Errorf("invalid post id")
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryPostRequest{
				VendorId: uint32(vendorID),
				PostId:   postID,
			}

			res, err := queryClient.Post(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// NewUpvotesQueryCmd defines the command to query a post from a vendor_id,post_id
func NewUpvotesQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upvotes [vendor-id] [post-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Query for upvotes by vendor ID and post ID",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query upvotes by vendor ID and post ID.

Example:
$ %s query curating upvotes 1 "123"
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			postID := strings.TrimSpace(args[1])
			if postID == "" {
				return fmt.Errorf("invalid post id")
			}

			vendorID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			qClient := types.NewQueryClient(clientCtx)
			queryUpvotes := &types.QueryUpvotesRequest{
				VendorId: uint32(vendorID),
				PostId:   postID,
			}

			res, err := qClient.Upvotes(context.Background(), queryUpvotes)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
