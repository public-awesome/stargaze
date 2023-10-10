package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "cron"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cron"
)

var PrivilegedContractsPrefix = []byte{0x01}
var ParamsKey = []byte{0x02}

func PrivilegedContractsKey(contractAddr sdk.AccAddress) []byte {
	return append(PrivilegedContractsPrefix, contractAddr...)
}
