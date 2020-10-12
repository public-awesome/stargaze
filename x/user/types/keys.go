package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "user"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierKey to be used for querierer msgs
	QuerierKey = ModuleName
)

var (
	// KeyPrefixVoucher 0x00 | voucher | vouched -> Vouch
	KeyPrefixVoucher = []byte{0x00}

	// KeyPrefixVouched 0x01 | vouched -> Vouch
	KeyPrefixVouched = []byte{0x01}
)

// VoucherKey gets the key four voucher | vouched
func VoucherKey(voucher, vouched sdk.AccAddress) []byte {
	return append(KeyPrefixVoucher, append(voucher.Bytes(), vouched.Bytes()...)...)
}

// VoucherPrefixKey gets the key for a voucher
func VoucherPrefixKey(voucher sdk.AccAddress) []byte {
	return append(KeyPrefixVoucher, voucher.Bytes()...)
}

// VouchedKey gets the key for a vouched address
func VouchedKey(vouched sdk.AccAddress) []byte {
	return append(KeyPrefixVouched, vouched.Bytes()...)
}
