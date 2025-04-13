package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/public-awesome/stargaze/v16/x/cron/types"
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdPromoteToPrivilegedContract())
	cmd.AddCommand(CmdDemoteFromPrivilegedContract())

	return cmd
}

func CmdPromoteToPrivilegedContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "promote-to-privilege-contract [contract_addr_bech32]",
		Short:   "Promotes the specified contract",
		Long:    "Promotes the given contract to privilege status which enables the contract to hook on to abci.BeginBlocker and abci.EndBlocker",
		Aliases: []string{"promote-contract"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			contract := args[0]
			msg := types.NewMsgPromoteToPrivilegedContract(clientCtx.GetFromAddress().String(), contract)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDemoteFromPrivilegedContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "demote-from-privilege-contract [contract_addr_bech32]",
		Short:   "Demotes the specified contract",
		Long:    "Demotes the given contract to privilege status which disables the contract to hook on to abci.BeginBlocker and abci.EndBlocker",
		Aliases: []string{"demote-contract"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			contract := args[0]
			msg := types.NewMsgDemoteFromPrivilegedContract(clientCtx.GetFromAddress().String(), contract)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
