package ante

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/x/authz"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type MinDepositDecorator struct {
	codec     codec.BinaryCodec
	govKeeper govkeeper.Keeper
}

func NewMinDepositDecorator(codec codec.BinaryCodec, gk govkeeper.Keeper) MinDepositDecorator {
	return MinDepositDecorator{
		codec,
		gk,
	}
}

func (dec MinDepositDecorator) checkDeposit(ctx sdk.Context, m sdk.Msg) error {
	switch msg := m.(type) {
	case *govtypes.MsgSubmitProposal:
		params := dec.govKeeper.GetDepositParams(ctx)
		coinDenom := params.MinDeposit[0]
		minDepositAmount := sdk.NewInt(1_000_000_000)
		c := msg.GetInitialDeposit()
		if c.AmountOf(coinDenom.Denom).LT(minDepositAmount) {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("min deposit cannot be lower than %s %s", minDepositAmount.String(), coinDenom.GetDenom()))
		}
	default:
		return nil
	}
	return nil
}

func (dec MinDepositDecorator) Validate(ctx sdk.Context, m sdk.Msg) error {
	err := dec.checkDeposit(ctx, m)
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
			err = dec.checkDeposit(ctx, wrappedMsg)
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
		err := dec.Validate(ctx, m)
		if err != nil {
			return ctx, err
		}
	}
	return next(ctx, tx, simulate)
}
