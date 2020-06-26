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
	"github.com/public-awesome/stakebird/x/curating/types"
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
		// GetCmdDelegate(cdc),
	)...)

	return stakeTxCmd
}

// GetCmdDelegate implements the delegate command.
func GetCmdDelegate(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delegate [validator-addr] [amount] [vendor-id] [post-id]",
		Args:  cobra.ExactArgs(4),
		Short: "Delegate liquid tokens to a validator for curating a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Delegate an amount of liquid coins to a validator from your wallet.
Example:
$ %s tx stake delegate cosmosvaloper1l2rsakp388kuv9k8qzq6lrm9taddae7fpx59wm 1000stake 1 2 --from mykey
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(auth.DefaultTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			// amount, err := sdk.ParseCoin(args[1])
			// if err != nil {
			// 	return err
			// }

			// delAddr := cliCtx.GetFromAddress()
			// valAddr, err := sdk.ValAddressFromBech32(args[0])
			// if err != nil {
			// 	return err
			// }

			// vendorID, err := strconv.ParseUint(args[2], 10, 64)
			// if err != nil {
			// 	return err
			// }

			// postID, err := strconv.ParseUint(args[3], 10, 64)
			// if err != nil {
			// 	return err
			// }

			// msg := types.NewMsgDelegate(vendorID, postID, delAddr, valAddr, amount)
			// return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{})
		},
	}
}

// GetCmdPost implements the delegate command.
func GetCmdPost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "post [deposit] [vendor-id] [post-id] [body] [reward_address]",
		Args:  cobra.MinimumNArgs(4),
		Short: "Create a post with a deposit",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a post with a deposit.
Example:
$ %s tx curating post 1000stake 1 2 "body"  --from mykey
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(auth.DefaultTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			stake, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			creator := cliCtx.GetFromAddress()

			vendorID, err := strconv.ParseUint(args[2], 10, 32)
			if err != nil {
				return err
			}

			var postID string
			if len(args) > 2 {
				postID = args[3]
			}

			var body string
			if len(args) > 3 {
				body = args[4]
			}

			var rewardAddrStr string
			var rewardAddr sdk.AccAddress
			if len(args) > 4 {
				rewardAddrStr = args[5]
			}
			if rewardAddrStr != "" {
				rewardAddr, err = sdk.AccAddressFromBech32(rewardAddrStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgPost(uint32(vendorID), postID, creator, rewardAddr, body, stake)

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
