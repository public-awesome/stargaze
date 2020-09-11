package cli

import (
	"bufio"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
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
		GetCmdVouch(cdc),
	)...)

	return stakeTxCmd
}

// GetCmdVouch implements the delegate command.
func GetCmdVouch(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vouch [vouched] [comment]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Vouch a user",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Vouch a user.
Example:
$ %s tx user vouch [vouched-address] "comment here" --from mykey
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(auth.DefaultTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			voucher := cliCtx.GetFromAddress()
			vouched, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			comment := ""
			if len(args) > 1 {
				comment = args[1]
			}

			msg := types.NewMsgVouch(voucher, vouched, comment)

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
