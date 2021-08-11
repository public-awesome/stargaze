package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	curatingtypes "github.com/public-awesome/stargaze/x/curating/types"
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
)

var (
	// KeyPrefixStake 0x00 | vendor_id | post_id | delegator -> Stake
	KeyPrefixStake = []byte{0x00}
)

// PostKey is the key used to store stakes for a post
func PostKey(vendorID uint32, postID curatingtypes.PostID) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	return append(KeyPrefixStake, append(vendorIDBz, postID.Bytes()...)...)
}

// StakeKey is the key used to store a stake
func StakeKey(vendorID uint32, postID curatingtypes.PostID, delegator sdk.AccAddress) []byte {
	return append(PostKey(vendorID, postID), delegator.Bytes()...)
}

// Uint32ToBigEndian - marshals uint32 to a bigendian byte slice so it can be sorted
func uint32ToBigEndian(i uint32) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b, i)
	return b
}
