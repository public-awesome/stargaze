package ante

import (
	"encoding/json"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
)

var _ sdk.AnteDecorator = FeeDecorator{}

type GlobalFeeReaderExpected interface {
	GetContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress) (types.ContractAuthorization, bool)
	GetCodeAuthorization(ctx sdk.Context, codeId uint64) (types.CodeAuthorization, bool)
	GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
}

type FeeDecorator struct {
	codec     codec.BinaryCodec
	feeKeeper GlobalFeeReaderExpected
}

func NewFeeDecorator(codec codec.BinaryCodec, fk GlobalFeeReaderExpected) FeeDecorator {
	return FeeDecorator{
		codec:     codec,
		feeKeeper: fk,
	}
}

// AnteHandle implements the AnteDecorator interface
func (mfd FeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	// Only check for minimum fees and global fee if the execution mode is CheckTx
	if !ctx.IsCheckTx() || simulate {
		return next(ctx, tx, simulate)
	}

	feeCoins := feeTx.GetFee().Sort()
	gas := feeTx.GetGas()
	msgs := feeTx.GetMsgs()

	// currently zero fees allowed only when all msgs are authorized to be zero fees
	// todo how to handle mixed msgs
	zeroFeeTx, err := mfd.containsOnlyZeroFeeMsgs(ctx, msgs)
	if err != nil {
		return ctx, err
	}

	if !zeroFeeTx {
		requiredFees := getMinGasPrice(ctx, int64(gas))
		if requiredFees.IsAllGTE(feeCoins) { // required fees > tx fees
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInsufficientFee, "Required fees "+requiredFees.String())
		}
	}

	return next(ctx, tx, simulate)
}

func (mfd FeeDecorator) containsOnlyZeroFeeMsgs(ctx sdk.Context, msgs []sdk.Msg) (bool, error) {
	for _, m := range msgs {
		switch msg := m.(type) {
		case *wasmtypes.MsgExecuteContract:
			{
				if !mfd.isZeroFeeMsg(ctx, msg) {
					return false, nil
				}
			}
		case *authz.MsgExec:
			{
				var authzMsgs []sdk.Msg
				for _, v := range msg.Msgs {
					var wrappedMsg sdk.Msg
					err := mfd.codec.UnpackAny(v, &wrappedMsg)
					if err != nil {
						return false, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "error decoding authz messages")
					}
					authzMsgs = append(authzMsgs, wrappedMsg)
				}
				return mfd.containsOnlyZeroFeeMsgs(ctx, authzMsgs)
			}
		}
	}

	return true, nil
}

func (mfd FeeDecorator) isZeroFeeMsg(ctx sdk.Context, msg *wasmtypes.MsgExecuteContract) bool {
	contactAddr := sdk.MustAccAddressFromBech32(msg.Contract)
	contractAuth, found := mfd.feeKeeper.GetContractAuthorization(ctx, contactAddr)
	if found {
		return mfd.isAuthorizedMethod(msg.GetMsg(), contractAuth.GetMethods())
	}
	codeId := mfd.feeKeeper.GetContractInfo(ctx, contactAddr).CodeID
	codeAuth, found := mfd.feeKeeper.GetCodeAuthorization(ctx, codeId)
	if found {
		return mfd.isAuthorizedMethod(msg.GetMsg(), codeAuth.GetMethods())
	}

	return false
}

func (mfd FeeDecorator) isAuthorizedMethod(jsonBytes wasmtypes.RawContractMessage, methods []string) bool {
	document := map[string]interface{}{}

	if len(methods) == 1 && methods[0] == "*" {
		return true
	}

	// contract method fetching taken from https://github.com/CosmWasm/wasmd/blob/4c906d5a53a255c551d6ed981a548cffe47ae9f0/x/wasm/types/json_matching.go
	if err := json.Unmarshal(jsonBytes, &document); err != nil {
		return false
	}
	if len(document) != 1 {
		return false
	}

	for topLevelKey := range document {
		for _, allowedKey := range methods {
			if allowedKey == topLevelKey {
				return true
			}
		}
		return false
	}
	return false
}

// https://github.com/cosmos/gaia/blob/79626dfe1d99c6c87850ffd83f5c54666c981f87/x/globalfee/ante/fee.go#L190
func getMinGasPrice(ctx sdk.Context, gasLimit int64) sdk.Coins {
	minGasPrices := ctx.MinGasPrices()
	// special case: if minGasPrices=[], requiredFees=[]
	if minGasPrices.IsZero() {
		return sdk.Coins{}
	}

	requiredFees := make(sdk.Coins, len(minGasPrices))
	// Determine the required fees by multiplying each required minimum gas
	// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
	glDec := sdk.NewDec(gasLimit)
	for i, gp := range minGasPrices {
		fee := gp.Amount.Mul(glDec)
		requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
	}

	return requiredFees.Sort()
}
