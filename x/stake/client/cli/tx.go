package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/public-awesome/stargaze/x/stake/types"
	"github.com/spf13/cobra"
)

// NewStakeTxCmd returns the stake command
func NewStakeTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake [vendor-id] [post-id] [amount] [validator-address] --from [key]",
		Args:  cobra.MinimumNArgs(4),
		Short: "Stake on a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Stake on a post.
Example:
$ %s tx stake stake 1 "2" 500 starsvaloper1deadbeef --from mykey
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			delegator := clientCtx.GetFromAddress()

			vendorID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			postID := args[1]
			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				panic("invalid amount")
			}

			valAddrStr := args[3]
			validator, err := sdk.ValAddressFromBech32(valAddrStr)
			if err != nil {
				return err
			}

			msg := types.NewMsgStake(
				uint32(vendorID), postID, delegator, validator, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewUnstakeTxCmd returns the unstake command
func NewUnstakeTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unstake [vendor-id] [post-id] [amount] --from [key]",
		Args:  cobra.MinimumNArgs(3),
		Short: "Unstake from a post",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Stake on a post.
Example:
$ %s tx stake unstake 1 "2" 500 --from mykey
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			delegator := clientCtx.GetFromAddress()

			vendorID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			postID := args[1]
			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				panic("invalid amount")
			}

			msg := types.NewMsgUnstake(uint32(vendorID), postID, delegator, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewBuyCreatorCoinTxCmd returns the stake command
//nolint:dupl
func NewBuyCreatorCoinTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy [username] [creatorAddr] [amount] [validator-address] --from [key]",
		Args:  cobra.MinimumNArgs(4),
		Short: "Buy a creator's coin",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Buy a creator's coin with the native token.
Example:
$ %s tx stake buy "satoshi" stars1deadbeef 21000000 starsvaloper1deadbeef --from mykey
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			buyer := clientCtx.GetFromAddress()

			username := args[0]
			if len(username) <= 3 {
				return errors.New("username too short")
			}

			creatorAddrStr := args[1]
			creator, err := sdk.AccAddressFromBech32(creatorAddrStr)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("invalid amount, must be an int value")
			}

			valAddrStr := args[3]
			validator, err := sdk.ValAddressFromBech32(valAddrStr)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyCreatorCoin(
				username, creator, buyer, validator, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewSellCreatorCoinTxCmd returns the stake command
//nolint:dupl
func NewSellCreatorCoinTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell [username] [creatorAddr] [amount] [validator-address] --from [key]",
		Args:  cobra.MinimumNArgs(4),
		Short: "Sell a creator's coin",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Sell a creator's coin with the native token.
Example:
$ %s tx stake sell "satoshi" stars1deadbeef 21000000 starsvaloper1deadbeef --from mykey
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			seller := clientCtx.GetFromAddress()

			username := args[0]
			if len(username) <= 3 {
				return errors.New("username too short")
			}

			creatorAddrStr := args[1]
			creator, err := sdk.AccAddressFromBech32(creatorAddrStr)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("invalid amount, must be an int value")
			}

			valAddrStr := args[3]
			validator, err := sdk.ValAddressFromBech32(valAddrStr)
			if err != nil {
				return err
			}

			msg := types.NewMsgSellCreatorCoin(
				username, creator, seller, validator, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
