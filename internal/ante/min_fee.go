package ante

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

type MinFeeDecorator struct {
	keeper wasmtypes.ViewKeeper
	codec  codec.BinaryCodec
}

func NewMinFeeDecorador(codec codec.BinaryCodec, keeper wasmtypes.ViewKeeper) MinFeeDecorator {
	return MinFeeDecorator{
		codec:  codec,
		keeper: keeper,
	}
}

func (mfd MinFeeDecorator) checkMinFee(ctx sdk.Context, m sdk.Msg, fee sdk.Coins) error {
	switch msg := m.(type) {
	case *wasmtypes.MsgExecuteContract:
		info := mfd.keeper.GetContractInfo(ctx, sdk.MustAccAddressFromBech32(msg.Contract))
		if mfd.keeper.IsPinnedCode(ctx, info.CodeID) {
			return nil
		}
		if fee.IsZero() {
			return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "must provide a fee to execute the contract")
		}
		return nil
	case *wasmtypes.MsgInstantiateContract:
		if mfd.keeper.IsPinnedCode(ctx, msg.CodeID) {
			return nil
		}
		if fee.IsZero() {
			return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "must provide a fee to instantiate the contract")
		}
		return nil
	default:
		return nil
	}
}

func (mfd MinFeeDecorator) Validate(ctx sdk.Context, m sdk.Msg, fee sdk.Coins) error {
	err := mfd.checkMinFee(ctx, m, fee)
	if err != nil {
		return err
	}
	if msg, ok := m.(*authz.MsgExec); ok {
		for _, v := range msg.Msgs {
			var wrappedMsg sdk.Msg
			err := mfd.codec.UnpackAny(v, &wrappedMsg)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "error decoding authz messages")
			}
			err = mfd.checkMinFee(ctx, wrappedMsg, fee)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (mfd MinFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "tx must be a FeeTx")
	}

	msgs := tx.GetMsgs()
	for _, m := range msgs {
		err := mfd.Validate(ctx, m, feeTx.GetFee())
		if err != nil {
			return ctx, err
		}
	}
	return next(ctx, tx, simulate)
}
