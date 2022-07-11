package ante

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/x/authz"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type MinCommissionDecorator struct {
	codec codec.BinaryCodec
}

func NewMinCommissionDecorator(codec codec.BinaryCodec) MinCommissionDecorator {
	return MinCommissionDecorator{
		codec,
	}
}

func checkCommission(m sdk.Msg) error {
	switch msg := m.(type) {
	case *stakingtypes.MsgCreateValidator:
		c := msg.Commission
		if c.Rate.LT(sdk.NewDecWithPrec(5, 2)) {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "commission can not be lower than 5%")
		}
	case *stakingtypes.MsgEditValidator:
		// if commission rate is nil, it means only other fields are being update and we must skip this validation
		if msg.CommissionRate == nil {
			return nil
		}
		if msg.CommissionRate.LT(sdk.NewDecWithPrec(5, 2)) {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "commission can not be lower than 5%")
		}
	default:
		return nil
	}
	return nil
}

func (dec MinCommissionDecorator) Validate(m sdk.Msg) error {
	err := checkCommission(m)
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
			err = checkCommission(wrappedMsg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (dec MinCommissionDecorator) AnteHandle(
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
