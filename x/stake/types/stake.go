package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	curatingtypes "github.com/public-awesome/stargaze/x/curating/types"
)

// NewStake allocates and returns a new `Stake` struct
func NewStake(
	vendorID uint32, postID curatingtypes.PostID,
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
