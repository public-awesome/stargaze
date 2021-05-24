package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Default parameter namespace
const (
	DefaultParamspace    string = ModuleName
	DefaultGenesisSender string = "stars1czlu4tvr3dg3ksuf8zak87eafztr2u004zyh5a"
	DefaultStakeDenom    string = "ustarx"
)

// Default vars
var (
	DefaultGenesisAllocation = sdk.NewInt64Coin(DefaultStakeDenom, 300_000_000_000_000)
)

// Parameter store keys
var (
	KeyGenesisAllocation = []byte("GenesisAllocation")
	KeyGenesisSender     = []byte("GenesisSender")
)

// NewParams creates a new Params object
func NewParams(
	genesisAllocation sdk.Coin,
	genesisSender string,
) Params {
	return Params{
		GenesisAllocation: genesisAllocation,
		GenesisSender:     genesisSender,
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultGenesisAllocation,
		DefaultGenesisSender,
	)
}
