package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/public-awesome/stakebird/x/user/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	stakeTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	stakeTxCmd.AddCommand(flags.PostCommands(
		GetCmdPost(cdc),
		GetCmdUpvote(cdc),
	)...)

	return stakeTxCmd
}

// GetCmdPost implements the delegate command.
func GetCmdPost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "post [vendor-id] [post-id] [body] [reward_address]",
		Args:  cobra.MinimumNArgs(3),
		Short: "Register a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a post.
Example:
$ %s tx user post 1 "2" "body" --from mykey
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(auth.DefaultTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			creator := cliCtx.GetFromAddress()

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

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdUpvote implements the upvote command.
func GetCmdUpvote(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "upvote [vendor-id] [post-id] [voteNum] [reward-addr]",
		Args:  cobra.MinimumNArgs(3),
		Short: "Upvote a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`user a post by upvoting.
Example:
$ %s tx user upvote 1 "2" 5 --from mykey
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(auth.DefaultTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			curator := cliCtx.GetFromAddress()

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
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
