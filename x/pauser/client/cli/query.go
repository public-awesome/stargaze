package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/public-awesome/stargaze/v18/x/pauser/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd builds query command group for the module.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetCmdParams())
	cmd.AddCommand(GetCmdIsContractPaused())
	cmd.AddCommand(GetCmdIsCodeIDPaused())
	cmd.AddCommand(GetCmdPausedContracts())
	cmd.AddCommand(GetCmdPausedCodeIDs())

	return cmd
}

func GetCmdParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "List the module params",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(
				cmd.Context(),
				&types.QueryParamsRequest{},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdIsContractPaused() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-paused [contract-address]",
		Short: "Check if a contract is paused",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.IsContractPaused(
				cmd.Context(),
				&types.QueryIsContractPausedRequest{
					ContractAddress: args[0],
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdIsCodeIDPaused() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-code-paused [code-id]",
		Short: "Check if a code ID is paused",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			codeID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.IsCodeIDPaused(
				cmd.Context(),
				&types.QueryIsCodeIDPausedRequest{
					CodeId: codeID,
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdPausedContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "paused-contracts",
		Short: "List all paused contracts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PausedContracts(
				cmd.Context(),
				&types.QueryPausedContractsRequest{},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdPausedCodeIDs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "paused-code-ids",
		Short: "List all paused code IDs",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PausedCodeIDs(
				cmd.Context(),
				&types.QueryPausedCodeIDsRequest{},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
