package types

import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "stake"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierKey to be used for querierer msgs
	QuerierKey = ModuleName

	// DefaultStakeDenom is the staking denom for the zone
	DefaultStakeDenom = "ustb"
)

var (
	// KeyPrefixPost 0x00 | vendor_id | post_id -> Post
	KeyPrefixPost = []byte{0x00}

	// KeyPrefixCurationQueue 0x02 | format(curation_end_time) -> []VPPair
	KeyPrefixCurationQueue = []byte{0x02}
)

// PostsKey is an index on all posts for a vendor
func PostsKey(vendorID uint32) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	return append(KeyPrefixPost, vendorIDBz...)
}

// PostKey is the key used to store a post
func PostKey(vendorID uint32, postIDBz []byte) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	return append(KeyPrefixPost, append(vendorIDBz, postIDBz...)...)
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
