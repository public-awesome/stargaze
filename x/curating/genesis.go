package curating

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/keeper"
	"github.com/public-awesome/stakebird/x/curating/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) (res []abci.ValidatorUpdate) {
	k.SetParams(ctx, data.Params)

	if k.GetRewardPoolBalance(ctx).IsZero() {
		err := k.InitializeRewardPool(ctx, k.GetParams(ctx).InitialRewardPool)
		if err != nil {
			panic(err)
		}
	}

	return []abci.ValidatorUpdate{}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	return types.NewGenesisState(k.GetParams(ctx))
}
