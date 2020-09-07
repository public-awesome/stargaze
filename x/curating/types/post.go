package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewPost allocates and returns a new `Post` struct
func NewPost(
	vendorID uint32, postIDHash []byte, bodyHash []byte, creator,
	rewardAccount sdk.AccAddress, curatingEndTime time.Time) Post {

	return Post{
		VendorID:        vendorID,
		PostIDHash:      postIDHash,
		Creator:         creator,
		RewardAccount:   rewardAccount,
		BodyHash:        bodyHash,
		CuratingEndTime: curatingEndTime,
	}
}
