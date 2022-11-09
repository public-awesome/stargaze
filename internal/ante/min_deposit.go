package ante

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/x/authz"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type MinDepositDecorator struct {
	codec codec.BinaryCodec
}

func NewMinDepositDecorator(codec codec.BinaryCodec) MinDepositDecorator {
	return MinDepositDecorator{
		codec,
	}
}

func checkDeposit(m sdk.Msg) error {
	switch msg := m.(type) {
	case *govtypes.MsgSubmitProposal:
		c := msg.GetInitialDeposit()
		if c.AmountOf("ustars").LT(sdk.NewInt(1_000_000_000)) {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "min deposit cannot be lower than 1,000 STARS")
		}
	default:
		return nil
	}
	return nil
}

func (dec MinDepositDecorator) Validate(m sdk.Msg) error {
	err := checkDeposit(m)
	if err != nil {
		return err
	}
	if msg, ok := m.(*authz.MsgExec); ok {
		for _, v := range msg.Msgs {
			var wrappedMsg sdk.Msg
			err := dec.codec.UnpackAny(v, &wrappedMsg)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "error decoding authz messages")
			}
			err = checkDeposit(wrappedMsg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (dec MinDepositDecorator) AnteHandle(
	ctx sdk.Context, tx sdk.Tx,
	simulate bool, next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	for _, m := range msgs {
		err := dec.Validate(m)
		if err != nil {
			return ctx, err
		}
	}
	return next(ctx, tx, simulate)
}
