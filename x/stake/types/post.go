package types

import (
	"time"

	"gopkg.in/yaml.v2"
)

// key: vendor_id | post_id
type Post struct {
	ID              uint64        `json:"id" yaml:"id"`
	VendorID        uint32        `json:"vendor_id" yaml:"vendor_id"`
	Body            string        `json:"body" yaml:"body"`
	VotingPeriod    time.Duration `json:"voting_period" yaml:"voting_period"`
	VotingStartTime time.Time     `json:"voting_start_time" yaml:"voting_start_time"`
}

func NewPost(id uint64, vendorID uint32, body string, votingPeriod time.Duration, votingStartTime time.Time) Post {
	return Post{
		ID:              id,
		VendorID:        vendorID,
		Body:            body,
		VotingPeriod:    votingPeriod,
		VotingStartTime: votingStartTime,
	}
}

func (p Post) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
