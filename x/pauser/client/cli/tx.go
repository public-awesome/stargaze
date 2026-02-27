package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/public-awesome/stargaze/v18/x/pauser/types"
	"github.com/spf13/cobra"
)

// GetTxCmd builds tx command group for the module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdPauseContract())
	cmd.AddCommand(CmdUnpauseContract())
	cmd.AddCommand(CmdPauseCodeID())
	cmd.AddCommand(CmdUnpauseCodeID())

	return cmd
}

func CmdPauseContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause-contract [contract-address]",
		Short: "Pause execution of a contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgPauseContract{
				Sender:          clientCtx.GetFromAddress().String(),
				ContractAddress: args[0],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdUnpauseContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unpause-contract [contract-address]",
		Short: "Unpause execution of a contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgUnpauseContract{
				Sender:          clientCtx.GetFromAddress().String(),
				ContractAddress: args[0],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdPauseCodeID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause-code-id [code-id]",
		Short: "Pause execution of all contracts with a code ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			codeID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := &types.MsgPauseCodeID{
				Sender: clientCtx.GetFromAddress().String(),
				CodeID: codeID,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdUnpauseCodeID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unpause-code-id [code-id]",
		Short: "Unpause execution of all contracts with a code ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			codeID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := &types.MsgUnpauseCodeID{
				Sender: clientCtx.GetFromAddress().String(),
				CodeID: codeID,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
