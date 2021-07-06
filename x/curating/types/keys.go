package types

import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "curating"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierKey to be used for querierer msgs
	QuerierKey = ModuleName

	// RewardPoolName is the name for the module account for the reward pool
	RewardPoolName = "reward_matching_pool"

	// VotingPoolName is the name of the global voting pool module account
	VotingPoolName = "voting_pool"

	// DefaultStakeDenom is the staking denom for the zone
	DefaultStakeDenom = "ustarx"

	// DefaultVoteDenom is the denom for quadratic voting
	DefaultVoteDenom = "ucredits"
)

var (
	// KeyPrefixPost 0x00 | vendor_id | post_id -> Post
	KeyPrefixPost = []byte{0x00}

	// KeyPrefixUpvote 0x01 | vendor_id | post_id | curator -> Upvote
	KeyPrefixUpvote = []byte{0x01}

	// KeyPrefixCurationQueue 0x02 | format(curation_end_time) -> []VPPair
	KeyPrefixCurationQueue = []byte{0x02}

	// PostIDKey for native posts (vendor = 0)
	PostIDKey = []byte{0x03}
)

// PostsKey is an index on all posts for a vendor
func PostsKey(vendorID uint32) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	return append(KeyPrefixPost, vendorIDBz...)
}

// PostKey is the key used to store a post
func PostKey(vendorID uint32, postID PostID) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	return append(KeyPrefixPost, append(vendorIDBz, postID.Bytes()...)...)
}

// UpvoteKey key is the key used to store an upvote
func UpvoteKey(vendorID uint32, postID PostID, curator sdk.AccAddress) []byte {
	return append(UpvotePrefixKey(vendorID, postID), curator.Bytes()...)
}

// UpvotePrefixKey 0x01|vendorID|postID|...
func UpvotePrefixKey(vendorID uint32, postID PostID) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	return append(KeyPrefixUpvote, append(vendorIDBz, postID.Bytes()...)...)
}

// CurationQueueByTimeKey gets the curation queue key by curation end time
func CurationQueueByTimeKey(curationEndTime time.Time) []byte {
	return append(KeyPrefixCurationQueue, sdk.FormatTimeBytes(curationEndTime)...)
}

// GetPostIDBytes returns the byte representation of the postlID
func GetPostIDBytes(postID uint64) (postIDBz []byte) {
	postIDBz = make([]byte, 8)
	binary.BigEndian.PutUint64(postIDBz, postID)
	return
}

// GetPostIDFromBytes returns postID in uint64 format from a byte array
func GetPostIDFromBytes(bz []byte) (postID uint64) {
	return binary.BigEndian.Uint64(bz)
}

// Uint32ToBigEndian - marshals uint32 to a bigendian byte slice so it can be sorted
func uint32ToBigEndian(i uint32) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b, i)
	return b
}
