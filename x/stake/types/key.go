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

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
)

var (
	// 0x11 | vendor_id | post_id
	KeyPrefixPost = []byte{0x11}

	// 0x12 | format(expire_time) | vendor_id | post_id | stake_id
	KeyPrefixVotingDelegationQueue = []byte{0x12}

	// 0x13
	KeyIndexStakeID = []byte{0x13}
)

func PostKey(vendorID uint32, postID uint64) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	postIDBz := sdk.Uint64ToBigEndian(postID)
	return append(KeyPrefixPost, append(vendorIDBz, postIDBz...)...)
}

func VotingDelegationQueueKey(completionTime time.Time, vendorID uint32, postID, stakeID uint64) []byte {
	bz := sdk.Uint64ToBigEndian(stakeID)
	return append(votingDelegationQueuePostIDPrefix(completionTime, vendorID, postID), bz...)
}

func votingDelegationQueuePostIDPrefix(completionTime time.Time, vendorID uint32, postID uint64) []byte {
	bz := sdk.Uint64ToBigEndian(postID)
	return append(votingDelegationQueueVendorPrefix(completionTime, vendorID), bz...)
}

func votingDelegationQueueVendorPrefix(completionTime time.Time, vendorID uint32) []byte {
	return append(VotingDelegationQueueTimeKeyPrefix(completionTime), uint32ToBigEndian(vendorID)...)
}

func VotingDelegationQueueTimeKeyPrefix(completionTime time.Time) []byte {
	bz := sdk.FormatTimeBytes(completionTime)
	return append(KeyPrefixVotingDelegationQueue, bz...)
}

// marshals uint32 to a big endian byte slice so it can be sorted
func uint32ToBigEndian(i uint32) []byte {
	b := make([]byte, 3)
	binary.BigEndian.PutUint32(b, i)
	return b
}

// returns bytes to be used as a key for a given stake index.
func StakeIndexToKey(index uint64) []byte {
	return sdk.Uint64ToBigEndian(index)
}

// returns a stake index for a given byte key
func StakeIndexFromKey(key []byte) uint64 {
	return bigEndianToUint64(key)
}

// returns an uint64 from big endian encoded bytes. If encoding
// is empty, zero is returned.
func bigEndianToUint64(bz []byte) uint64 {
	if len(bz) == 0 {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}
