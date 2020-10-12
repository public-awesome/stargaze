package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrWithdrawTooOften withdraw too often
	ErrWithdrawTooOften = sdkerrors.Register(ModuleName, 100, "Each address can withdraw only once")
	// ErrFaucetKeyEmpty empty key
	ErrFaucetKeyEmpty = sdkerrors.Register(ModuleName, 101, "Armor should not be empty")
	// ErrFaucetKeyExisted key already exists
	ErrFaucetKeyExisted = sdkerrors.Register(ModuleName, 102, "Faucet key existed")
	// ErrInvalidCoinAmount invalid coin amount
	ErrInvalidCoinAmount = sdkerrors.Register(ModuleName, 103, "Invalid coin amount")
	// ErrKeyNotFound is not on store
	ErrKeyNotFound = sdkerrors.Register(ModuleName, 104, "Key not found")
)
