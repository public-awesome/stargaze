package types

const (
	// ModuleName defines the module name
	ModuleName = "alloc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// Fairburn pool name
	FairburnPoolName = "fairburn_pool"

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_alloc"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
