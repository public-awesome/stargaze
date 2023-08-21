package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/public-awesome/stargaze/v12/x/globalfee/types"
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

	cmd.AddCommand(CmdSetCodeAuthorization())
	cmd.AddCommand(CmdRemoveCodeAuthorization())
	cmd.AddCommand(CmdSetContractAuthorization())
	cmd.AddCommand(CmdRemoveContractAuthorization())

	return cmd
}

func CmdSetCodeAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-code-authorization [code-id] [methods]",
		Short: "Creates or updates the gasless operation authorization for the given code id and for the provided methods",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			codeID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			methods := strings.Split(args[1], ",")

			msg := types.NewMsgSetCodeAuthorization(
				clientCtx.GetFromAddress().String(),
				codeID,
				methods,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveCodeAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-code-authorization [code-id]",
		Short: "Removes any previously set code authorizations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			codeID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveCodeAuthorization(
				clientCtx.GetFromAddress().String(),
				codeID,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSetContractAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-contract-authorization [contract-address] [methods]",
		Short: "Creates or updates the gasless operation authorization for the given contract address and for the provided methods",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			methods := strings.Split(args[1], ",")

			msg := types.NewMsgSetContractAuthorization(
				clientCtx.GetFromAddress().String(),
				args[0],
				methods,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveContractAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-contract-authorization [contract-address]",
		Short: "Removes any previously set contract authorizations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveContractAuthorization(
				clientCtx.GetFromAddress().String(),
				args[0],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
