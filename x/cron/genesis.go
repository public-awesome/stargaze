package cron

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v8/x/cron/keeper"
	"github.com/public-awesome/stargaze/v8/x/cron/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, addr := range genState.GetPrivilegedContractAddresses() {
		contractAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			panic(err)
		}
		if k.HasContractInfo(ctx, contractAddr) {
			k.SetPrivileged(ctx, contractAddr)
		}
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	k.IteratePrivileged(ctx, func(addr sdk.AccAddress) bool {
		genesis.PrivilegedContractAddresses = append(genesis.PrivilegedContractAddresses, addr.String())
		return false
	})
	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
