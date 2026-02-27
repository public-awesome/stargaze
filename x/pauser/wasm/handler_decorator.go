package wasm

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/v2/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/pauser/keeper"
	"github.com/public-awesome/stargaze/v17/x/pauser/types"
)

// NewPauseMessageHandlerDecorator returns a decorator that intercepts wasm Execute/Migrate
// messages to paused contracts. It accepts a pointer to keeper to handle late initialization.
func NewPauseMessageHandlerDecorator(pauseKeeper *keeper.Keeper) func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &pauseMessenger{
			wrapped:     old,
			pauseKeeper: pauseKeeper,
		}
	}
}

type pauseMessenger struct {
	wrapped     wasmkeeper.Messenger
	pauseKeeper *keeper.Keeper
}

func (pm *pauseMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, [][]*codectypes.Any, error) {
	if msg.Wasm != nil {
		var targetContract string
		switch {
		case msg.Wasm.Execute != nil:
			targetContract = msg.Wasm.Execute.ContractAddr
		case msg.Wasm.Migrate != nil:
			targetContract = msg.Wasm.Migrate.ContractAddr
		}

		if len(targetContract) > 0 {
			targetAddr, err := sdk.AccAddressFromBech32(targetContract)
			if err == nil && pm.pauseKeeper.IsExecutionPaused(ctx, targetAddr) {
				return nil, nil, nil, types.ErrContractPaused
			}
		}
	}
	return pm.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}
