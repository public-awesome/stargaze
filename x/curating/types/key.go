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

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName

	// RewardPoolName is the name for the module account for the reward pool
	RewardPoolName = "reward_matching_pool"

	// VotingPoolName is the name of the global voting pool module account
	VotingPoolName = "voting_pool"

	// DefaultStakeDenom is the staking denom for the zone
	DefaultStakeDenom = "ufuel"
)

var (
	// 0x00 | vendor_id | post_id -> Post
	KeyPrefixPost = []byte{0x00}

	// 0x01 | vendor_id | post_id | curator -> Upvote
	KeyPrefixUpvote = []byte{0x01}

	// 0x02 | format(curation_end_time) -> []VPPair
	KeyPrefixCurationQueue = []byte{0x02}
)

var lenTime = len(sdk.FormatTimeBytes(time.Now()))

func PostKey(vendorID uint32, postIDHash []byte) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	return append(KeyPrefixPost, append(vendorIDBz, postIDHash...)...)
}

func UpvoteKey(vendorID uint32, postIDHash []byte, curator sdk.AccAddress) []byte {
	return append(UpvotePrefixKey(vendorID, postIDHash), curator.Bytes()...)
}

// UpvotePrefixKey 0x01|vendorID|postID|...
func UpvotePrefixKey(vendorID uint32, postIDHash []byte) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	return append(KeyPrefixUpvote, append(vendorIDBz, postIDHash...)...)
}

// CurationQueueByTimeKey gets the curation queue key by curation end time
func CurationQueueByTimeKey(curationEndTime time.Time) []byte {
	return append(KeyPrefixCurationQueue, sdk.FormatTimeBytes(curationEndTime)...)
}

// Uint32ToBigEndian - marshals uint32 to a bigendian byte slice so it can be sorted
func uint32ToBigEndian(i uint32) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b, i)
	return b
}
