package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPost(
	vendorID uint32, postID string, bodyHash []byte, creator,
	rewardAccount sdk.AccAddress, deposit sdk.Coin, curatingEndTime time.Time) Post {

	return Post{
		VendorID:        vendorID,
		PostID:          postID,
		Creator:         creator,
		RewardAccount:   rewardAccount,
		BodyHash:        bodyHash,
		Deposit:         deposit,
		CuratingEndTime: curatingEndTime,
	}
}
