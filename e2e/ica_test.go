package e2e

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	chantypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/relayer"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestInterchainAccounts is a test case that performs simulations and assertions around some basic
// features and packet flows surrounding interchain accounts. See: https://github.com/cosmos/interchain-accounts-demo
func TestInterchainAccounts(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	t.Parallel()

	client, network := interchaintest.DockerSetup(t)

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)
	ctx := context.Background()

	// Get both chains
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name: "icad",
			ChainConfig: ibc.ChainConfig{
				Images: []ibc.DockerImage{{Repository: "ghcr.io/cosmos/ibc-go-icad", Version: "v0.1.7", UidGid: "1025:1025"}},
			},
		},
		{
			Name:        "stargaze",
			ChainConfig: stargazeCfg,
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	icadChain, starsdChain := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	// Get a relayer instance
	r := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		relayer.StartupFlags("-p", "events", "-b", "100"),
	).Build(t, client, network)

	// Build the network; spin up the chains and configure the relayer
	const pathName = "test-path"
	const relayerName = "relayer"

	ic := interchaintest.NewInterchain().
		AddChain(icadChain).
		AddChain(starsdChain).
		AddRelayer(r, relayerName).
		AddLink(interchaintest.InterchainLink{
			Chain1:  icadChain,
			Chain2:  starsdChain,
			Relayer: r,
			Path:    pathName,
		})

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: true,
	}))

	// Fund a user account on chain1 and chain2
	const userFunds = int64(10_000_000_000)
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), userFunds, icadChain, starsdChain)
	icadUser := users[0]
	starsdUser := users[1]

	// Generate a new IBC path
	err = r.GeneratePath(ctx, eRep, icadChain.Config().ChainID, starsdChain.Config().ChainID, pathName)
	require.NoError(t, err)

	// Create new clients
	err = r.CreateClients(ctx, eRep, pathName, ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	err = testutil.WaitForBlocks(ctx, 2, icadChain, starsdChain)
	require.NoError(t, err)

	// Create a new connection
	err = r.CreateConnections(ctx, eRep, pathName)
	require.NoError(t, err)

	err = testutil.WaitForBlocks(ctx, 2, icadChain, starsdChain)
	require.NoError(t, err)

	// Query for the newly created connection
	connections, err := r.GetConnections(ctx, eRep, icadChain.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 1, len(connections))

	// Start the relayer and set the cleanup function.
	err = r.StartRelayer(ctx, eRep, pathName)
	require.NoError(t, err)

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				t.Logf("an error occurred while stopping the relayer: %s", err)
			}
		},
	)

	// Register a new interchain account on starsdChain, on behalf of the user acc on icadChain
	icadAddr := icadUser.(*cosmos.CosmosWallet).FormattedAddressWithPrefix(icadChain.Config().Bech32Prefix)
	registerICA := []string{
		icadChain.Config().Bin, "tx", "intertx", "register",
		"--from", icadAddr,
		"--connection-id", connections[0].ID,
		"--chain-id", icadChain.Config().ChainID,
		"--home", icadChain.HomeDir(),
		"--node", icadChain.GetRPCAddress(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	_, _, err = icadChain.Exec(ctx, registerICA, nil)
	require.NoError(t, err)

	ir := cosmos.DefaultEncoding().InterfaceRegistry

	c2h, err := starsdChain.Height(ctx)
	require.NoError(t, err)

	channelFound := func(found *chantypes.MsgChannelOpenConfirm) bool {
		return found.PortId == "icahost"
	}

	// Wait for channel open confirm
	_, err = cosmos.PollForMessage(ctx, starsdChain, ir,
		c2h, c2h+30, channelFound)
	require.NoError(t, err)

	// Query for the newly registered interchain account
	queryICA := []string{
		icadChain.Config().Bin, "query", "intertx", "interchainaccounts", connections[0].ID, icadAddr,
		"--chain-id", icadChain.Config().ChainID,
		"--home", icadChain.HomeDir(),
		"--node", icadChain.GetRPCAddress(),
	}
	stdout, _, err := icadChain.Exec(ctx, queryICA, nil)
	require.NoError(t, err)

	icaAddr := parseInterchainAccountField(stdout)
	require.NotEmpty(t, icaAddr)

	// Get initial account balances
	starsdAddr := starsdUser.(*cosmos.CosmosWallet).FormattedAddressWithPrefix(starsdChain.Config().Bech32Prefix)

	starsdOrigBal, err := starsdChain.GetBalance(ctx, starsdAddr, starsdChain.Config().Denom)
	require.NoError(t, err)

	icaOrigBal, err := starsdChain.GetBalance(ctx, icaAddr, starsdChain.Config().Denom)
	require.NoError(t, err)

	// Send funds to ICA from user account on starsd
	transferAmount := math.NewInt(1000)
	transfer := ibc.WalletAmount{
		Address: icaAddr,
		Denom:   starsdChain.Config().Denom,
		Amount:  transferAmount,
	}
	err = starsdChain.SendFunds(ctx, starsdUser.KeyName(), transfer)
	require.NoError(t, err)

	starsdBal, err := starsdChain.GetBalance(ctx, starsdAddr, starsdChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, starsdBal, starsdOrigBal.Sub(transferAmount))

	icaBal, err := starsdChain.GetBalance(ctx, icaAddr, starsdChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, icaBal, icaOrigBal.Add(transferAmount))

	// Build bank transfer msg
	rawMsg, err := json.Marshal(map[string]any{
		"@type":        "/cosmos.bank.v1beta1.MsgSend",
		"from_address": icaAddr,
		"to_address":   starsdAddr,
		"amount": []map[string]any{
			{
				"denom":  starsdChain.Config().Denom,
				"amount": transferAmount.String(),
			},
		},
	})
	require.NoError(t, err)

	// Send bank transfer msg to ICA on starsd from the user account on icad
	sendICATransfer := []string{
		icadChain.Config().Bin, "tx", "intertx", "submit", string(rawMsg),
		"--connection-id", connections[0].ID,
		"--from", icadAddr,
		"--chain-id", icadChain.Config().ChainID,
		"--home", icadChain.HomeDir(),
		"--node", icadChain.GetRPCAddress(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	_, _, err = icadChain.Exec(ctx, sendICATransfer, nil)
	require.NoError(t, err)

	// Wait for tx to be relayed
	c1h, err := icadChain.Height(ctx)
	require.NoError(t, err)

	ackFound := func(found *chantypes.MsgAcknowledgement) bool {
		return found.Packet.Sequence == 1 &&
			found.Packet.SourcePort == "icacontroller-"+icadAddr &&
			found.Packet.DestinationPort == "icahost"
	}

	// Wait for ack
	_, err = cosmos.PollForMessage(ctx, icadChain, ir, c1h, c1h+10, ackFound)
	require.NoError(t, err)

	// Assert that the funds have been received by the user account on starsd
	starsdBal, err = starsdChain.GetBalance(ctx, starsdAddr, starsdChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, starsdBal, starsdOrigBal)

	// Assert that the funds have been removed from the ICA on chain2
	icaBal, err = starsdChain.GetBalance(ctx, icaAddr, starsdChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, icaBal, icaOrigBal)
}

// parseInterchainAccountField takes a slice of bytes which should be returned when querying for an ICA via
// the 'intertx interchainaccounts' cmd and splices out the actual address portion.
func parseInterchainAccountField(stdout []byte) string {
	// After querying an ICA the stdout should look like the following,
	// interchain_account_address: cosmos1p76n3mnanllea4d3av0v0e42tjj03cae06xq8fwn9at587rqp23qvxsv0j
	// So we split the string at the : and then grab the address and return.
	parts := strings.SplitN(string(stdout), ":", 2)
	icaAddr := strings.TrimSpace(parts[1])
	return icaAddr
}
