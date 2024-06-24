package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

const (
	MaximumReceiverLength = 2048
	MaximumMemoLength     = 32_768
	MaximumOwnerLength    = 2048
	MaxSize               = 500_000
)

type CheckDecorator struct {
	cdc codec.BinaryCodec
}

func NewCheckDecorator(cdc codec.BinaryCodec) CheckDecorator {
	return CheckDecorator{
		cdc: cdc,
	}
}

func (cd CheckDecorator) CheckMessage(m sdk.Msg) error {
	switch msg := m.(type) {
	case *ibctransfertypes.MsgTransfer:
		if msg.Size() > MaxSize {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "msg size is too large")
		}
	case *icacontrollertypes.MsgSendTx:
		if msg.Size() > MaxSize {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "msg size is too large")
		}
	case *icacontrollertypes.MsgRegisterInterchainAccount:
		if msg.Size() > MaxSize {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "msg size is too large")
		}
	}
	return nil
}

func (cd CheckDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if ctx.IsCheckTx() {
		for _, m := range tx.GetMsgs() {
			err := cd.CheckMessage(m)
			if err != nil {
				return ctx, err
			}
		}
	}
	return next(ctx, tx, simulate)
}
