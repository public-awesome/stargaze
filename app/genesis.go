package app

import (
	"encoding/json"

	codecstd "github.com/cosmos/cosmos-sdk/codec/std"
)

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() GenesisState {
	cdc := codecstd.MakeCodec(ModuleBasics)
	return ModuleBasics.DefaultGenesis(cdc)
}
