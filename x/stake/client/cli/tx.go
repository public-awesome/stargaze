package cli

// // GetTxCmd returns the transaction commands for this module
// func GetTxCmd(cdc *codec.Codec) *cobra.Command {
// 	stakeTxCmd := &cobra.Command{
// 		Use:                        types.ModuleName,
// 		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
// 		DisableFlagParsing:         true,
// 		SuggestionsMinimumDistance: 2,
// 		RunE:                       client.ValidateCmd,
// 	}

// 	stakeTxCmd.AddCommand(flags.PostCommands(
// 		GetCmdDelegate(cdc),
// 	)...)

// 	return stakeTxCmd
// }

// // GetCmdDelegate implements the delegate command.
// func GetCmdDelegate(cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "delegate [validator-addr] [amount]",
// 		Args:  cobra.ExactArgs(2),
// 		Short: "Delegate liquid tokens to a validator",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`Delegate an amount of liquid coins to a validator from your wallet.
// Example:
// $ %s tx stake delegate cosmosvaloper1l2rsakp388kuv9k8qzq6lrm9taddae7fpx59wm 1000stake --from mykey
// `,
// 				version.ClientName,
// 			),
// 		),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(auth.DefaultTxEncoder(cdc))
// 			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

// 			// amount, err := sdk.ParseCoin(args[1])
// 			// if err != nil {
// 			// 	return err
// 			// }

// 			// delAddr := cliCtx.GetFromAddress()
// 			// valAddr, err := sdk.ValAddressFromBech32(args[0])
// 			// if err != nil {
// 			// 	return err
// 			// }

// 			// msg := types.NewMsgDelegate(delAddr, valAddr, amount)
// 			// return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
// 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{})
// 		},
// 	}
// }
