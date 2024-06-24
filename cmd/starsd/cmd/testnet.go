package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"cosmossdk.io/math"
	"cosmossdk.io/math/unsafe"
	cmtcfg "github.com/cometbft/cometbft/config"
	cmttypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/public-awesome/stargaze/v15/app/params"
	"github.com/spf13/cobra"
)

var (
	testnetFlagNumValidators        = "v"
	testnetFlagOutputDir            = "output-dir"
	testnetFlagInitialCoins         = "coins"
	testnetFlagInitialStakingAmount = "initial-staking-amount"
	testnetFlagDockerTag            = "docker-tag"
	tesnetFlagStakeDenom            = "stake-denom"
	testnetDefaultDenom             = "ustars"
	testnetFlagUnbondingPeriod      = "unbonding-period"
)

func NewTestnetCmd(mbm module.BasicManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Initialize files for stargaze testnet",
		Long: `testnet will create a "v" num of directories with files belonging to validators.
This configuration is strictly for docker compose bootstrapping.
		`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			args := testnetArguments{}
			var err error
			args.outputDir, _ = cmd.Flags().GetString(testnetFlagOutputDir)
			args.chainID, _ = cmd.Flags().GetString(flags.FlagChainID)
			// args.minGasPrices, _ = cmd.Flags().GetString(server.FlagMinGasPrices)
			args.numValidators, _ = cmd.Flags().GetInt(testnetFlagNumValidators)
			args.stakeAmount, err = cmd.Flags().GetInt64(testnetFlagInitialStakingAmount)
			if err != nil {
				return err
			}
			args.stakeDenom, err = cmd.Flags().GetString(tesnetFlagStakeDenom)
			if err != nil {
				return err
			}

			args.validatorCoins, err = cmd.Flags().GetString(testnetFlagInitialCoins)
			if err != nil {
				return err
			}
			err = initTestnet(cmd, args, mbm)
			if err != nil {
				fmt.Println(err)
			}
			return err
		},
	}
	cmd.Flags().Int(testnetFlagNumValidators, 4, "Number of validators to initialize the testnet with")
	cmd.Flags().StringP(testnetFlagOutputDir, "o", "./stargaze-testnet", "Directory to store initialization data for the testnet")
	cmd.Flags().String(tesnetFlagStakeDenom, testnetDefaultDenom, "app's stake denom")
	cmd.Flags().String(testnetFlagDockerTag, "latest", "docker tag for testnet command")
	cmd.Flags().String(testnetFlagUnbondingPeriod, "72h", "app's unbonding period")
	cmd.Flags().Int64(testnetFlagInitialStakingAmount, 100000000,
		"Flag initial staking amount: 100000000")
	cmd.Flags().String(testnetFlagInitialCoins, fmt.Sprintf("1000000000%s", testnetDefaultDenom),
		"Validator genesis coins: 100000ustars")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")

	return cmd
}

// // NewTestNetworkFixture returns a new WasmApp AppConstructor for network simulation tests
// func NewTestNetworkFixture() network.TestFixture {
// 	dir, err := os.MkdirTemp("", "stargaze")
// 	if err != nil {
// 		panic(fmt.Sprintf("failed creating temporary directory: %v", err))
// 	}
// 	defer os.RemoveAll(dir)
// 	var emptyWasmOptions []wasmkeeper.Option
// 	app := stargazeapp.NewStargazeApp(log.NewNopLogger(), dbm.NewMemDB(), nil, true, simtestutil.NewAppOptionsWithFlagHome(dir), emptyWasmOptions)
// 	appCtr := func(val network.ValidatorI) servertypes.Application {
// 		return stargazeapp.NewStargazeApp(
// 			val.GetCtx().Logger, dbm.NewMemDB(), nil, true,
// 			simtestutil.NewAppOptionsWithFlagHome(val.GetCtx().Config.RootDir),
// 			emptyWasmOptions,
// 			bam.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
// 			bam.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
// 			bam.SetChainID(val.GetCtx().Viper.GetString(flags.FlagChainID)),
// 		)
// 	}

//		return network.TestFixture{
//			AppConstructor: appCtr,
//			// GenesisState:   app.DefaultGenesis(),
//			EncodingConfig: testutil.TestEncodingConfig{
//				InterfaceRegistry: app.InterfaceRegistry(),
//				Codec:             app.AppCodec(),
//				TxConfig:          app.TxConfig(),
//				Amino:             app.LegacyAmino(),
//			},
//		}
//	}
func getAppConfig(chainID, gasPrices string, statesync, fullNode bool) (string, params.CustomAppConfig) {
	template, customConfig := params.DefaultConfig()
	customAppConfig := customConfig.(params.CustomAppConfig)

	customAppConfig.MinGasPrices = gasPrices

	customAppConfig.Pruning = "everything"
	if fullNode {
		customAppConfig.Pruning = "nothing"
	}

	if statesync {
		customAppConfig.StateSync.SnapshotInterval = 2000
		customAppConfig.StateSync.SnapshotKeepRecent = 2
	}

	customAppConfig.GRPC.Enable = true
	customAppConfig.API.Enable = true
	customAppConfig.API.EnableUnsafeCORS = true
	customAppConfig.API.Swagger = true
	customAppConfig.GRPCWeb.Enable = true
	customAppConfig.Telemetry.Enabled = true
	customAppConfig.Telemetry.PrometheusRetentionTime = 60
	customAppConfig.Telemetry.EnableHostnameLabel = false
	customAppConfig.Telemetry.GlobalLabels = [][]string{{"chain_id", chainID}}

	return template, customAppConfig
}

type testnetArguments struct {
	chainID        string
	outputDir      string
	numValidators  int
	stakeDenom     string
	stakeAmount    int64
	validatorCoins string
}

// TestnetNode holds configuration for nodes
type TestnetNode struct {
	Name             string
	OutsidePortRange string
	InsidePortRange  string
	APIPort          string
	GRPCPort         string
}

const nodeDirPerm = 0o755

func initTestnet(cmd *cobra.Command, args testnetArguments, mbm module.BasicManager) error {
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return err
	}

	if args.chainID == "" {
		args.chainID = "chain-" + unsafe.Str(6)
	}

	var (
		genAccounts []authtypes.GenesisAccount
		genBalances []banktypes.Balance
		genFiles    []string
	)

	nodeIDs := make([]string, args.numValidators)
	valPubKeys := make([]cryptotypes.PubKey, args.numValidators)

	initialPort := 26656
	allocatedPorts := 4
	nodes := make([]TestnetNode, 0)

	for i := 0; i < args.numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", "node", i)
		nodeDir := filepath.Join(args.outputDir, nodeDirName, "starsd")
		gentxsDir := filepath.Join(args.outputDir, "gentxs")

		// for docker
		endPort := initialPort + allocatedPorts
		testnetNode := TestnetNode{
			Name:             nodeDirName,
			OutsidePortRange: fmt.Sprintf("%d-%d", initialPort, initialPort+2),
			InsidePortRange:  fmt.Sprintf("%d-%d", 26656, 26656+2),
			APIPort:          fmt.Sprintf("%d", initialPort+3),
			GRPCPort:         fmt.Sprintf("%d", initialPort+4),
		}
		nodes = append(nodes, testnetNode)
		initialPort = endPort + 1

		nodeConfig := cmtcfg.DefaultConfig()

		nodeConfig.SetRoot(nodeDir)
		nodeConfig.Moniker = nodeDirName
		nodeConfig.RPC.ListenAddress = "tcp://0.0.0.0:26657"
		nodeConfig.RPC.CORSAllowedOrigins = []string{"*"}

		if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
			if err := os.RemoveAll(args.outputDir); err != nil {
				return err
			}
			return err
		}

		nodeConfig.Moniker = nodeDirName
		var err error
		nodeIDs[i], valPubKeys[i], err = genutil.InitializeNodeValidatorFiles(nodeConfig)
		if err != nil {
			if err := os.RemoveAll(args.outputDir); err != nil {
				return err
			}
			return err
		}
		memo := fmt.Sprintf("%s@%s:26656", nodeIDs[i], testnetNode.Name)
		genFiles = append(genFiles, nodeConfig.GenesisFile())
		inBuf := bufio.NewReader(cmd.InOrStdin())
		kb, err := keyring.New(sdk.KeyringServiceName(), "test", nodeDir, inBuf, clientCtx.Codec)
		if err != nil {
			return err
		}

		keyringAlgos, _ := kb.SupportedAlgorithms()
		algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
		if err != nil {
			return err
		}

		addr, secret, err := testutil.GenerateSaveCoinKey(kb, nodeDirName, "", true, algo)
		if err != nil {
			_ = os.RemoveAll(args.outputDir)
			return err
		}

		info := map[string]string{"secret": secret}

		cliPrint, err := json.Marshal(info)
		if err != nil {
			return err
		}

		// save private key seed words
		if err := writeFile(fmt.Sprintf("%v.json", "key_seed"), nodeDir, cliPrint); err != nil {
			return err
		}

		valCoins, err := sdk.ParseCoinsNormalized(args.validatorCoins)
		if err != nil {
			return err
		}

		genBalances = append(genBalances, banktypes.Balance{Address: addr.String(), Coins: valCoins})
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(addr, nil, 0, 0))

		valStr, err := clientCtx.TxConfig.SigningContext().ValidatorAddressCodec().BytesToString(sdk.ValAddress(addr))
		if err != nil {
			return err
		}
		createValMsg, err := stakingtypes.NewMsgCreateValidator(
			valStr,
			valPubKeys[i],
			sdk.NewInt64Coin(args.stakeDenom, args.stakeAmount),
			stakingtypes.NewDescription(nodeDirName, "", "", "", ""),
			stakingtypes.NewCommissionRates(math.LegacyNewDecWithPrec(5, 2), math.LegacyNewDecWithPrec(25, 2), math.LegacyNewDecWithPrec(5, 2)),
			math.OneInt(),
		)
		if err != nil {
			return err
		}

		txBuilder := clientCtx.TxConfig.NewTxBuilder()
		if err := txBuilder.SetMsgs(createValMsg); err != nil {
			return err
		}

		txBuilder.SetMemo(memo)

		txFactory := tx.Factory{}
		txFactory = txFactory.
			WithChainID(args.chainID).
			WithMemo(memo).
			WithKeybase(kb).
			WithTxConfig(clientCtx.TxConfig)

		if err := tx.Sign(cmd.Context(), txFactory, nodeDirName, txBuilder, true); err != nil {
			return err
		}

		txBz, err := clientCtx.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
		if err != nil {
			return err
		}

		if err := writeFile(fmt.Sprintf("%v.json", nodeDirName), gentxsDir, txBz); err != nil {
			return err
		}

		// first node has statesync and full archival node
		template, config := getAppConfig(args.chainID, "0ustars", i == 0, i == 0)
		srvconfig.SetConfigTemplate(template)
		srvconfig.WriteConfigFile(filepath.Join(nodeDir, "config", "app.toml"), config)

	}
	for i, node := range nodes {
		fmt.Printf("Node %s, id [%s] outside port: %s \n", node.Name, nodeIDs[i], node.OutsidePortRange)
	}

	genAccounts = authtypes.SanitizeGenesisAccounts(genAccounts)
	genBalances = banktypes.SanitizeGenesisBalances(genBalances)
	if err := initGenFiles(clientCtx, mbm, args.chainID, genAccounts, genBalances, genFiles, args.numValidators); err != nil {
		return err
	}

	err = collectGenFiles(
		clientCtx, args.chainID, nodeIDs, valPubKeys, args.numValidators,
		args.outputDir, "node", "starsd",
	)
	if err != nil {
		return err
	}
	dockerTag, err := cmd.Flags().GetString(testnetFlagDockerTag)
	if err != nil {
		return err
	}
	def, err := docker(nodes, dockerTag)
	if err != nil {
		return err
	}

	err = writeFile("docker-compose.yml", args.outputDir, []byte(def))
	if err != nil {
		return err
	}

	cmd.PrintErrf("Successfully initialized %d node directories\n", args.numValidators)
	return nil
}

func collectGenFiles(
	clientCtx client.Context, chainID string,
	nodeIDs []string, valPubKeys []cryptotypes.PubKey, numValidators int,
	outputDir, nodeDirPrefix, nodeDaemonHome string,
) error {
	genBalIterator := banktypes.GenesisBalancesIterator{}
	genTime := time.Now()

	for i := 0; i < numValidators; i++ {
		nodeConfig := cmtcfg.DefaultConfig()
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")
		nodeConfig.Moniker = nodeDirName

		nodeConfig.RPC.ListenAddress = "tcp://0.0.0.0:26657"
		nodeConfig.RPC.CORSAllowedOrigins = []string{"*"}

		nodeConfig.SetRoot(nodeDir)

		nodeID, valPubKey := nodeIDs[i], valPubKeys[i]
		initCfg := genutiltypes.NewInitConfig(chainID, gentxsDir, nodeID, valPubKey)

		appGenesis, err := genutiltypes.AppGenesisFromFile(nodeConfig.GenesisFile())
		if err != nil {
			return err
		}

		_, err = genutil.GenAppStateFromConfig(clientCtx.Codec, clientCtx.TxConfig, nodeConfig, initCfg, appGenesis, genBalIterator, genutiltypes.DefaultMessageValidator,
			clientCtx.TxConfig.SigningContext().ValidatorAddressCodec())
		if err != nil {
			return err
		}

		appGenesis.GenesisTime = genTime
		genFile := nodeConfig.GenesisFile()

		if err := appGenesis.SaveAs(genFile); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(name, dir string, contents []byte) error {
	file := filepath.Join(dir, name)

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("could not create directory %q: %w", dir, err)
	}

	return os.WriteFile(file, contents, 0o600)
}

func initGenFiles(
	clientCtx client.Context, mbm module.BasicManager, chainID string,
	genAccounts []authtypes.GenesisAccount, genBalances []banktypes.Balance,
	genFiles []string, numValidators int,
) error {
	appGenState := mbm.DefaultGenesis(clientCtx.Codec)

	// set the accounts in the genesis state
	var authGenState authtypes.GenesisState
	clientCtx.Codec.MustUnmarshalJSON(appGenState[authtypes.ModuleName], &authGenState)

	accounts, err := authtypes.PackAccounts(genAccounts)
	if err != nil {
		return err
	}

	authGenState.Accounts = accounts
	appGenState[authtypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&authGenState)

	// set the balances in the genesis state
	var bankGenState banktypes.GenesisState
	clientCtx.Codec.MustUnmarshalJSON(appGenState[banktypes.ModuleName], &bankGenState)

	bankGenState.Balances = banktypes.SanitizeGenesisBalances(genBalances)
	for _, bal := range bankGenState.Balances {
		bankGenState.Supply = bankGenState.Supply.Add(bal.Coins...)
	}
	appGenState[banktypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&bankGenState)

	appGenState = PrepareGenesis(clientCtx, appGenState, TestnetGenesisParams())

	appGenStateJSON, err := json.MarshalIndent(appGenState, "", "  ")
	if err != nil {
		return err
	}

	appGenesis := genutiltypes.NewAppGenesisWithVersion(chainID, appGenStateJSON)
	consensusParams := cmttypes.DefaultConsensusParams()

	consensusParams.ABCI.VoteExtensionsEnableHeight = 2
	consensusParams.Block.MaxBytes = 10 * 1024 * 1024
	consensusParams.Block.MaxGas = 300_000_000
	appGenesis.Consensus = &genutiltypes.ConsensusGenesis{
		Validators: nil,
		Params:     consensusParams,
	}

	// generate empty genesis files for each validator and save
	for i := 0; i < numValidators; i++ {
		if err := appGenesis.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

const dockerComposeDefinition = `# Stargaze Testnet
version: '3.5'
services:{{range $node := .Nodes }}
	{{ $node.Name }}:
		image: publicawesome/stargaze:{{ $.Tag }}
		pull_policy: always
		restart: always
		ports:
			- {{ $node.OutsidePortRange}}:{{ $node.InsidePortRange}}
			- {{ $node.APIPort}}:1317
			- {{ $node.GRPCPort}}:9090
		volumes:
			- ./{{$node.Name}}/starsd:/data/.starsd/
{{end}}
`

func docker(nodes []TestnetNode, tag string) (string, error) {
	def := strings.ReplaceAll(dockerComposeDefinition, "\t", "  ")
	t, err := template.New("definition").Parse(def)
	if err != nil {
		return "", err
	}
	d := struct {
		Nodes []TestnetNode
		Tag   string
	}{Nodes: nodes, Tag: tag}

	buf := bytes.NewBufferString("")
	err = t.Execute(buf, d)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
