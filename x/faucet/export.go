package faucet

import (
	"github.com/public-awesome/stakebird/x/faucet/internal/keeper"
	"github.com/public-awesome/stakebird/x/faucet/internal/types"
)

// exported consts
const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

// exported vars
var (
	NewKeeper = keeper.NewKeeper
)

// Keeper exports internal keeper for wiring
type (
	Keeper = keeper.Keeper
)
