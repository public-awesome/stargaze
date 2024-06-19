package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
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
	cdc codec.BinaryCodec
}

func NewCheckDecorator(cdc codec.BinaryCodec) CheckDecorator {
	return CheckDecorator{
		cdc: cdc,
	}
}

func (cdc CheckDecorator) CheckMessage(m sdk.Msg) error {
	switch msg := m.(type) {
	case *ibctransfertypes.MsgTransfer:
		if len(msg.Receiver) > MaximumReceiverLength {
			return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver")
		}
		if len(msg.Memo) > MaximumMemoLength {
			return errorsmod.Wrap(sdkerrors.ErrInvalidType, "invalid memo")
		}
	case *icacontrollertypes.MsgSendTx:
		if len(msg.Owner) > MaximumOwnerLength {
			return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner")
		}
	case *icacontrollertypes.MsgRegisterInterchainAccount:
		if len(msg.Owner) > MaximumOwnerLength {
			return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner")
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
			if msg, ok := m.(*authz.MsgExec); ok {
				for _, v := range msg.Msgs {
					var wrappedMsg sdk.Msg
					err := cd.cdc.UnpackAny(v, &wrappedMsg)
					if err != nil {
						return ctx, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "error decoding authz messages")
					}
					err = cd.CheckMessage(wrappedMsg)
					if err != nil {
						return ctx, err
					}
				}
			}
		}
	}
	return next(ctx, tx, simulate)
}
