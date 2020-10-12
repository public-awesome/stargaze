package types

const (
	// ModuleName is the name of the module
	ModuleName = "faucet"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierKey to be used for querierer msgs
	QuerierKey = ModuleName
)

var (

	// FaucetStoreKey is the key used to store the default faucet key
	FaucetStoreKey = []byte("DefaultFaucetStoreKey")
)
