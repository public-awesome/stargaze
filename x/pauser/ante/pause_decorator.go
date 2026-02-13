package ante

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/public-awesome/stargaze/v17/x/pauser/types"
)

var _ sdk.AnteDecorator = PauseDecorator{}

// maxNestedMsgDepth is the maximum allowed nesting depth for authz.MsgExec messages.
const maxNestedMsgDepth = 2

// ContractPauseKeeper defines the interface needed by the ante handler.
type ContractPauseKeeper interface {
	IsExecutionPaused(ctx sdk.Context, contractAddr sdk.AccAddress) bool
}

// PauseDecorator rejects MsgExecuteContract transactions targeting paused contracts.
type PauseDecorator struct {
	pauseKeeper ContractPauseKeeper
}

// NewPauseDecorator creates a new PauseDecorator.
func NewPauseDecorator(pauseKeeper ContractPauseKeeper) PauseDecorator {
	return PauseDecorator{pauseKeeper: pauseKeeper}
}

// AnteHandle checks if any MsgExecuteContract in the tx targets a paused contract.
// It also recursively inspects messages nested inside authz.MsgExec.
func (pd PauseDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if err := pd.checkMsg(ctx, msg, 0); err != nil {
			return ctx, err
		}
	}
	return next(ctx, tx, simulate)
}

// checkMsg inspects a single message for paused contract execution,
// recursing into authz.MsgExec to prevent bypass. Returns an error if
// nesting exceeds maxNestedMsgDepth.
func (pd PauseDecorator) checkMsg(ctx sdk.Context, msg sdk.Msg, depth int) error {
	if depth > maxNestedMsgDepth {
		return types.ErrNestedMsgTooDeep
	}

	switch msg := msg.(type) {
	case *wasmtypes.MsgExecuteContract:
		contractAddr, err := sdk.AccAddressFromBech32(msg.Contract)
		if err != nil {
			return nil
		}
		if pd.pauseKeeper.IsExecutionPaused(ctx, contractAddr) {
			return types.ErrContractPaused
		}
	case *authz.MsgExec:
		innerMsgs, err := msg.GetMessages()
		if err != nil {
			return nil
		}
		for _, innerMsg := range innerMsgs {
			if err := pd.checkMsg(ctx, innerMsg, depth+1); err != nil {
				return err
			}
		}
	}
	return nil
}
