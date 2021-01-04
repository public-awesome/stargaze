package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewStake allocates and returns a new `Stake` struct
func NewStake(
	vendorID uint32, postID []byte,
	delegator sdk.AccAddress, validator sdk.ValAddress, amount sdk.Int) Stake {

	return Stake{
		VendorID:  vendorID,
		PostID:    postID,
		Delegator: delegator.String(),
		Validator: validator.String(),
		Amount:    amount,
	}
}

// Stakes is a collection of Stake 🥩
type Stakes []Stake
