package ante

import (
	"context"
	"encoding/json"
	"errors"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stargaze/v13/x/globalfee/types"
)

var _ sdk.AnteDecorator = FeeDecorator{}

type GlobalFeeReaderExpected interface {
	GetContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress) (types.ContractAuthorization, bool)
	GetCodeAuthorization(ctx sdk.Context, codeID uint64) (types.CodeAuthorization, bool)
	GetContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
	GetParams(ctx sdk.Context) types.Params
}

type StakingReaderExpected interface {
	BondDenom(ctx context.Context) (string, error)
}

type FeeDecorator struct {
	codec         codec.BinaryCodec
	feeKeeper     GlobalFeeReaderExpected
	stakingKeeper StakingReaderExpected
}

func NewFeeDecorator(codec codec.BinaryCodec, fk GlobalFeeReaderExpected, sk StakingReaderExpected) FeeDecorator {
	return FeeDecorator{
		codec:         codec,
		feeKeeper:     fk,
		stakingKeeper: sk,
	}
}

// AnteHandle implements the AnteDecorator interface
func (mfd FeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must implement the sdk.FeeTx interface")
	}

	// Only check for minimum fees and global fee if the execution mode is CheckTx
	if !ctx.IsCheckTx() || simulate {
		return next(ctx, tx, simulate)
	}

	msgs := feeTx.GetMsgs()

	// currently accepting zero fee transactions only when the tx contains only the authorized operations that can bypass the minimum fee
	onlyZeroFeeMsgs := mfd.containsOnlyZeroFeeMsgs(ctx, msgs)

	return mfd.checkFees(ctx, feeTx, tx, onlyZeroFeeMsgs, simulate, next) // https://github.com/cosmos/gaia/blob/6fe097e3280baa360a28b59a29b8cca964a5ae97/x/globalfee/ante/fee.go
}

func (mfd FeeDecorator) containsOnlyZeroFeeMsgs(ctx sdk.Context, msgs []sdk.Msg) bool {
	for _, m := range msgs {
		switch msg := m.(type) {
		case *wasmtypes.MsgExecuteContract:
			{
				if !mfd.isZeroFeeMsg(ctx, msg) {
					return false
				}
			}
		default:
			return false
		}
	}

	return true
}

func (mfd FeeDecorator) isZeroFeeMsg(ctx sdk.Context, msg *wasmtypes.MsgExecuteContract) bool {
	contactAddr := sdk.MustAccAddressFromBech32(msg.Contract)
	contractAuth, found := mfd.feeKeeper.GetContractAuthorization(ctx, contactAddr)
	if found {
		return isAuthorizedMethod(msg.GetMsg(), contractAuth.GetMethods())
	}
	codeID := mfd.feeKeeper.GetContractInfo(ctx, contactAddr).CodeID
	codeAuth, found := mfd.feeKeeper.GetCodeAuthorization(ctx, codeID)
	if found {
		return isAuthorizedMethod(msg.GetMsg(), codeAuth.GetMethods())
	}

	return false
}

func isAuthorizedMethod(jsonBytes wasmtypes.RawContractMessage, methods []string) bool {
	document := map[string]interface{}{}

	if len(methods) == 1 && methods[0] == "*" {
		return true
	}

	if err := jsonBytes.ValidateBasic(); err != nil {
		return false
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

// The fee checking ante mechanism below is based on the x/GlobalFee/ante from cosmos/gaia
// https://github.com/cosmos/gaia/blob/6fe097e3280baa360a28b59a29b8cca964a5ae97/x/globalfee/ante/fee.go
func (mfd FeeDecorator) checkFees(ctx sdk.Context, feeTx sdk.FeeTx, tx sdk.Tx, onlyZeroFeeMsgs bool, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	feeCoins := feeTx.GetFee().Sort()
	gas := feeTx.GetGas()

	// Get required Global Fee set by module
	requiredGlobalFees, err := mfd.getGlobalFee(ctx, feeTx)
	if err != nil {
		return ctx, err
	}

	// Get local minimum-gas-prices set by the validator node
	localFees := getMinGasPrice(ctx, int64(gas))

	// CombinedFeeRequirement should never be empty since
	// global fee is set to its default value, i.e. 0ustars, if empty
	combinedFeeRequired := combinedFeeRequirement(requiredGlobalFees, localFees)
	if len(combinedFeeRequired) == 0 {
		return ctx, errorsmod.Wrapf(sdkerrors.ErrNotFound, "required fees are not setup.")
	}

	nonZeroCoinFeesReq, zeroCoinFeesDenomReq := getNonZeroFees(combinedFeeRequired)

	// feeCoinsNonZeroDenom contains non-zero denominations from the combinedFeeRequirement
	//
	// feeCoinsNoZeroDenom is used to check if the fees meets the requirement imposed by nonZeroCoinFeesReq
	// when feeCoins does not contain zero coins' denoms in combinedFeeRequirement
	// feeCoinsNonZeroDenom, feeCoinsZeroDenom := splitCoinsByDenoms(feeCoins, zeroCoinFeesDenomReq)
	feeCoinsNonZeroDenom, feeCoinsZeroDenom := splitCoinsByDenoms(feeCoins, zeroCoinFeesDenomReq)

	// Check that the fees are in expected denominations.
	// if feeCoinsNoZeroDenom=[], DenomsSubsetOf returns true
	// if feeCoinsNoZeroDenom is not empty, but nonZeroCoinFeesReq empty, return false
	if !feeCoinsNonZeroDenom.DenomsSubsetOf(nonZeroCoinFeesReq) {
		return ctx, errorsmod.Wrapf(sdkerrors.ErrInsufficientFee, "fee is not a subset of required fees; got %s, required: %s", feeCoins, combinedFeeRequired)
	}

	// only check feeCoinsNoZeroDenom has coins IsAnyGTE than nonZeroCoinFeesReq
	// when feeCoins does not contain denoms of zero denoms in combinedFeeRequirement
	if !onlyZeroFeeMsgs && len(feeCoinsZeroDenom) == 0 {
		// special case: when feeCoins=[] and there is zero coin in fee requirement
		if len(feeCoins) == 0 && len(zeroCoinFeesDenomReq) != 0 {
			return next(ctx, tx, simulate)
		}

		// Check that the amounts of the fees are greater or equal than
		// the expected amounts, i.e., at least one feeCoin amount must
		// be greater or equal to one of the combined required fees.

		// if feeCoinsNoZeroDenom=[], return false
		// if nonZeroCoinFeesReq=[], return false (this situation should not happen
		// because when nonZeroCoinFeesReq empty, and DenomsSubsetOf check passed,
		// the tx should already passed before)
		if !feeCoinsNonZeroDenom.IsAnyGTE(nonZeroCoinFeesReq) {
			return ctx, errorsmod.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeCoins, combinedFeeRequired)
		}
	}

	return next(ctx, tx, simulate)
}

// getGlobalFee returns the global fees for a given fee tx's gas
// (might also return 0denom if globalMinGasPrice is 0)
// sorted in ascending order.
// Note that ParamStoreKeyMinGasPrices type requires coins sorted.
func (mfd FeeDecorator) getGlobalFee(ctx sdk.Context, feeTx sdk.FeeTx) (sdk.Coins, error) {
	var (
		globalMinGasPrices sdk.DecCoins
		err                error
	)

	globalMinGasPrices = mfd.feeKeeper.GetParams(ctx).MinimumGasPrices

	// global fee is empty set, set global fee to 0uatom
	if len(globalMinGasPrices) == 0 {
		globalMinGasPrices, err = mfd.defaultZeroGlobalFee(ctx)
		if err != nil {
			return sdk.Coins{}, err
		}
	}
	requiredGlobalFees := make(sdk.Coins, len(globalMinGasPrices))
	// Determine the required fees by multiplying each required minimum gas
	// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
	glDec := sdkmath.LegacyNewDec(int64(feeTx.GetGas()))
	for i, gp := range globalMinGasPrices {
		fee := gp.Amount.Mul(glDec)
		requiredGlobalFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
	}

	return requiredGlobalFees.Sort(), nil
}

// getMinGasPrice returns the validator's minimum gas prices
// fees given a gas limit
func getMinGasPrice(ctx sdk.Context, gasLimit int64) sdk.Coins {
	minGasPrices := ctx.MinGasPrices()
	// special case: if minGasPrices=[], requiredFees=[]
	if minGasPrices.IsZero() {
		return sdk.Coins{}
	}

	requiredFees := make(sdk.Coins, len(minGasPrices))
	// Determine the required fees by multiplying each required minimum gas
	// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
	glDec := sdkmath.LegacyNewDec(gasLimit)
	for i, gp := range minGasPrices {
		fee := gp.Amount.Mul(glDec)
		requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
	}

	return requiredFees.Sort()
}

// combinedFeeRequirement returns the global fee and min_gas_price combined and sorted.
// Both globalFees and minGasPrices must be valid, but combinedFeeRequirement
// does not validate them, so it may return 0denom.
// if globalfee is empty, combinedFeeRequirement return sdk.Coins{}
func combinedFeeRequirement(globalFees, minGasPrices sdk.Coins) sdk.Coins {
	// empty min_gas_price
	if len(minGasPrices) == 0 {
		return globalFees
	}
	// empty global fee is not possible if we set default global fee
	if len(globalFees) == 0 && len(minGasPrices) != 0 {
		return sdk.Coins{}
	}

	// if min_gas_price denom is in globalfee, and the amount is higher than globalfee, add min_gas_price to allFees
	var allFees sdk.Coins
	for _, fee := range globalFees {
		// min_gas_price denom in global fee
		ok, c := find(minGasPrices, fee.Denom)
		if ok && c.Amount.GT(fee.Amount) {
			allFees = append(allFees, c)
		} else {
			allFees = append(allFees, fee)
		}
	}

	return allFees.Sort()
}

// getNonZeroFees returns the given fees nonzero coins
// and a map storing the zero coins's denoms
func getNonZeroFees(fees sdk.Coins) (sdk.Coins, map[string]bool) {
	requiredFeesNonZero := sdk.Coins{}
	requiredFeesZeroDenom := map[string]bool{}

	for _, gf := range fees {
		if gf.IsZero() {
			requiredFeesZeroDenom[gf.Denom] = true
		} else {
			requiredFeesNonZero = append(requiredFeesNonZero, gf)
		}
	}

	return requiredFeesNonZero.Sort(), requiredFeesZeroDenom
}

// splitCoinsByDenoms returns the given coins split in two whether
// their demon is or isn't found in the given denom map.
func splitCoinsByDenoms(feeCoins sdk.Coins, denomMap map[string]bool) (feeCoinsNonZeroDenom sdk.Coins, feeCoinsZeroDenom sdk.Coins) {
	for _, fc := range feeCoins {
		_, found := denomMap[fc.Denom]
		if found {
			feeCoinsZeroDenom = append(feeCoinsZeroDenom, fc)
		} else {
			feeCoinsNonZeroDenom = append(feeCoinsNonZeroDenom, fc)
		}
	}

	return feeCoinsNonZeroDenom.Sort(), feeCoinsZeroDenom.Sort()
}

func (mfd FeeDecorator) defaultZeroGlobalFee(ctx sdk.Context) ([]sdk.DecCoin, error) {
	bondDenom, err := mfd.getBondDenom(ctx)
	if err != nil {
		return nil, err
	}
	if bondDenom == "" {
		return nil, errors.New("empty staking bond denomination")
	}
	return []sdk.DecCoin{sdk.NewDecCoinFromDec(bondDenom, sdkmath.LegacyNewDec(0))}, nil
}

// find replaces the functionality of Coins.find from SDK v0.46.x
func find(coins sdk.Coins, denom string) (bool, sdk.Coin) {
	switch len(coins) {
	case 0:
		return false, sdk.Coin{}

	case 1:
		coin := coins[0]
		if coin.Denom == denom {
			return true, coin
		}
		return false, sdk.Coin{}

	default:
		midIdx := len(coins) / 2 // 2:1, 3:1, 4:2
		coin := coins[midIdx]
		switch {
		case denom < coin.Denom:
			return find(coins[:midIdx], denom)
		case denom == coin.Denom:
			return true, coin
		default:
			return find(coins[midIdx+1:], denom)
		}
	}
}

func (mfd FeeDecorator) getBondDenom(ctx sdk.Context) (string, error) {
	return mfd.stakingKeeper.BondDenom(ctx)
}
