package e2e

import (
	"context"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	interchaintest "github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
	"github.com/strangelove-ventures/interchaintest/v4/testreporter"
	"github.com/strangelove-ventures/interchaintest/v4/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"gopkg.in/yaml.v2"
)

func TestIBCHooks(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	var (
		ctx                = context.Background()
		client, network    = interchaintest.DockerSetup(t)
		rep                = testreporter.NewNopReporter()
		eRep               = rep.RelayerExecReporter(t)
		chainIdA, chainIdB = "chain-a", "chain-b"
		chainA, chainB     *cosmos.CosmosChain

		// Each network will contain 1 validator and 0 full nodes.
		// This is to keep overhead down so the tests do not eat up resources and take unnecessarily long to complete.
		numVals      = 1
		numFullNodes = 0
	)

	// Use the default Stargaze Chain Config and override the chain IDs for each network.
	baseCfg := stargazeCfg

	baseCfg.ChainID = chainIdA
	configA := baseCfg

	baseCfg.ChainID = chainIdB
	configB := baseCfg

	// Build our chain factory with 4 distinct chains
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          "stargaze",
			ChainConfig:   configA,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          "stargaze",
			ChainConfig:   configB,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	chainA, chainB = chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	r := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
	).Build(t, client, network)

	const pathAB = "ab"
	ic := interchaintest.NewInterchain().
		AddChain(chainA).
		AddChain(chainB).
		AddRelayer(r, "relayer").
		AddLink(interchaintest.InterchainLink{
			Chain1:  chainA,
			Chain2:  chainB,
			Relayer: r,
			Path:    pathAB,
		})

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),

		SkipPathCreation: false,
	}))
	t.Cleanup(func() {
		_ = ic.Close()
	})

	const userFunds = int64(10_000_000_000)
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), userFunds, chainA, chainB)
	user1 := users[0]
	user2 := users[1]

	// Wait a few blocks for relayer to start and for user accounts to be created
	err = testutil.WaitForBlocks(ctx, 5, chainA, chainB)
	require.NoError(t, err)

	// Store and init the test contract on Chain B
	contractAddress := setupCounterContract(ctx, chainB, user2, t)

	channel, err := ibc.GetTransferChannel(ctx, r, eRep, chainA.Config().ChainID, chainB.Config().ChainID)
	require.NoError(t, err)

	transfer := ibc.WalletAmount{
		Address: contractAddress,
		Denom:   chainA.Config().Denom,
		Amount:  int64(1),
	}

	memo := ibc.TransferOptions{
		Memo: fmt.Sprintf(`{"wasm":{"contract":"%s","msg":%s}}`, contractAddress, `{"increment":{}}`),
	}

	// Initial transfer. Account is created by the wasm execute is not so we must do this twice to properly set up
	transferTx, err := chainA.SendIBCTransfer(ctx, channel.ChannelID, user1.KeyName(), transfer, memo)
	require.NoError(t, err)
	chainAHeight, err := chainA.Height(ctx)
	require.NoError(t, err)

	_, err = testutil.PollForAck(ctx, chainA, chainAHeight-5, chainAHeight+55, transferTx.Packet)
	require.NoError(t, err)

}

func setupCounterContract(ctx context.Context, chainB *cosmos.CosmosChain, user2 ibc.Wallet, t *testing.T) string {
	codeId, err := chainB.StoreContract(ctx, user2.KeyName(), "artifacts/contracts/ibchooks_counter.wasm")
	if err != nil {
		require.NoError(t, err)
	}

	initMsg := `{"count":0}`
	cmd := []string{
		chainB.Config().Bin, "tx", "wasm", "instantiate", codeId, initMsg,
		"--label", "test counter contract",
		"--no-admin",
		"--node", chainB.GetRPCAddress(),
		"--home", chainB.HomeDir(),
		"--chain-id", chainB.Config().ChainID,
		"--from", user2.KeyName(),
		"--gas", "500000",
		"--keyring-dir", chainB.HomeDir(),
		"--keyring-backend", keyring.BackendTest,
		"-b", "block",
		"-y",
	}
	_, _, err = chainB.Exec(ctx, cmd, nil)
	require.NoError(t, err, "could not instantiate the contract")

	cmd = []string{
		chainB.Config().Bin, "q", "wasm", "list-contract-by-code", codeId,
		"--node", chainB.GetRPCAddress(),
		"--home", chainB.HomeDir(),
		"--chain-id", chainB.Config().ChainID,
	}
	stdout, _, err := chainB.Exec(ctx, cmd, nil)
	require.NoError(t, err, "could not list the contracts")
	contactsRes := cosmos.QueryContractResponse{}
	err = yaml.Unmarshal(stdout, &contactsRes)
	require.NoError(t, err, "could not unmarshal query contract response")
	contractAddress := contactsRes.Contracts[0]
	return contractAddress
}
