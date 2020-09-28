package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

// GetQueryCmd returns the query commands for posts and upvotes
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "curating",
		Short:                      "Curating query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand(
		NewPostQueryCmd(),
		NewUpvotesQueryCmd(),
	)

	return queryCmd
}

// NewTxCmd returns the transaction commands for the curation module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "curating",
		Short:                      "Curating transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewPostTxCmd(),
		NewUpvoteTxCmd(),
	)

	return txCmd
}
