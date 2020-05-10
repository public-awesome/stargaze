package types

import (
	"time"
)

func NewPost(id, vendorID uint64, body string, votingPeriod time.Duration, votingStartTime time.Time) Post {
	return Post{
		ID:              id,
		VendorID:        vendorID,
		Body:            body,
		VotingPeriod:    votingPeriod,
		VotingStartTime: votingStartTime,
	}
}
