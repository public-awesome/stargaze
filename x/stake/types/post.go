package types

import (
	"time"
)

func NewPost(id uint64, vendorID uint32, body string, votingPeriod time.Duration, votingStartTime time.Time) Post {
	return Post{
		ID:              id,
		VendorID:        vendorID,
		Body:            body,
		VotingPeriod:    votingPeriod,
		VotingStartTime: votingStartTime,
	}
}
