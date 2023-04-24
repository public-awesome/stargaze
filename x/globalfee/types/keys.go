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
	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_globalfee"
	// RouterKey is the message route f
	RouterKey = ModuleName
)

var (
	CodeAuthorizationPrefix     = []byte{0x00}
	ContractAuthorizationPrefix = []byte{0x01}
)

func GetCodeAuthorizationPrefix(codeId uint64) []byte {
	return append(CodeAuthorizationPrefix, sdk.Uint64ToBigEndian(codeId)...)
}

func GetContractAuthorizationPrefix(contractAddress sdk.AccAddress) []byte {
	return append(ContractAuthorizationPrefix, contractAddress...)
}
