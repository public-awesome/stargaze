package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// NewMining returns a new Mining
func NewMining(minter sdk.AccAddress, coin sdk.Coin) Mining {
	return Mining{
		Minter:   minter.String(),
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

// DenomConfig holds configuration for each individual denom minting
type DenomConfig struct {
	Amount         int64
	BurnBeforeMint bool
}
