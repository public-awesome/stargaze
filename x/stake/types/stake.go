package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewStake allocates and returns a new `Stake` struct
func NewStake(vendorID uint32, postIDBz []byte, bodyHash []byte, creator,
	rewardAccount sdk.AccAddress, curatingEndTime time.Time) Stake {

	return Stake{
		VendorID:        vendorID,
		PostID:          postIDBz,
		Creator:         creator.String(),
		RewardAccount:   rewardAccount.String(),
		BodyHash:        bodyHash,
		CuratingEndTime: curatingEndTime,
	}
}
