package cli

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/public-awesome/stakebird/x/faucet/internal/types"
	"github.com/spf13/cobra"
)

// NewMintCmd returns the mint command
func NewMintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [denom]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Mint [denom] coin to sender andress",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Mint
Example:
$ %s tx faucet mint ustb --from address
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			denom := strings.TrimSpace(args[0])
			sender := clientCtx.GetFromAddress()

			if sender.Empty() {
				return fmt.Errorf("must provide a --from key")
			}
			msg := types.NewMsgMint(sender, sender, time.Now().Unix(), denom)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewPublishCmd returns the publish command
func NewPublishCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publish",
		Args:  cobra.MinimumNArgs(0),
		Short: "Publish current account as an public faucet. Do NOT add many coins in this account",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Publish
Example:
$ %s tx faucet publish --from faucet_key
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			buf := bufio.NewReader(cmd.InOrStdin())
			encryptPassword, err := input.GetPassword("Enter passphrase to encrypt the exported key:", buf)
			if err != nil {
				return err
			}
			armor, err := clientCtx.Keyring.ExportPrivKeyArmor(clientCtx.GetFromName(), encryptPassword)
			if err != nil {
				return err
			}

			msg := types.NewMsgFaucetKey(clientCtx.GetFromAddress(), armor)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewLoadKeyCmd returns the load key command
func NewLoadKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "load-key",
		Args:  cobra.MinimumNArgs(0),
		Short: "Loads the faucet key from chain and adds it to the local keyring for minting coins",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Load
Example:
$ %s tx faucet load-key
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryFaucetKeyRequest{}

			res, err := queryClient.FaucetKey(context.Background(), req)
			if err != nil {
				return err
			}
			buf := bufio.NewReader(cmd.InOrStdin())
			decryptPassword, err := input.GetPassword("Enter passphrase to decrypt faucet key:", buf)
			if err != nil {
				return err
			}
			err = clientCtx.Keyring.ImportPrivKey("faucet", res.FaucetKey.Armor, decryptPassword)
			if err != nil {
				return err
			}
			return clientCtx.PrintString("Faucet key loaded\n")
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
