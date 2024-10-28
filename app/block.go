package app

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PreBlocker application updates every pre block
func (a *App) PreBlocker(ctx sdk.Context, _ *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	return a.ModuleManager.PreBlock(ctx)
}

// Precommiter application updates every commit
func (a *App) Precommiter(ctx sdk.Context) {
	if err := a.ModuleManager.Precommit(ctx); err != nil {
		panic(err)
	}
}

// PrepareCheckStater application updates every commit
func (a *App) PrepareCheckStater(ctx sdk.Context) {
	if err := a.ModuleManager.PrepareCheckState(ctx); err != nil {
		panic(err)
	}
}

// BeginBlocker application updates every begin block
func (a *App) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	return a.ModuleManager.BeginBlock(ctx)
}

// EndBlocker application updates every end block
func (a *App) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	return a.ModuleManager.EndBlock(ctx)
}
