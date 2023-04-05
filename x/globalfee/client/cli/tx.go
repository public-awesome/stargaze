package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
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
	// proposals for the same
	cmd.AddCommand(CmdProposalSetCodeAuthorization())
	cmd.AddCommand(CmdProposalRemoveCodeAuthorization())
	cmd.AddCommand(CmdProposalSetContractAuthorization())
	cmd.AddCommand(CmdProposalRemoveContractAuthorization())

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

			codeId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			methods := strings.Split(args[1], ",")

			msg := types.NewMsgSetCodeAuthorization(
				clientCtx.GetFromAddress().String(),
				codeId,
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

			codeId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveCodeAuthorization(
				clientCtx.GetFromAddress().String(),
				codeId,
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

func CmdProposalSetCodeAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-code-authorization-proposal [code-id] [methods]",
		Short: "Creates a proposal which creates or updates the gasless operation authorization for the given code id and for the provided methods",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, proposalTitle, proposalDescr, deposit, err := getProposalInfo(cmd)
			if err != nil {
				return err
			}

			codeId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			methods := strings.Split(args[1], ",")

			prop := types.SetCodeAuthorizationProposal{
				Title:       proposalTitle,
				Description: proposalDescr,
				CodeAuthorization: &types.CodeAuthorization{
					CodeId:  codeId,
					Methods: methods,
				},
			}

			msg, err := govtypes.NewMsgSubmitProposal(&prop, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(cli.FlagTitle, "", "Title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "Description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "Deposit of proposal")

	return cmd
}

func CmdProposalRemoveCodeAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-code-authorization-proposal [code-id]",
		Short: "Creates a proposal which removes any previously set code authorizations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, proposalTitle, proposalDescr, deposit, err := getProposalInfo(cmd)
			if err != nil {
				return err
			}

			codeId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			content := types.RemoveCodeAuthorizationProposal{
				Title:       proposalTitle,
				Description: proposalDescr,
				CodeId:      codeId,
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
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(cli.FlagTitle, "", "Title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "Description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "Deposit of proposal")

	return cmd
}

func CmdProposalSetContractAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-contract-authorization-proposal [contract-address] [methods]",
		Short: "Creates a proposal which creates or updates the gasless operation authorization for the given contract address and for the provided methods",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, proposalTitle, proposalDescr, deposit, err := getProposalInfo(cmd)
			if err != nil {
				return err
			}

			methods := strings.Split(args[1], ",")

			content := types.SetContractAuthorizationProposal{
				Title:       proposalTitle,
				Description: proposalDescr,
				ContractAuthorization: &types.ContractAuthorization{
					ContractAddress: args[0],
					Methods:         methods,
				},
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
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(cli.FlagTitle, "", "Title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "Description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "Deposit of proposal")

	return cmd
}

func CmdProposalRemoveContractAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-contract-authorization [contract-address]",
		Short: "Creates a proposal which removes any previously set contract authorizations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, proposalTitle, proposalDescr, deposit, err := getProposalInfo(cmd)
			if err != nil {
				return err
			}
			content := types.RemoveContractAuthorizationProposal{
				Title:           proposalTitle,
				Description:     proposalDescr,
				ContractAddress: args[0],
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
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(cli.FlagTitle, "", "Title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "Description of proposal")
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

	proposalDescr, err := cmd.Flags().GetString(cli.FlagDescription)
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
