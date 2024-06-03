package e2e

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	chantypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/public-awesome/stargaze/v14/app"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/relayer"
	"github.com/strangelove-ventures/interchaintest/v8/testreporter"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestInterchainAccounts is a test case that performs simulations and assertions around some basic
// features and packet flows surrounding interchain accounts.
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
	stargazeCfg1.ChainID = "stargaze-c" // stargaze controller
	stargazeCfg2 := stargazeCfg
	stargazeCfg2.ChainID = "stargaze-h" // stargaze host

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
	userFunds := math.NewInt(10_000_000_000)
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
	ir := app.MakeEncodingConfig().InterfaceRegistry

	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				t.Logf("an error occurred while stopping the relayer: %s", err)
			}
		},
	)

	// Register a new interchain account on hostChain, on behalf of the user acc on controllerChain
	icaMetadata := icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: connection.ID,
		HostConnectionId:       connection.Counterparty.ConnectionId,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}
	icaMetadataBytes, err := icatypes.ModuleCdc.MarshalJSON(&icaMetadata)
	require.NoError(t, err)
	version := string(icaMetadataBytes)
	registerICA := []string{
		controllerChain.Config().Bin, "tx", "interchain-accounts", "controller", "register", connection.ID,
		"--version", version,
		"--from", controllerUser.FormattedAddress(),
		"--chain-id", controllerChain.Config().ChainID,
		"--home", controllerChain.HomeDir(),
		"--node", controllerChain.GetRPCAddress(),
		"--keyring-backend", keyring.BackendTest,
		"--output", "json",
		"-y",
	}
	_, _, err = controllerChain.Exec(ctx, registerICA, nil)
	require.NoError(t, err)

	controllerHeight, err := controllerChain.Height(ctx)
	require.NoError(t, err)

	// Wait for channel open confirm
	_, err = cosmos.PollForMessage(ctx, controllerChain, ir,
		controllerHeight, controllerHeight+10, func(found *chantypes.MsgChannelOpenAck) bool {
			return found.PortId == "icacontroller-"+controllerUser.FormattedAddress()
		})
	require.NoError(t, err)

	// Query for the newly registered interchain account address
	queryICA := []string{
		controllerChain.Config().Bin, "query", "interchain-accounts", "controller", "interchain-account", controllerUser.FormattedAddress(), connection.ID,
		"--home", controllerChain.HomeDir(),
		"--node", controllerChain.GetRPCAddress(),
	}
	stdout, _, err := controllerChain.Exec(ctx, queryICA, nil)
	require.NoError(t, err)

	icaAddr := parseInterchainAccountField(stdout)
	require.NotEmpty(t, icaAddr)

	// Get initial account balances
	hostOrigBal, err := hostChain.GetBalance(ctx, hostUser.FormattedAddress(), hostChain.Config().Denom)
	require.NoError(t, err)

	controllerOrigBal, err := hostChain.GetBalance(ctx, icaAddr, hostChain.Config().Denom)
	require.NoError(t, err)
	require.Zero(t, controllerOrigBal.Int64()) // ensuring ica address is zero balance

	// Send funds to ICA from faucet account on host chain
	transferAmount := math.NewInt(1000)
	transfer := ibc.WalletAmount{
		Address: icaAddr,
		Denom:   hostChain.Config().Denom,
		Amount:  transferAmount,
	}
	err = hostChain.SendFunds(ctx, hostUser.KeyName(), transfer)
	require.NoError(t, err)

	err = testutil.WaitForBlocks(ctx, 1, hostChain)
	require.NoError(t, err)

	// Ensuring the balances are correct
	hostBal, err := hostChain.GetBalance(ctx, hostUser.FormattedAddress(), hostChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, hostBal, hostOrigBal.Sub(transferAmount)) // after sending funds to ICA
	icaBal, err := hostChain.GetBalance(ctx, icaAddr, hostChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, icaBal, controllerOrigBal.Add(transferAmount)) // after receiving funds from faucet

	// Build bank transfer msg
	msg := SendMsg{
		Type:        "/cosmos.bank.v1beta1.MsgSend",
		FromAddress: icaAddr,
		ToAddress:   hostUser.FormattedAddress(),
		Amount: []Amount{
			{
				Denom:  hostChain.Config().Denom,
				Amount: transferAmount.String(),
			},
		},
	}
	msgBytes, err := json.Marshal(msg)
	require.NoError(t, err)
	var sdkmsg sdk.Msg
	cdc := codec.NewProtoCodec(ir)
	err = cdc.UnmarshalInterfaceJSON(msgBytes, &sdkmsg)
	require.NoError(t, err)
	icaPacketDataBytes, err := icatypes.SerializeCosmosTx(cdc, []proto.Message{sdkmsg}, "proto3")
	require.NoError(t, err)
	icaPacketBytes, err := cdc.MarshalJSON(&icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: icaPacketDataBytes,
		Memo: "icatest",
	})
	require.NoError(t, err)

	// Send bank transfer msg to ICA on host chain from the user account on icad
	sendICATransfer := []string{
		controllerChain.Config().Bin, "tx", "interchain-accounts", "controller", "send-tx", connection.ID, string(icaPacketBytes),
		"--from", controllerUser.FormattedAddress(),
		"--chain-id", controllerChain.Config().ChainID,
		"--home", controllerChain.HomeDir(),
		"--node", controllerChain.GetRPCAddress(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	_, _, err = controllerChain.Exec(ctx, sendICATransfer, nil)
	require.NoError(t, err)

	// Wait for tx to be relayed and acknowledged
	controllerHeight, err = controllerChain.Height(ctx)
	require.NoError(t, err)
	_, err = cosmos.PollForMessage(ctx, controllerChain, ir, controllerHeight, controllerHeight+10, func(found *chantypes.MsgAcknowledgement) bool {
		return found.Packet.Sequence == 1 &&
			found.Packet.SourcePort == "icacontroller-"+controllerUser.FormattedAddress() &&
			found.Packet.DestinationPort == "icahost"
	})
	require.NoError(t, err)

	// Assert that the funds have been received by the user account on host
	hostBal, err = hostChain.GetBalance(ctx, hostUser.FormattedAddress(), hostChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, hostBal.Int64(), hostOrigBal.Int64())

	// Assert that the funds have been removed from the ICA on chain2
	icaBal, err = hostChain.GetBalance(ctx, icaAddr, hostChain.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, icaBal, controllerOrigBal)
}

// parseInterchainAccountField takes a slice of bytes which should be returned when querying for an ICA via
// the 'interchain-accounts controller interchain-account' cmd and splices out the actual address portion.
func parseInterchainAccountField(stdout []byte) string {
	// After querying an ICA the stdout should look like the following,
	// address: cosmos1p76n3mnanllea4d3av0v0e42tjj03cae06xq8fwn9at587rqp23qvxsv0j
	// So we split the string at the : and then grab the address and return.
	parts := strings.SplitN(string(stdout), ":", 2)
	icaAddr := strings.TrimSpace(parts[1])
	return icaAddr
}

type SendMsg struct {
	Type        string   `json:"@type"`
	FromAddress string   `json:"from_address"`
	ToAddress   string   `json:"to_address"`
	Amount      []Amount `json:"amount"`
}
type Amount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
