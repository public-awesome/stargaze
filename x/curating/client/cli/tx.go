package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/public-awesome/stakebird/x/curating/types"
	"github.com/spf13/cobra"
)

// NewPostTxCmd returns the post command
func NewPostTxCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "post [vendor-id] [post-id] [body] [reward_address] --from [key]",
		Args:  cobra.MinimumNArgs(3),
		Short: "Register a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a post.
Example:
$ %s tx curating post 1 "2" "body" --from mykey
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress()

			vendorID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			postID := args[1]
			body := args[2]

			var rewardAddrStr string
			var rewardAddr sdk.AccAddress
			if len(args) > 3 {
				rewardAddrStr = args[3]
			}
			if rewardAddrStr != "" {
				rewardAddr, err = sdk.AccAddressFromBech32(rewardAddrStr)
				if err != nil {
					return err
				}
			}
			msg := types.NewMsgPost(uint32(vendorID), postID, creator, rewardAddr, body)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewUpvoteTxCmd returns the upvote command
func NewUpvoteTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upvote [vendor-id] [post-id] [voteNum] [reward-addr] --from [key]",
		Args:  cobra.MinimumNArgs(3),
		Short: "Upvote a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Curating a post by upvoting.
Example:
$ %s tx curating upvote 1 "2" 5 --from mykey
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			curator := clientCtx.GetFromAddress()

			vendorID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			postID := args[1]

			voteNum, err := strconv.ParseUint(args[2], 10, 32)
			if err != nil {
				return err
			}

			var rewardAddrStr string
			var rewardAddr sdk.AccAddress
			if len(args) > 3 {
				rewardAddrStr = args[3]
			}
			if rewardAddrStr != "" {
				rewardAddr, err = sdk.AccAddressFromBech32(rewardAddrStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgUpvote(
				uint32(vendorID), postID, curator, rewardAddr, int32(voteNum))

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
