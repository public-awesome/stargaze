package types

import (
	"time"
)

func NewPost(id, vendorID uint64, body string, voteEnd time.Time) Post {
	return Post{
		ID:       id,
		VendorID: vendorID,
		Body:     body,
		VoteEnd:  voteEnd,
	}
}
