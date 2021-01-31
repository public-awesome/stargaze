package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/public-awesome/stargaze/x/stake/types"
	"github.com/spf13/cobra"
)

// NewStakesQueryCmd defines the command to query posts from a vendor_id
func NewStakesQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stakes [vendor-id] [post-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Query for stakes by vendor and post id",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query stakes by vendor and post id.
Example:
$ %s query stake stakes 1 123
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
			req := &types.QueryStakesRequest{
				VendorId: uint32(vendorID),
				PostId:   postID,
			}

			res, err := queryClient.Stakes(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// NewStakeQueryCmd defines the command to query a stake
func NewStakeQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake [vendor-id] [post-id] [delegator]",
		Args:  cobra.ExactArgs(3),
		Short: "Query for a specific stake",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for a specific stake.
Example:
$ %s query stake stake 1 123 stbdeadbeef
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

			delegator := strings.TrimSpace(args[2])
			if delegator == "" {
				return fmt.Errorf("invalid delegator")
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryStakeRequest{
				VendorId:  uint32(vendorID),
				PostId:    postID,
				Delegator: delegator,
			}

			res, err := queryClient.Stake(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
