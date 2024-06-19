package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

const (
	MaximumReceiverLength = 2048
	MaximumMemoLength     = 32768
	MaximumOwnerLength    = 2048
)

type CheckDecorator struct {
}

func NewCheckDecorator() CheckDecorator {
	return CheckDecorator{}
}

func (cd CheckDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if ctx.IsCheckTx() {
		for _, m := range tx.GetMsgs() {
			switch msg := m.(type) {
			case *ibctransfertypes.MsgTransfer:
				if len(msg.Receiver) > MaximumReceiverLength {
					return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver")
				}
				if len(msg.Memo) > MaximumMemoLength {
					return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidType, "invalid memo")
				}
			case *icacontrollertypes.MsgSendTx:
				if len(msg.Owner) > MaximumOwnerLength {
					return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner")
				}
			case *icacontrollertypes.MsgRegisterInterchainAccount:
				if len(msg.Owner) > MaximumOwnerLength {
					return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner")
				}
			default:
				return next(ctx, tx, simulate)
			}
		}
	}
	return next(ctx, tx, simulate)
}
