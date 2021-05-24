package dao

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/dao/keeper"
	"github.com/public-awesome/stargaze/x/dao/types"
)

// InitGenesis initializes the dao module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, dk types.DistKeeper, bk types.BankKeeper, genState types.GenesisState) {
	funder, err := sdk.AccAddressFromBech32(genState.Params.Funder)
	if err != nil {
		panic(err)
	}
	daoFund := bk.GetAllBalances(ctx, funder)

	err = dk.FundCommunityPool(ctx, daoFund, funder)
	if err != nil {
		panic(err)
	}

	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the dao module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
