package types

const (
	// ModuleName is the module name.
	ModuleName = "pauser"
	// StoreKey is the module KV storage prefix key.
	StoreKey = ModuleName
	// QuerierRoute is the querier route for the module.
	QuerierRoute = ModuleName
	// RouterKey is the message route.
	RouterKey = ModuleName
)

var (
	PausedContractPrefix = []byte{0x01}
	PausedCodeIDPrefix   = []byte{0x02}

	// ParamsKey stores the module params
	ParamsKey = []byte{0x03}
)
