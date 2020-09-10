package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMining returns a new Mining
func NewMining(minter sdk.AccAddress, coin sdk.Coin) Mining {
	return Mining{
		Minter:   minter,
		LastTime: 0,
		Total:    coin,
	}
}

// NewFaucetKey create a instance
func NewFaucetKey(armor string) FaucetKey {
	return FaucetKey{
		Armor: armor,
	}
}
