package cmd

import (
	"errors"
	"io"
	"os"

	"cosmossdk.io/log"
	tmcfg "github.com/cometbft/cometbft/config"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	confixcmd "cosmossdk.io/tools/confix/cmd"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/public-awesome/stargaze/v13/app"
	"github.com/public-awesome/stargaze/v13/app/params"
)

const EnvironmentPrefix = "STARGAZE"

// NewRootCmd creates a new root command for wasmd. It is called once in the
// main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(app.Bech32PrefixAccAddr, app.Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(app.Bech32PrefixValAddr, app.Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(app.Bech32PrefixConsAddr, app.Bech32PrefixConsPub)
	cfg.SetAddressVerifier(wasmtypes.VerifyAddressLen())
	cfg.Seal()

	tempApp := app.NewStargazeApp(log.NewNopLogger(), dbm.NewMemDB(), nil, true, simtestutil.NewAppOptionsWithFlagHome(tempDir()), []wasmkeeper.Option{})
	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: tempApp.InterfaceRegistry(),
		Codec:             tempApp.AppCodec(),
		TxConfig:          tempApp.TxConfig(),
		Amino:             tempApp.LegacyAmino(),
	}

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome).
		WithViper(EnvironmentPrefix)

	rootCmd := &cobra.Command{
		Use:   version.AppName,
		Short: "Stargaze App",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())
			// if the command is version, we don't need to initialize the client context
			if cmd.Name() == "version" {
				return nil
			}
			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}
			customTemplate, customParams := params.DefaultConfig()
			tmconfig := tmcfg.DefaultConfig()
			return server.InterceptConfigsPreRunHandler(cmd, customTemplate, customParams, tmconfig)
		},
	}

	initRootCmd(rootCmd, encodingConfig.TxConfig, encodingConfig.InterfaceRegistry, encodingConfig.Codec, app.ModuleBasics)

	return rootCmd, encodingConfig
}

func initRootCmd(
	rootCmd *cobra.Command,
	txConfig client.TxConfig,
	interfaceRegistry codectypes.InterfaceRegistry,
	appCodec codec.Codec,
	basicManager module.BasicManager,
) {
	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		debug.Cmd(),
		confixcmd.ConfigCommand(),
		Bech32Cmd(),
		//pruning.Cmd(newApp, app.DefaultNodeHome),
	//	snapshot.Cmd(newApp),
	)
	server.AddCommands(rootCmd, app.DefaultNodeHome, newApp, appExport, addModuleInitFlags)

	// add keybase, auxiliary RPC, query, genesis, and tx child commands
	rootCmd.AddCommand(
		server.StatusCommand(),
		genutilcli.Commands(txConfig, app.ModuleBasics, app.DefaultNodeHome),
		queryCommand(),
		txCommand(),
		keys.Commands(),
	)

}

// func oldInitRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
// 	rootCmd.AddCommand(
// 		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
// 		tmcli.NewCompletionCmd(rootCmd, true),
// 		debug.Cmd(),
// 		config.Cmd(),
// 		Bech32Cmd(),
// 		pruning.PruningCmd(newApp),
// 		snapshot.Cmd(newApp),
// 	)

// 	server.AddCommands(rootCmd, app.DefaultNodeHome, newApp, appExport, addModuleInitFlags)

// 	wasmcli.ExtendUnsafeResetAllCmd(rootCmd)

// 	// add keybase, auxiliary RPC, query, and tx child commands
// 	rootCmd.AddCommand(
// 		server.StatusCommand(),
// 		genutilcli.GenesisCoreCommand(encodingConfig.TxConfig, app.ModuleBasics, app.DefaultNodeHome),
// 		queryCommand(),
// 		txCommand(),
// 		keys.Commands(),
// 	)
// }

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
	wasm.AddModuleInitFlags(startCmd)
	startCmd.PreRunE = chainPreRuns(CheckLibwasmVersion, startCmd.PreRunE)
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		rpc.QueryEventForTxCmd(),
		server.QueryBlockCmd(),
		server.QueryBlockResultsCmd(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		authcmd.GetSimulateCmd(),
	)

	return cmd
}

func newApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	baseappOptions := server.DefaultBaseappOptions(appOpts)

	var wasmOpts []wasmkeeper.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	return app.NewStargazeApp(
		logger, db, traceStore, true,
		appOpts,
		wasmOpts,
		baseappOptions...,
	)

}

func appExport(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (servertypes.ExportedApp, error) {
	var stargazeApp *app.App
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home is not set")
	}

	loadLatest := height == -1

	var emptyWasmOpts []wasmkeeper.Option
	stargazeApp = app.NewStargazeApp(
		logger,
		db,
		traceStore,
		loadLatest,
		appOpts,
		emptyWasmOpts,
	)

	if height != -1 {
		if err := stargazeApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return stargazeApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, modulesToExport)
}

var tempDir = func() string {
	dir, err := os.MkdirTemp("", "starsd")
	if err != nil {
		panic("failed to create temp dir: " + err.Error())
	}
	defer os.RemoveAll(dir)

	return dir
}
