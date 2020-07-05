package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/public-awesome/stakebird/x/funding/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	fundingTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	fundingTxCmd.AddCommand(flags.PostCommands(
		GetCmdBuy(cdc),
	)...)

	return fundingTxCmd
}

func GetCmdBuy(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "buy [max_amount] [quantity]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Buy FUEL with ATOM reserves from the bonding curve",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Locks collateral that will be used as reserves for the bonding curve. Mints new FUEL.
Example:
$ %s tx funding buy 1000stake --from mykey
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(auth.DefaultTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			maxAmount, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}

			quantity, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			senderAddr := cliCtx.GetFromAddress()
			msg := types.NewMsgBuy(maxAmount, quantity, senderAddr)

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
