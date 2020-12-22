package types

import (
	"encoding/binary"
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

// StakeKey is the key used to store a stake
func StakeKey(vendorID uint32, postIDBz []byte) []byte {
	vendorIDBz := uint32ToBigEndian(vendorID)
	return append(KeyPrefixStake, append(vendorIDBz, postIDBz...)...)
}

// Uint32ToBigEndian - marshals uint32 to a bigendian byte slice so it can be sorted
func uint32ToBigEndian(i uint32) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b, i)
	return b
}
