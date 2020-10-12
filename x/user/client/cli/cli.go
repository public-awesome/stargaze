package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the query commands for user module
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "user",
		Short:                      "User query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand()

	return queryCmd
}

// NewTxCmd returns the transaction commands for the user module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "user",
		Short:                      "User transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand()

	return txCmd
}
