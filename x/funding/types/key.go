package types

const (
	// ModuleName is the name of the module
	ModuleName = "funding"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName

	Counterparty = "ibczeroxfer"
	ReserveDenom = "stake"
	Denom        = "ufuel"
)

var (
	// ChannelStoreKeyPrefix is a prefix for storing channel
	ChannelStoreKeyPrefix = []byte{0x03}
)

// ChannelStoreKey is a function to generate key for each verified channel in store
func ChannelStoreKey(chainName, channelPort string) []byte {
	buf := append(ChannelStoreKeyPrefix, []byte(chainName)...)
	buf = append(buf, []byte(channelPort)...)
	return buf
}
