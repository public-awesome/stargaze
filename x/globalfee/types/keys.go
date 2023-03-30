package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the module name.
	ModuleName = "globalfee"
	// StoreKey is the module KV storage prefix key.
	StoreKey = ModuleName
	// QuerierRoute is the querier route for the module.
	QuerierRoute = ModuleName
)

var (
	CodeAuthorizationPrefix     = []byte{0x00}
	ContractAuthorizationPrefix = []byte{0x01}
)

func GetCodeAuthorizationPrefix(codeId uint64) []byte {
	return append(CodeAuthorizationPrefix, i64tob(codeId)...)
}

func GetContractAuthorizationPrefix(contractAddress sdk.AccAddress) []byte {
	return append(CodeAuthorizationPrefix, contractAddress...)
}

func i64tob(val uint64) []byte {
	r := make([]byte, 8)
	for i := uint64(0); i < 8; i++ {
		r[i] = byte((val >> (i * 8)) & 0xff)
	}
	return r
}
