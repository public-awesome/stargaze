package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPost(
	bodyHash []byte, creator, rewardAccount sdk.AccAddress,
	deposit sdk.Coin, curatingEndTime time.Time) Post {

	return Post{
		Creator:         creator,
		RewardAccount:   rewardAccount,
		BodyHash:        bodyHash,
		Deposit:         deposit,
		CuratingEndTime: curatingEndTime,
	}
}
