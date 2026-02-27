package cron

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v18/x/cron/keeper"
	"github.com/public-awesome/stargaze/v18/x/cron/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	params := genState.GetParams()
	if err := k.SetParams(ctx, params); err != nil {
		panic(err)
	}

	for _, addr := range genState.GetPrivilegedContractAddresses() {
		contractAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			panic(err)
		}
		err = k.SetPrivileged(ctx, contractAddr)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}
	genesis.Params = params

	k.IteratePrivileged(ctx, func(addr sdk.AccAddress) bool {
		genesis.PrivilegedContractAddresses = append(genesis.PrivilegedContractAddresses, addr.String())
		return false
	})
	return genesis
}
