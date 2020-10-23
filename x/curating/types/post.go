package types

import (
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

// func (cp CuratedPost) String() string {
// 	postID, err := strconv.ParseInt(string(cp.PostID), 10, 64)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return strconv.FormatInt(postID, 10)
// }

// String implements the stringer interface for Post
func (p *Post) String() string {
	// out, err := yaml.Marshal(p)
	// if err != nil {
	// 	return ""
	// }
	// return string(out)
	return "hello"
}
