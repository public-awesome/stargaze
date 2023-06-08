package globalfee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v11/x/globalfee/keeper"
	"github.com/public-awesome/stargaze/v11/x/globalfee/types"
)

// InitGenesis initializes the module genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	params := genState.GetParams()
	k.SetParams(ctx, params)

	for _, ca := range genState.GetCodeAuthorizations() {
		err := k.SetCodeAuthorization(ctx, ca)
		if err != nil {
			panic(err)
		}
	}

	for _, ca := range genState.GetContractAuthorizations() {
		err := k.SetContractAuthorization(ctx, ca)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis exports the module genesis for the current block.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	params := k.GetParams(ctx)
	genesis.Params = params
	k.IterateCodeAuthorizations(ctx, func(ca types.CodeAuthorization) bool {
		genesis.CodeAuthorizations = append(genesis.CodeAuthorizations, ca)
		return false
	})
	k.IterateContractAuthorizations(ctx, func(ca types.ContractAuthorization) bool {
		genesis.ContractAuthorizations = append(genesis.ContractAuthorizations, ca)
		return false
	})
	return genesis
}
