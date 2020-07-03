package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPost(
	vendorID uint32, postIDHash []byte, bodyHash []byte, creator,
	rewardAccount sdk.AccAddress, deposit sdk.Coin, curatingEndTime time.Time) Post {

	return Post{
		VendorID:        vendorID,
		PostIDHash:      postIDHash,
		Creator:         creator,
		RewardAccount:   rewardAccount,
		BodyHash:        bodyHash,
		Deposit:         deposit,
		CuratingEndTime: curatingEndTime,
	}
}
