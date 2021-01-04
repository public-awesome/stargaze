package types

import (
	"encoding/json"
	"time"

	"github.com/bwmarrin/snowflake"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewUpvote fills and Upvote struct
func NewUpvote(
	vendorID uint32, postIDBz []byte, curator, rewardAccount sdk.AccAddress,
	voteAmount sdk.Coin, curatedTime time.Time) Upvote {

	return Upvote{
		VendorID:      vendorID,
		PostID:        postIDBz,
		Curator:       curator.String(),
		RewardAccount: rewardAccount.String(),
		VoteAmount:    voteAmount,
		CuratedTime:   curatedTime,
	}
}

// Upvotes is a collection of Upvote objects
type Upvotes []Upvote

type upvotePretty struct {
	VendorID      uint32   `json:"vendor_id" yaml:"vendor_id"`
	PostID        string   `json:"post_id" yaml:"post_id"`
	Curator       string   `json:"curator" yaml:"curator"`
	RewardAccount string   `json:"reward_account" yaml:"reward_account"`
	VoteAmount    sdk.Coin `json:"vote_amount" yaml:"vote_amount"`
	CuratedTime   string   `json:"curated_time" yaml:"curated_time"`
}

// MarshalJSON defines custom encoding scheme
func (u Upvote) MarshalJSON() ([]byte, error) {
	var temp [8]byte
	copy(temp[:], u.PostID) // convert a postID byte slice into a fixed 8 byte array
	postID := snowflake.ParseIntBytes(temp)

	out, err := json.Marshal(upvotePretty{
		VendorID:      u.VendorID,
		PostID:        postID.String(),
		Curator:       u.Curator,
		RewardAccount: u.RewardAccount,
		VoteAmount:    u.VoteAmount,
		CuratedTime:   u.CuratedTime.Format(time.RFC3339),
	})
	if err != nil {
		return out, err
	}

	return out, nil
}
