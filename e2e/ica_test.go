package e2e

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	chantypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
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

	stargazeCfg1 := stargazeCfg
	stargazeCfg1.ChainID = "stargaze-1"
	stargazeCfg2 := stargazeCfg
	stargazeCfg2.ChainID = "stargaze-2"

	// Get both chains
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:        "stargaze",
			ChainConfig: stargazeCfg1,
		},
		{
			Name:        "stargaze",
			ChainConfig: stargazeCfg2,
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	controllerChain, hostChain := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

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
		AddChain(controllerChain).
		AddChain(hostChain).
		AddRelayer(r, relayerName).
		AddLink(interchaintest.InterchainLink{
			Chain1:  controllerChain,
			Chain2:  hostChain,
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
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), userFunds, controllerChain, hostChain)
	controllerUser := users[0]
	hostUser := users[1]

	// Generate a new IBC path
	err = r.GeneratePath(ctx, eRep, controllerChain.Config().ChainID, hostChain.Config().ChainID, pathName)
	require.NoError(t, err)

	// Create new clients
	err = r.CreateClients(ctx, eRep, pathName, ibc.CreateClientOptions{TrustingPeriod: "330h"})
	require.NoError(t, err)

	err = testutil.WaitForBlocks(ctx, 2, controllerChain, hostChain)
	require.NoError(t, err)

	// Create a new connection
	err = r.CreateConnections(ctx, eRep, pathName)
	require.NoError(t, err)

	err = testutil.WaitForBlocks(ctx, 2, controllerChain, hostChain)
	require.NoError(t, err)

	// Query for the newly created connection
	connections, err := r.GetConnections(ctx, eRep, controllerChain.Config().ChainID)
	require.NoError(t, err)
	require.Equal(t, 2, len(connections))

	connection := connections[0]

	// Start the relayer and set the cleanup function.
	err = r.StartRelayer(ctx, eRep, pathName)
	require.NoError(t, err)

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				t.Logf("an error occurred while stopping the relayer: %s", err)
			}
			// err = os.Remove("msg.json")
			// if err != nil {
			// 	t.Logf("an error occurred while removing the msg.json file: %s", err)
			// }
		},
	)

	// Register a new interchain account on starsdChain, on behalf of the user acc on icadChain
	registerICA := []string{
		controllerChain.Config().Bin, "tx", "interchain-accounts", "controller", "register", connection.ID,
		"--from", controllerUser.FormattedAddress(),
		"--chain-id", controllerChain.Config().ChainID,
		"--home", controllerChain.HomeDir(),
		"--node", controllerChain.GetRPCAddress(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	_, _, err = controllerChain.Exec(ctx, registerICA, nil)
	require.NoError(t, err)

	ir := cosmos.DefaultEncoding().InterfaceRegistry

	c2h, err := hostChain.Height(ctx)
	require.NoError(t, err)

	channelFound := func(found *chantypes.MsgChannelOpenConfirm) bool {
		return found.PortId == "icahost"
	}

	// Wait for channel open confirm
	_, err = cosmos.PollForMessage(ctx, hostChain, ir,
		c2h, c2h+30, channelFound)
	require.NoError(t, err)

	// Query for the newly registered interchain account
	queryICA := []string{
		controllerChain.Config().Bin, "query", "interchain-accounts", "controller", "interchain-account", controllerUser.FormattedAddress(), connection.ID,
		"--chain-id", controllerChain.Config().ChainID,
		"--home", controllerChain.HomeDir(),
		"--node", controllerChain.GetRPCAddress(),
	}
	stdout, _, err := controllerChain.Exec(ctx, queryICA, nil)
	require.NoError(t, err)
	t.Log(string(stdout))

	icaAddr := parseInterchainAccountField(stdout)
	require.NotEmpty(t, icaAddr)

	// Get initial account balances

	starsdOrigBal, err := hostChain.GetBalance(ctx, hostUser.FormattedAddress(), hostChain.Config().Denom)
	require.NoError(t, err)

	icaOrigBal, err := hostChain.GetBalance(ctx, icaAddr, hostChain.Config().Denom)
	require.NoError(t, err)

	// Send funds to ICA from user account on starsd
	transferAmount := math.NewInt(1000)
	transfer := ibc.WalletAmount{
		Address: icaAddr,
		Denom:   hostChain.Config().Denom,
		Amount:  transferAmount,
	}
	err = hostChain.SendFunds(ctx, hostUser.KeyName(), transfer)
	require.NoError(t, err)

	starsdBal, err := hostChain.GetBalance(ctx, hostUser.FormattedAddress(), hostChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, starsdBal, starsdOrigBal.Sub(transferAmount))

	icaBal, err := hostChain.GetBalance(ctx, icaAddr, hostChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, icaBal, icaOrigBal.Add(transferAmount))

	// Build bank transfer msg
	rawMsg, err := json.Marshal(map[string]any{
		"@type":        "/cosmos.bank.v1beta1.MsgSend",
		"from_address": icaAddr,
		"to_address":   hostUser.FormattedAddress(),
		"amount": []map[string]any{
			{
				"denom":  hostChain.Config().Denom,
				"amount": transferAmount.String(),
			},
		},
	})
	require.NoError(t, err)
	err = os.WriteFile("msg.json", rawMsg, 0644)
	require.NoError(t, err)

	// Send bank transfer msg to ICA on starsd from the user account on icad
	sendICATransfer := []string{
		controllerChain.Config().Bin, "tx", "interchain-accounts", "controller", "send-tx", connection.ID, string(rawMsg),
		"--from", controllerUser.FormattedAddress(),
		"--chain-id", controllerChain.Config().ChainID,
		"--home", controllerChain.HomeDir(),
		"--node", controllerChain.GetRPCAddress(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	_, _, err = controllerChain.Exec(ctx, sendICATransfer, nil)
	require.NoError(t, err)

	// Wait for tx to be relayed
	c1h, err := controllerChain.Height(ctx)
	require.NoError(t, err)

	ackFound := func(found *chantypes.MsgAcknowledgement) bool {
		return found.Packet.Sequence == 1 &&
			found.Packet.SourcePort == "icacontroller-"+controllerUser.FormattedAddress() &&
			found.Packet.DestinationPort == "icahost"
	}

	// Wait for ack
	_, err = cosmos.PollForMessage(ctx, controllerChain, ir, c1h, c1h+10, ackFound)
	require.NoError(t, err)

	// Assert that the funds have been received by the user account on starsd
	starsdBal, err = hostChain.GetBalance(ctx, hostUser.FormattedAddress(), hostChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, starsdBal, starsdOrigBal)

	// Assert that the funds have been removed from the ICA on chain2
	icaBal, err = hostChain.GetBalance(ctx, icaAddr, hostChain.Config().Denom)
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
