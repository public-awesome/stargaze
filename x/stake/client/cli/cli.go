package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the query commands for posts and upvotes
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "stake",
		Short:                      "Stake query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand(
		NewStakesQueryCmd(),
		NewStakeQueryCmd(),
	)

	return queryCmd
}

// NewTxCmd returns the transaction commands for the stake module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "stake",
		Short:                      "Stake transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewStakeTxCmd(),
		NewUnstakeTxCmd(),
		NewBuyCreatorCoinTxCmd(),
		NewSellCreatorCoinTxCmd(),
	)

	return txCmd
}
