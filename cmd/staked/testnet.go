package main

// DONTCOVER

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmconfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/public-awesome/stakebird/app"
	"github.com/public-awesome/stakebird/x/curating"
)

var (
	flagNodeDirPrefix        = "node-dir-prefix"
	flagNumValidators        = "v"
	flagOutputDir            = "output-dir"
	flagNodeDaemonHome       = "node-daemon-home"
	flagNodeCLIHome          = "node-cli-home"
	flagStartingIPAddress    = "starting-ip-address"
	flagInitialCoins         = "coins"
	flagInitialStakingAmount = "initial-staking-amount"
	flagCurationWindow       = "curation-window"
	defaultKeyringBackend    = "test"
)

// get cmd to initialize all files for tendermint testnet and application
func testnetCmd(ctx *server.Context, cdc *codec.Codec,
	mbm module.BasicManager, genBalIterator bank.GenesisBalancesIterator,
) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Initialize files for a staked testnet",
		Long: `testnet will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).
Note, strict routability for addresses is turned off in the config file.
Example:
	staked testnet --v 4 --output-dir ./output --starting-ip-address 192.168.10.2
	`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			config := ctx.Config

			outputDir := viper.GetString(flagOutputDir)
			chainID := viper.GetString(flags.FlagChainID)
			minGasPrices := viper.GetString(server.FlagMinGasPrices)
			nodeDirPrefix := viper.GetString(flagNodeDirPrefix)
			nodeDaemonHome := viper.GetString(flagNodeDaemonHome)
			nodeCLIHome := viper.GetString(flagNodeCLIHome)
			startingIPAddress := viper.GetString(flagStartingIPAddress)
			numValidators := viper.GetInt(flagNumValidators)

			return InitTestnet(cmd, config, cdc, mbm, genBalIterator, outputDir, chainID,
				minGasPrices, nodeDirPrefix, nodeDaemonHome, nodeCLIHome, startingIPAddress, numValidators)
		},
	}

	cmd.Flags().Int(flagNumValidators, 4,
		"Number of validators to initialize the testnet with")
	cmd.Flags().StringP(flagOutputDir, "o", "./mytestnet",
		"Directory to store initialization data for the testnet")
	cmd.Flags().String(flagNodeDirPrefix, "node",
		"Prefix the directory name for each node with (node results in node0, node1, ...)")
	cmd.Flags().String(flagNodeDaemonHome, "staked",
		"Home directory of the node's daemon configuration")
	cmd.Flags().String(flagNodeCLIHome, "stakecli",
		"Home directory of the node's cli configuration")
	cmd.Flags().String(flagInitialCoins, fmt.Sprintf("1000000000%s", app.DefaultStakeDenom),
		"Validator genesis coins: 100000ustb,1000000uatom")
	cmd.Flags().Int64(flagInitialStakingAmount, 100000000,
		"Flag initial staking amount: 100000000")
	// nolint:lll
	cmd.Flags().String(flagStartingIPAddress, "192.168.0.1",
		"Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...)")
	cmd.Flags().String(
		flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	// nolint:lll
	cmd.Flags().String(
		server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", app.DefaultStakeDenom),
		"Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake)")
	cmd.Flags().String(flags.FlagKeyringBackend, defaultKeyringBackend, "Select keyring's backend (os|file|test)")
	cmd.Flags().String(flagStakeDenom, app.DefaultStakeDenom, "app's stake denom")
	cmd.Flags().String(flagUnbondingPeriod, app.DefaultUnbondingPeriod, "app's unbonding period")
	cmd.Flags().String(flagCurationWindow, "72h",
		"Curation Window for post expiration: 72h, 3h, 90m")

	return cmd
}

const nodeDirPerm = 0755

// TestnetNode holds configuration for nodes
type TestnetNode struct {
	Name             string
	OutsidePortRange string
	InsidePortRange  string
}

// InitTestnet the testnet
func InitTestnet(
	cmd *cobra.Command, config *tmconfig.Config, cdc *codec.Codec,
	mbm module.BasicManager, genBalIterator bank.GenesisBalancesIterator,
	outputDir, chainID, minGasPrices, nodeDirPrefix, nodeDaemonHome,
	nodeCLIHome, startingIPAddress string, numValidators int,
) error {

	if chainID == "" {
		chainID = "chain-" + tmrand.NewRand().Str(6)
	}

	monikers := make([]string, numValidators)
	nodeIDs := make([]string, numValidators)
	valPubKeys := make([]crypto.PubKey, numValidators)

	appConfig := srvconfig.DefaultConfig()
	appConfig.MinGasPrices = minGasPrices

	//nolint:prealloc
	var (
		genAccounts []authexported.GenesisAccount
		genBalances []bank.Balance
		genFiles    []string
	)

	inBuf := bufio.NewReader(cmd.InOrStdin())
	initialPort := 26656
	allocatedPorts := 2

	nodes := make([]TestnetNode, 0)
	// generate private keys, node IDs, and initial transactions
	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		clientDir := filepath.Join(outputDir, nodeDirName, nodeCLIHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")
		endPort := initialPort + allocatedPorts
		testnetNode := TestnetNode{
			Name:             nodeDirName,
			OutsidePortRange: fmt.Sprintf("%d-%d", initialPort, endPort),
			InsidePortRange:  fmt.Sprintf("%d-%d", 26656, 26656+allocatedPorts),
		}
		nodes = append(nodes, testnetNode)
		initialPort = endPort + 1
		config.SetRoot(nodeDir)
		config.RPC.ListenAddress = "tcp://0.0.0.0:26657"

		if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		if err := os.MkdirAll(clientDir, nodeDirPerm); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		monikers = append(monikers, nodeDirName)
		config.Moniker = nodeDirName
		nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(config)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}
		nodeIDs[i] = nodeID
		valPubKeys[i] = valPubKey
		memo := fmt.Sprintf("%s@%s:26656", nodeIDs[i], testnetNode.Name)
		genFiles = append(genFiles, config.GenesisFile())

		kb, err := keyring.New(
			sdk.KeyringServiceName(),
			viper.GetString(flags.FlagKeyringBackend),
			clientDir,
			inBuf,
		)
		if err != nil {
			return err
		}

		keyPass := clientkeys.DefaultKeyPass
		addr, secret, err := server.GenerateSaveCoinKey(kb, nodeDirName, keyPass, true)
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
		if err = writeFile(fmt.Sprintf("%v.json", "key_seed"), clientDir, cliPrint); err != nil {
			return err
		}
		stakeDenom := viper.GetString(flagStakeDenom)

		initialCoins := viper.GetString(flagInitialCoins)
		valCoins, err := sdk.ParseCoins(initialCoins)
		if err != nil {
			return err
		}

		genBalances = append(genBalances, bank.Balance{Address: addr, Coins: valCoins.Sort()})
		genAccounts = append(genAccounts, auth.NewBaseAccount(addr, nil, 0, 0))

		stakingAmount := viper.GetInt64(flagInitialStakingAmount)

		msg := staking.NewMsgCreateValidator(
			sdk.ValAddress(addr),
			valPubKeys[i],
			sdk.NewCoin(stakeDenom, sdk.NewInt(stakingAmount)),
			staking.NewDescription(nodeDirName, "", "", "", ""),
			staking.NewCommissionRates(sdk.OneDec(), sdk.OneDec(), sdk.OneDec()),
			sdk.OneInt(),
		)

		tx := auth.NewStdTx([]sdk.Msg{msg}, auth.StdFee{}, []auth.StdSignature{}, memo)
		txBldr := auth.NewTxBuilderFromCLI(inBuf).WithChainID(chainID).WithMemo(memo).WithKeybase(kb)

		signedTx, err := txBldr.SignStdTx(nodeDirName, clientkeys.DefaultKeyPass, tx, false)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		txBytes, err := cdc.MarshalJSON(signedTx)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		// gather gentxs folder
		if err := writeFile(fmt.Sprintf("%v.json", nodeDirName), gentxsDir, txBytes); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		// TODO: Rename config file to server.toml as it's not particular to Gaia
		// (REF: https://github.com/cosmos/cosmos-sdk/issues/4125).
		appConfigPath := filepath.Join(nodeDir, "config/app.toml")
		srvconfig.WriteConfigFile(appConfigPath, appConfig)
	}
	stakeDenom := viper.GetString(flagStakeDenom)
	unbondingPeriod := viper.GetString(flagUnbondingPeriod)
	if err := initGenFiles(
		cdc,
		mbm,
		chainID,
		stakeDenom,
		unbondingPeriod,
		genAccounts,
		genBalances,
		genFiles,
		numValidators,
	); err != nil {
		return err
	}

	err := collectGenFiles(
		cdc, config, chainID, monikers, nodeIDs, valPubKeys, numValidators,
		outputDir, nodeDirPrefix, nodeDaemonHome, genBalIterator,
	)
	if err != nil {
		return err
	}

	def, err := docker(nodes)
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

// nolint:interfacer
func initGenFiles(
	cdc *codec.Codec, mbm module.BasicManager, chainID, stakeDenom, unbondingPeriod string,
	genAccounts []authexported.GenesisAccount, genBalances []bank.Balance,
	genFiles []string, numValidators int,
) error {

	appGenState := mbm.DefaultGenesis(cdc)

	// set the accounts in the genesis state
	var authGenState auth.GenesisState
	cdc.MustUnmarshalJSON(appGenState[auth.ModuleName], &authGenState)

	authGenState.Accounts = genAccounts
	appGenState[auth.ModuleName] = cdc.MustMarshalJSON(authGenState)

	// set the balances in the genesis state
	var bankGenState bank.GenesisState
	cdc.MustUnmarshalJSON(appGenState[bank.ModuleName], &bankGenState)

	bankGenState.Balances = genBalances
	appGenState[bank.ModuleName] = cdc.MustMarshalJSON(bankGenState)

	// curating module
	curationWindow := viper.GetString(flagCurationWindow)
	curationWindowDuration, err := time.ParseDuration(curationWindow)
	if err != nil {
		return err
	}

	var curatingGenState curating.GenesisState
	cdc.MustUnmarshalJSON(appGenState[curating.ModuleName], &curatingGenState)
	curatingGenState.Params.CurationWindow = curationWindowDuration
	appGenState[curating.ModuleName] = cdc.MustMarshalJSON(curatingGenState)

	appGenStateJSON, err := codec.MarshalJSONIndent(cdc, appGenState)
	if err != nil {
		return err
	}

	genDoc := types.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}
	appState, err := initGenesis(cdc, &genDoc, stakeDenom, unbondingPeriod)
	if err != nil {
		return err
	}
	genDoc.AppState, err = cdc.MarshalJSON(appState)
	if err != nil {
		return err
	}

	// generate empty genesis files for each validator and save
	for i := 0; i < numValidators; i++ {
		if err := genDoc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

func collectGenFiles(
	cdc *codec.Codec, config *tmconfig.Config, chainID string,
	monikers, nodeIDs []string, valPubKeys []crypto.PubKey,
	numValidators int, outputDir, nodeDirPrefix, nodeDaemonHome string,
	genBalIterator bank.GenesisBalancesIterator,
) error {

	var appState json.RawMessage
	genTime := tmtime.Now()

	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")
		moniker := monikers[i]
		config.Moniker = nodeDirName

		config.SetRoot(nodeDir)

		nodeID, valPubKey := nodeIDs[i], valPubKeys[i]
		initCfg := genutil.NewInitConfig(chainID, gentxsDir, moniker, nodeID, valPubKey)

		genDoc, err := types.GenesisDocFromFile(config.GenesisFile())
		if err != nil {
			return err
		}

		nodeAppState, err := genutil.GenAppStateFromConfig(cdc, config, initCfg, *genDoc, genBalIterator)
		if err != nil {
			return err
		}

		if appState == nil {
			// set the canonical application state (they should not differ)
			appState = nodeAppState
		}

		genFile := config.GenesisFile()

		// overwrite each validator's genesis file to have a canonical genesis time
		if err := genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	err := tmos.EnsureDir(writePath, 0700)
	if err != nil {
		return err
	}

	err = tmos.WriteFile(file, contents, 0600)
	if err != nil {
		return err
	}

	return nil
}

const dockerComposeDefinition = `# Stakebird Testnet
version: '3.1'

services:{{range $node := .Nodes }}
	{{ $node.Name }}:
		image: publicawesome/stakebird
		ports:
			- {{ $node.OutsidePortRange}}:{{ $node.InsidePortRange}}
		volumes:
			- ./{{$node.Name}}/staked:/data/.staked/
{{end}}

	rest-server:
		image: publicawesome/stakebird
		ports:
			- 1317:1317
		command:
			- stakecli
			- rest-server
			- --laddr
			- tcp://:1317
			- --node
			- tcp://{{ (index .Nodes 0).Name }}:26657
			- --trust-node

`

func docker(nodes []TestnetNode) (string, error) {
	def := strings.ReplaceAll(dockerComposeDefinition, "\t", "  ")
	t, err := template.New("definition").Parse(def)
	if err != nil {
		return "", err
	}
	d := struct {
		Nodes []TestnetNode
	}{Nodes: nodes}

	buf := bytes.NewBufferString("")
	err = t.Execute(buf, d)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
