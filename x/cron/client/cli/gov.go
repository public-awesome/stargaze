package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/public-awesome/stargaze/v12/x/cron/types"
)

func ProposalSetPrivilegeContractCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "promote-to-privilege-contract [contract_addr_bech32]",
		Short:   "Create a proposal to promote the given contract",
		Long:    "Create a proposal to promote the given contract to privilege status which enables the contract to hook on to abci.BeginBlocker and abci.EndBlocker",
		Aliases: []string{"promote-contract"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, proposalTitle, proposalDescr, deposit, err := getProposalInfo(cmd)
			if err != nil {
				return err
			}

			contract := args[0]

			content := types.PromoteToPrivilegedContractProposal{
				Title:       proposalTitle,
				Description: proposalDescr,
				Contract:    contract,
			}

			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
		SilenceUsage: true,
	}

	// proposal flagsExecute
	cmd.Flags().String(cli.FlagTitle, "", "Title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "Description of proposal") //nolint:staticcheck
	cmd.Flags().String(cli.FlagDeposit, "", "Deposit of proposal")
	return cmd
}

func ProposalUnsetPrivilegeContractCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "demote-from-privilege-contract [contract_addr_bech32]",
		Short:   "Create a proposal to demote the given contract",
		Long:    "Create a proposal to demote the given contract privilege status which disables the contract to be called from abci.BeginBlocker and abci.EndBlocker",
		Aliases: []string{"demote-contract"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, proposalTitle, proposalDescr, deposit, err := getProposalInfo(cmd)
			if err != nil {
				return err
			}

			contract := args[0]

			content := types.DemotePrivilegedContractProposal{
				Title:       proposalTitle,
				Description: proposalDescr,
				Contract:    contract,
			}

			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
		SilenceUsage: true,
	}

	// proposal flagsExecute
	cmd.Flags().String(cli.FlagTitle, "", "Title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "Description of proposal") //nolint:staticcheck
	cmd.Flags().String(cli.FlagDeposit, "", "Deposit of proposal")
	return cmd
}

func getProposalInfo(cmd *cobra.Command) (client.Context, string, string, sdk.Coins, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return client.Context{}, "", "", nil, err
	}

	proposalTitle, err := cmd.Flags().GetString(cli.FlagTitle)
	if err != nil {
		return clientCtx, proposalTitle, "", nil, err
	}

	proposalDescr, err := cmd.Flags().GetString(cli.FlagDescription) //nolint:staticcheck
	if err != nil {
		return client.Context{}, proposalTitle, proposalDescr, nil, err
	}

	depositArg, err := cmd.Flags().GetString(cli.FlagDeposit)
	if err != nil {
		return client.Context{}, proposalTitle, proposalDescr, nil, err
	}

	deposit, err := sdk.ParseCoinsNormalized(depositArg)
	if err != nil {
		return client.Context{}, proposalTitle, proposalDescr, deposit, err
	}

	return clientCtx, proposalTitle, proposalDescr, deposit, nil
}
