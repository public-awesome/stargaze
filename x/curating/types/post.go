package types

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewPost allocates and returns a new `Post` struct
func NewPost(
	vendorID uint32, postIDBz []byte, bodyHash []byte, creator,
	rewardAccount sdk.AccAddress, curatingEndTime time.Time) Post {

	return Post{
		VendorID:        vendorID,
		PostID:          postIDBz,
		Creator:         creator.String(),
		RewardAccount:   rewardAccount.String(),
		BodyHash:        bodyHash,
		CuratingEndTime: curatingEndTime,
	}
}

// CuratedPost is an application-specific wrapper around a `Post`
type CuratedPost struct {
	*Post
}

func (cp CuratedPost) String() string {
	postID, err := strconv.ParseInt(string(cp.PostID), 10, 64)
	if err != nil {
		panic(err)
	}

	return strconv.FormatInt(postID, 10)

	// return cp.String()
}
