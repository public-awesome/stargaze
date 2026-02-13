package wasm_test

import (
	"testing"

	wasmvmtypes "github.com/CosmWasm/wasmvm/v2/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/public-awesome/stargaze/v17/testutil/keeper"
	"github.com/public-awesome/stargaze/v17/x/pauser/types"
	pauserwasm "github.com/public-awesome/stargaze/v17/x/pauser/wasm"
	"github.com/stretchr/testify/require"
)

type mockMessenger struct {
	dispatchMsgFn func(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error)
}

func (m *mockMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
	if m.dispatchMsgFn != nil {
		return m.dispatchMsgFn(ctx, contractAddr, contractIBCPortID, msg)
	}
	return nil, nil, nil, nil
}

func TestPauseMessengerPausedContractRejected(t *testing.T) {
	k, ctx := keepertest.PauserKeeper(t)

	// Pause TestContract2
	err := k.SetPausedContract(ctx, types.PausedContract{
		ContractAddress: keepertest.TestContract2,
		PausedBy:        "someone",
	})
	require.NoError(t, err)

	wrapped := &mockMessenger{}
	decorator := pauserwasm.NewPauseMessageHandlerDecorator(&k)
	messenger := decorator(wrapped)

	callerAddr := sdk.MustAccAddressFromBech32(keepertest.TestContract1)

	// Wasm execute targeting a paused contract
	msg := wasmvmtypes.CosmosMsg{
		Wasm: &wasmvmtypes.WasmMsg{
			Execute: &wasmvmtypes.ExecuteMsg{
				ContractAddr: keepertest.TestContract2,
				Msg:          []byte(`{"increment":{}}`),
				Funds:        nil,
			},
		},
	}

	_, _, _, err = messenger.DispatchMsg(ctx, callerAddr, "", msg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrContractPaused)
}

func TestPauseMessengerUnpausedContractAllowed(t *testing.T) {
	k, ctx := keepertest.PauserKeeper(t)

	called := false
	wrapped := &mockMessenger{
		dispatchMsgFn: func(_ sdk.Context, _ sdk.AccAddress, _ string, _ wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
			called = true
			return nil, nil, nil, nil
		},
	}
	decorator := pauserwasm.NewPauseMessageHandlerDecorator(&k)
	messenger := decorator(wrapped)

	callerAddr := sdk.MustAccAddressFromBech32(keepertest.TestContract1)

	// Wasm execute targeting a non-paused contract
	msg := wasmvmtypes.CosmosMsg{
		Wasm: &wasmvmtypes.WasmMsg{
			Execute: &wasmvmtypes.ExecuteMsg{
				ContractAddr: keepertest.TestContract2,
				Msg:          []byte(`{"increment":{}}`),
				Funds:        nil,
			},
		},
	}

	_, _, _, err := messenger.DispatchMsg(ctx, callerAddr, "", msg)
	require.NoError(t, err)
	require.True(t, called, "wrapped messenger should have been called")
}

func TestPauseMessengerPausedByCodeIDRejected(t *testing.T) {
	k, ctx := keepertest.PauserKeeper(t)

	// TestContract2 has CodeID 2 per the mock
	err := k.SetPausedCodeID(ctx, types.PausedCodeID{
		CodeID:   2,
		PausedBy: "someone",
	})
	require.NoError(t, err)

	wrapped := &mockMessenger{}
	decorator := pauserwasm.NewPauseMessageHandlerDecorator(&k)
	messenger := decorator(wrapped)

	callerAddr := sdk.MustAccAddressFromBech32(keepertest.TestContract1)

	msg := wasmvmtypes.CosmosMsg{
		Wasm: &wasmvmtypes.WasmMsg{
			Execute: &wasmvmtypes.ExecuteMsg{
				ContractAddr: keepertest.TestContract2,
				Msg:          []byte(`{"increment":{}}`),
				Funds:        nil,
			},
		},
	}

	_, _, _, err = messenger.DispatchMsg(ctx, callerAddr, "", msg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrContractPaused)
}

func TestPauseMessengerNonExecuteWasmMsgAllowed(t *testing.T) {
	k, ctx := keepertest.PauserKeeper(t)

	// Pause the contract
	err := k.SetPausedContract(ctx, types.PausedContract{
		ContractAddress: keepertest.TestContract1,
		PausedBy:        "someone",
	})
	require.NoError(t, err)

	called := false
	wrapped := &mockMessenger{
		dispatchMsgFn: func(_ sdk.Context, _ sdk.AccAddress, _ string, _ wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
			called = true
			return nil, nil, nil, nil
		},
	}
	decorator := pauserwasm.NewPauseMessageHandlerDecorator(&k)
	messenger := decorator(wrapped)

	callerAddr := sdk.MustAccAddressFromBech32(keepertest.TestContract1)

	// Wasm instantiate (not execute) should pass through even when contract is paused
	msg := wasmvmtypes.CosmosMsg{
		Wasm: &wasmvmtypes.WasmMsg{
			Instantiate: &wasmvmtypes.InstantiateMsg{
				CodeID: 1,
				Msg:    []byte(`{"count":0}`),
				Funds:  nil,
				Label:  "test",
			},
		},
	}

	_, _, _, err = messenger.DispatchMsg(ctx, callerAddr, "", msg)
	require.NoError(t, err)
	require.True(t, called, "wrapped messenger should have been called")
}

func TestPauseMessengerNonWasmMsgAllowed(t *testing.T) {
	k, ctx := keepertest.PauserKeeper(t)

	called := false
	wrapped := &mockMessenger{
		dispatchMsgFn: func(_ sdk.Context, _ sdk.AccAddress, _ string, _ wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
			called = true
			return nil, nil, nil, nil
		},
	}
	decorator := pauserwasm.NewPauseMessageHandlerDecorator(&k)
	messenger := decorator(wrapped)

	callerAddr := sdk.MustAccAddressFromBech32(keepertest.TestContract1)

	// Bank send message (not wasm at all) should always pass through
	msg := wasmvmtypes.CosmosMsg{
		Bank: &wasmvmtypes.BankMsg{
			Send: &wasmvmtypes.SendMsg{
				ToAddress: keepertest.TestContract2,
				Amount:    wasmvmtypes.Array[wasmvmtypes.Coin]{{Denom: "ustars", Amount: "1000"}},
			},
		},
	}

	_, _, _, err := messenger.DispatchMsg(ctx, callerAddr, "", msg)
	require.NoError(t, err)
	require.True(t, called, "wrapped messenger should have been called")
}
