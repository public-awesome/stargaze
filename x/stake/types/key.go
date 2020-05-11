package types

import (
	"encoding/binary"
	"fmt"
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

var lenTime = len(sdk.FormatTimeBytes(time.Now()))

func PostKey(vendorID, postID uint64) []byte {
	vendorIDBz := sdk.Uint64ToBigEndian(vendorID)
	postIDBz := sdk.Uint64ToBigEndian(postID)
	return append(KeyPrefixPost, append(vendorIDBz, postIDBz...)...)
}

func VotingDelegationQueueKey(endTime time.Time, vendorID, postID, stakeID uint64) []byte {
	bz := sdk.Uint64ToBigEndian(stakeID)
	return append(votingDelegationQueuePostIDPrefix(endTime, vendorID, postID), bz...)
}

func votingDelegationQueuePostIDPrefix(endTime time.Time, vendorID, postID uint64) []byte {
	bz := sdk.Uint64ToBigEndian(postID)
	return append(votingDelegationQueueVendorPrefix(endTime, vendorID), bz...)
}

func votingDelegationQueueVendorPrefix(endTime time.Time, vendorID uint64) []byte {
	return append(VotingDelegationQueueTimeKeyPrefix(endTime), sdk.Uint64ToBigEndian(vendorID)...)
}

func VotingDelegationQueueTimeKeyPrefix(endTime time.Time) []byte {
	bz := sdk.FormatTimeBytes(endTime)
	return append(KeyPrefixVotingDelegationQueue, bz...)
}

// returns bytes to be used as a key for a given stake index.
func StakeIndexToKey(index uint64) []byte {
	return sdk.Uint64ToBigEndian(index)
}

// returns a stake index for a given byte key
func StakeIndexFromKey(key []byte) uint64 {
	return bigEndianToUint64(key)
}

func SplitVotingDelegationQueueKey(key []byte) (endTime time.Time, vendorID, postID, stakeID uint64) {
	lenVendorID := 8
	lenPostID := 8
	lenStakeID := 8

	if len(key[1:]) != lenTime+lenVendorID+lenPostID+lenStakeID {
		panic(fmt.Sprintf("unexpected key length (%d ≠ %d)", len(key[1:]), lenTime+24))
	}

	endTime, err := sdk.ParseTimeBytes(key[1 : 1+lenTime])
	if err != nil {
		panic(err)
	}

	vendorID = binary.BigEndian.Uint64(key[1+lenTime : 1+lenTime+lenVendorID])
	postID = binary.BigEndian.Uint64(key[1+lenTime+lenVendorID : 1+lenTime+lenVendorID+lenPostID])
	stakeID = binary.BigEndian.Uint64(key[1+lenTime+lenVendorID+lenPostID:])

	return endTime, vendorID, postID, stakeID
}

// returns an uint64 from big endian encoded bytes. If encoding
// is empty, one is returned.
func bigEndianToUint64(bz []byte) uint64 {
	if len(bz) == 0 {
		// start with an index of 1 (easier debugging)
		return 1
	}

	return binary.BigEndian.Uint64(bz)
}
