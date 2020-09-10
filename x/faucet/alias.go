package faucet

import (
	"github.com/cosmos/modules/incubator/faucet/internal/keeper"
	"github.com/cosmos/modules/incubator/faucet/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey

	MAINNET = "mainnet"
	TESTNET = "testnet"
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper = keeper.Keeper
)
