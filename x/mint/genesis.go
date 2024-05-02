package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v14/x/mint/keeper"
	"github.com/public-awesome/stargaze/v14/x/mint/types"
)

// InitGenesis new mint genesis
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, ak types.AccountKeeper, data *types.GenesisState) {
	keeper.SetMinter(ctx, data.Minter)
	if err := keeper.SetParams(ctx, data.Params); err != nil {
		panic(err)
	}
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	minter, err := keeper.GetMinter(ctx)
	if err != nil {
		panic(err)
	}
	params, err := keeper.GetParams(ctx)
	if err != nil {
		panic(err)
	}
	return types.NewGenesisState(minter, params)
}
