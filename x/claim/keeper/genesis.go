package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v6/x/claim/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) []abci.ValidatorUpdate {
	k.CreateModuleAccount(ctx, data.ModuleAccountBalance)
	if data.Params.AirdropEnabled && data.Params.AirdropStartTime.Equal(time.Time{}) {
		data.Params.AirdropStartTime = ctx.BlockTime()
	}
	err := k.SetClaimRecords(ctx, data.ClaimRecords)
	if err != nil {
		panic(err)
	}
	k.SetParams(ctx, data.Params)
	return nil
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	params := k.GetParams(ctx)
	genesis.ModuleAccountBalance = k.GetModuleAccountBalance(ctx)
	genesis.Params = params
	genesis.ClaimRecords = k.ClaimRecords(ctx)
	return genesis
}
