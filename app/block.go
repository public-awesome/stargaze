package app

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PreBlocker application updates every pre block
func (a *App) PreBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	// call app's preblocker first in case there is changes made on upgrades
	// that can modify state and lead to serialization/deserialization issues
	resp, err := a.ModuleManager.PreBlock(ctx)
	if err != nil {
		return resp, err
	}

	// oracle preblocker sends empty response pre block so it can ignored
	_, err = a.oraclePreBlocker(ctx, req)
	if err != nil {
		return &sdk.ResponsePreBlock{}, err
	}

	// return resp from app's preblocker which can return consensus param changed flag
	return resp, nil
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
