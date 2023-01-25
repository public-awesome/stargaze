package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v8/x/cron/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group cron queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetCmdListPrivilegedContracts())
	// this line is used by starport scaffolding # 1

	return cmd
}

// GetCmdListPrivilegedContracts lists all privileged contracts
func GetCmdListPrivilegedContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-privileged",
		Short:   "List all privileged contract addresses",
		Long:    "List all contract addresses which have been elevated to privileged status",
		Aliases: []string{"privileged-contracts", "privileged", "lpc"},
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ListPrivileged(
				cmd.Context(),
				&types.QueryListPrivilegedRequest{},
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
