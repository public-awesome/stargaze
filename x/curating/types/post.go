package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPost(id, vendorID uint64, hash string, creator sdk.AccAddress, stake sdk.Coin, curationEndTime time.Time) Post {
	return Post{
		Creator:         creator,
		Hash:            hash,
		Stake:           stake,
		CuratingEndTime: curationEndTime,
	}
}
