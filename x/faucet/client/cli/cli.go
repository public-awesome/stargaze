package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/public-awesome/stakebird/x/faucet/internal/types"
)

// GetQueryCmd returns the query commands for posts and upvotes
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Faucet query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand()

	return queryCmd
}

// NewTxCmd returns the transaction commands for the curation module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Faucet transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewMintCmd(),
		NewPublishCmd(),
		NewLoadKeyCmd(),
	)

	return txCmd
}
