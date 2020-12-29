package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewStake allocates and returns a new `Stake` struct
func NewStake(validator sdk.ValAddress, amount sdk.Int) Stake {

	return Stake{
		Validator: validator.String(),
		Amount:    amount,
	}
}
