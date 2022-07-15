package cmd

// DONTCOVER

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/spf13/cobra"
	tmconfig "github.com/tendermint/tendermint/config"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	flagNodeDirPrefix        = "node-dir-prefix"
	flagNumValidators        = "v"
	flagOutputDir            = "output-dir"
	flagNodeDaemonHome       = "node-daemon-home"
	flagStartingIPAddress    = "starting-ip-address"
	flagDockerTag            = "docker-tag"
	flagStakeDenom           = "stake-denom"
	flagUnbondingPeriod      = "unbonding-period"
	flagInitialCoins         = "coins"
	flagInitialStakingAmount = "initial-staking-amount"
	flagSnapshotFile         = "snapshot-file"
	defaultBondDenom         = "ustars"
)

// get cmd to initialize all files for tendermint testnet and application
func TestnetCmd(mbm module.BasicManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Initialize files for a simapp testnet",
		Long: `testnet will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Example:
	gaiad testnet --v 4 --output-dir ./output --starting-ip-address 192.168.10.2
	`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			outputDir, _ := cmd.Flags().GetString(flagOutputDir)
			keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			minGasPrices, _ := cmd.Flags().GetString(server.FlagMinGasPrices)
			nodeDirPrefix, _ := cmd.Flags().GetString(flagNodeDirPrefix)
			nodeDaemonHome, _ := cmd.Flags().GetString(flagNodeDaemonHome)
			startingIPAddress, _ := cmd.Flags().GetString(flagStartingIPAddress)
			numValidators, _ := cmd.Flags().GetInt(flagNumValidators)
			algo, _ := cmd.Flags().GetString(flags.FlagKeyAlgorithm)

			return InitTestnet(
				clientCtx, cmd, config, mbm, banktypes.GenesisBalancesIterator{}, outputDir, chainID, minGasPrices,
				nodeDirPrefix, nodeDaemonHome, startingIPAddress, keyringBackend, algo, numValidators,
			)
		},
	}

	cmd.Flags().Int(flagNumValidators, 4, "Number of validators to initialize the testnet with")
	cmd.Flags().StringP(flagOutputDir, "o", "./mytestnet", "Directory to store initialization data for the testnet")
	cmd.Flags().String(flagNodeDirPrefix, "node", "Prefix the directory name for each node with (node results in node0, node1, ...)")
	cmd.Flags().String(flagNodeDaemonHome, "starsd", "Home directory of the node's daemon configuration")
	cmd.Flags().String(flagStartingIPAddress, "192.168.0.1", "Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...)")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(server.FlagMinGasPrices, fmt.Sprintf("0%s", defaultBondDenom), "Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake)")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	cmd.Flags().String(flags.FlagKeyAlgorithm, string(hd.Secp256k1Type), "Key signing algorithm to generate keys for")
	cmd.Flags().String(flagStakeDenom, defaultBondDenom, "app's stake denom")
	cmd.Flags().Int64(flagInitialStakingAmount, 100000000,
		"Flag initial staking amount: 100000000")
	cmd.Flags().String(flagInitialCoins, fmt.Sprintf("1000000000%s", defaultBondDenom),
		"Validator genesis coins: 100000ustars")
	cmd.Flags().String(flagDockerTag, "latest", "docker tag for testnet command")
	cmd.Flags().String(flagUnbondingPeriod, "72h", "app's unbonding period")
	cmd.Flags().String(flagSnapshotFile, "", "snapshot filename for airdrop")

	return cmd
}

const nodeDirPerm = 0o755

func getSimappConfig(chainID, gasPrices string, statesync bool) *srvconfig.Config {
	simappConfig := srvconfig.DefaultConfig()
	simappConfig.MinGasPrices = gasPrices
	simappConfig.PruningKeepEvery = "0"
	simappConfig.PruningInterval = "11"
	simappConfig.PruningKeepRecent = "5"
	if statesync {
		simappConfig.PruningKeepEvery = "2000"
		simappConfig.StateSync.SnapshotInterval = 2000
		simappConfig.StateSync.SnapshotKeepRecent = 2
	}

	simappConfig.API.Enable = true
	simappConfig.API.EnableUnsafeCORS = true
	simappConfig.API.Swagger = true
	simappConfig.GRPCWeb.EnableUnsafeCORS = true
	simappConfig.Telemetry.Enabled = true
	simappConfig.Telemetry.PrometheusRetentionTime = 60
	simappConfig.Telemetry.EnableHostnameLabel = false
	simappConfig.Telemetry.GlobalLabels = [][]string{{"chain_id", chainID}}
	return simappConfig
}

// Initialize the testnet
func InitTestnet(
	clientCtx client.Context,
	cmd *cobra.Command,
	nodeConfig *tmconfig.Config,
	mbm module.BasicManager,
	genBalIterator banktypes.GenesisBalancesIterator,
	outputDir,
	chainID,
	minGasPrices,
	nodeDirPrefix,
	nodeDaemonHome,
	startingIPAddress,
	keyringBackend,
	algoStr string,
	numValidators int,
) error {
	if chainID == "" {
		chainID = "chain-" + tmrand.NewRand().Str(6)
	}

	nodeIDs := make([]string, numValidators)
	valPubKeys := make([]cryptotypes.PubKey, numValidators)

	var (
		genAccounts []authtypes.GenesisAccount
		genBalances []banktypes.Balance
		genFiles    []string
	)

	inBuf := bufio.NewReader(cmd.InOrStdin())

	initialPort := 26656
	allocatedPorts := 4
	nodes := make([]TestnetNode, 0)
	stakeDenom, err := cmd.Flags().GetString(flagStakeDenom)
	if err != nil {
		return err
	}

	initialCoins, err := cmd.Flags().GetString(flagInitialCoins)
	if err != nil {
		return err
	}
	valCoins, err := sdk.ParseCoinsNormalized(initialCoins)
	if err != nil {
		return err
	}

	// generate private keys, node IDs, and initial transactions
	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")

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

		nodeConfig.SetRoot(nodeDir)
		nodeConfig.RPC.CORSAllowedOrigins = []string{"*"}
		nodeConfig.RPC.ListenAddress = "tcp://0.0.0.0:26657"

		if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		nodeConfig.Moniker = nodeDirName
		var err error
		nodeIDs[i], valPubKeys[i], err = genutil.InitializeNodeValidatorFiles(nodeConfig)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		memo := fmt.Sprintf("%s@%s:26656", nodeIDs[i], testnetNode.Name)
		genFiles = append(genFiles, nodeConfig.GenesisFile())

		kb, err := keyring.New(sdk.KeyringServiceName(), keyringBackend, nodeDir, inBuf)
		if err != nil {
			return err
		}

		keyringAlgos, _ := kb.SupportedAlgorithms()
		algo, err := keyring.NewSigningAlgoFromString(algoStr, keyringAlgos)
		if err != nil {
			return err
		}

		addr, secret, err := testutil.GenerateSaveCoinKey(kb, nodeDirName, "", true, algo)
		if err != nil {
			_ = os.RemoveAll(outputDir)
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

		genBalances = append(genBalances, banktypes.Balance{Address: addr.String(), Coins: valCoins.Sort()})
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(addr, nil, 0, 0))

		stakingAmount, err := cmd.Flags().GetInt64(flagInitialStakingAmount)
		if err != nil {
			return err
		}
		createValMsg, err := stakingtypes.NewMsgCreateValidator(
			sdk.ValAddress(addr),
			valPubKeys[i],
			sdk.NewCoin(stakeDenom, sdk.NewInt(stakingAmount)),
			stakingtypes.NewDescription(nodeDirName, "", "", "", ""),
			stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 2), sdk.NewDecWithPrec(25, 2), sdk.NewDecWithPrec(5, 2)),
			sdk.OneInt(),
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
			WithChainID(chainID).
			WithMemo(memo).
			WithKeybase(kb).
			WithTxConfig(clientCtx.TxConfig)

		if err := tx.Sign(txFactory, nodeDirName, txBuilder, true); err != nil {
			return err
		}

		txBz, err := clientCtx.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
		if err != nil {
			return err
		}

		if err := writeFile(fmt.Sprintf("%v.json", nodeDirName), gentxsDir, txBz); err != nil {
			return err
		}

		simappConfig := getSimappConfig(chainID, minGasPrices, i == 0)
		srvconfig.WriteConfigFile(filepath.Join(nodeDir, "config/app.toml"), simappConfig)
	}
	for i, node := range nodes {
		fmt.Printf("Node %s, id [%s] outside port: %s \n", node.Name, nodeIDs[i], node.OutsidePortRange)
	}

	genAccounts = authtypes.SanitizeGenesisAccounts(genAccounts)
	genBalances = banktypes.SanitizeGenesisBalances(genBalances)

	snapshotFile, err := cmd.Flags().GetString(flagSnapshotFile)
	if err != nil {
		return err
	}
	if err := initGenFiles(clientCtx, mbm, chainID, snapshotFile, genAccounts, genBalances, genFiles, numValidators); err != nil {
		return err
	}

	err = collectGenFiles(
		clientCtx, nodeConfig, chainID, nodeIDs, valPubKeys, numValidators,
		outputDir, nodeDirPrefix, nodeDaemonHome, genBalIterator,
	)
	if err != nil {
		return err
	}
	dockerTag, err := cmd.Flags().GetString(flagDockerTag)
	if err != nil {
		return err
	}
	def, err := docker(nodes, dockerTag)
	if err != nil {
		return err
	}

	err = writeFile("docker-compose.yml", outputDir, []byte(def))

	if err != nil {
		return err
	}

	cmd.PrintErrf("Successfully initialized %d node directories\n", numValidators)
	return nil
}

func initGenFiles(
	clientCtx client.Context, mbm module.BasicManager, chainID, snapshotFileName string,
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

	bankGenState.Balances = genBalances
	appGenState[banktypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&bankGenState)

	appGenStateJSON, err := json.MarshalIndent(appGenState, "", "  ")
	if err != nil {
		return err
	}

	genDoc := &types.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}

	snapshot := &Snapshot{}
	if snapshotFileName != "" {
		// read snapshot.json and parse into struct
		snapshotFile, err := ioutil.ReadFile(filepath.Clean(snapshotFileName))
		if err != nil {
			return err
		}
		err = json.Unmarshal(snapshotFile, snapshot)
		if err != nil {
			return err
		}
	}

	appGenState, genDoc, err = PrepareGenesis(clientCtx, appGenState, genDoc, TestnetGenesisParams(), chainID, *snapshot)
	if err != nil {
		return err
	}
	// validate genesis state
	if err = mbm.ValidateGenesis(clientCtx.Codec, clientCtx.TxConfig, appGenState); err != nil {
		return fmt.Errorf("error validating genesis file: %s", err.Error())
	}

	// save genesis
	appStateJSON, err := json.Marshal(appGenState)
	if err != nil {
		return fmt.Errorf("failed to marshal application genesis state: %w", err)
	}

	genDoc.AppState = appStateJSON
	// generate empty genesis files for each validator and save
	for i := 0; i < numValidators; i++ {
		if err := genDoc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

func collectGenFiles(
	clientCtx client.Context, nodeConfig *tmconfig.Config, chainID string,
	nodeIDs []string, valPubKeys []cryptotypes.PubKey, numValidators int,
	outputDir, nodeDirPrefix, nodeDaemonHome string, genBalIterator banktypes.GenesisBalancesIterator,
) error {
	var appState json.RawMessage
	genTime := tmtime.Now()

	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")
		nodeConfig.Moniker = nodeDirName

		nodeConfig.SetRoot(nodeDir)

		nodeID, valPubKey := nodeIDs[i], valPubKeys[i]
		initCfg := genutiltypes.NewInitConfig(chainID, gentxsDir, nodeID, valPubKey)

		genDoc, err := types.GenesisDocFromFile(nodeConfig.GenesisFile())
		if err != nil {
			return err
		}

		nodeAppState, err := genutil.GenAppStateFromConfig(clientCtx.Codec, clientCtx.TxConfig, nodeConfig, initCfg, *genDoc, genBalIterator)
		if err != nil {
			return err
		}

		if appState == nil {
			// set the canonical application state (they should not differ)
			appState = nodeAppState
		}

		genFile := nodeConfig.GenesisFile()

		// overwrite each validator's genesis file to have a canonical genesis time
		if err := genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(name string, dir string, contents []byte) error {
	file := filepath.Join(dir, name)
	err := tmos.EnsureDir(dir, 0o755)
	if err != nil {
		return err
	}

	err = tmos.WriteFile(file, contents, 0o644)
	if err != nil {
		return err
	}

	return nil
}

// TestnetNode holds configuration for nodes
type TestnetNode struct {
	Name             string
	OutsidePortRange string
	InsidePortRange  string
	APIPort          string
	GRPCPort         string
}

const dockerComposeDefinition = `# Stargaze Testnet
version: '3.1'
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
