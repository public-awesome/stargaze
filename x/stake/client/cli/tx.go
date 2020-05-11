package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/rocket-protocol/stakebird/x/stake/types"
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
		GetCmdDelegate(cdc),
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

			amount, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			delAddr := cliCtx.GetFromAddress()
			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			vendorID, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			postID, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgDelegate(vendorID, postID, delAddr, valAddr, amount)
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdPost implements the delegate command.
func GetCmdPost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "post [validator-addr] [amount] [vendor-id] [post-id] [body] [voting-period]",
		Args:  cobra.ExactArgs(4),
		Short: "Delegate liquid tokens to a validator for creating a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Delegate an amount of liquid coins to a validator from your wallet.
Example:
$ %s tx stake post cosmosvaloper1l2rsakp388kuv9k8qzq6lrm9taddae7fpx59wm 1000stake 1 2 "body" 72h  --from mykey
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(auth.DefaultTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			amount, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			delAddr := cliCtx.GetFromAddress()
			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			vendorID, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			postID, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			var body string
			if len(args) > 3 {
				body = args[4]
			}

			var votingPeriod time.Duration
			if len(args) > 4 {
				votingPeriod, err = time.ParseDuration(args[5])
				if err != nil {
					panic("Failed parsing voting period")
				}
			}

			msgDel := types.NewMsgDelegate(vendorID, postID, delAddr, valAddr, amount)
			msg := types.NewMsgPost(body, msgDel, votingPeriod)

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
