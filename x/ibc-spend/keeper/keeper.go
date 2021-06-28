package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/public-awesome/stargaze/x/ibc-spend/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc            codec.Marshaler
		storeKey       sdk.StoreKey
		memKey         sdk.StoreKey
		ak             types.AccountKeeper
		transferKeeper types.TransferKeeper
		distrKeeper    types.DistributionKeeper
		bankKeeper     types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	ak types.AccountKeeper,
	transferKeeper types.TransferKeeper,
	distrKeeper types.DistributionKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {

	// ensure module account is set
	// if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
	// panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	// }

	return &Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		ak:             ak,
		transferKeeper: transferKeeper,
		distrKeeper:    distrKeeper,
		bankKeeper:     bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetIBCSpendAccount returns the ModuleAccount
func (k Keeper) GetIBCSpendAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.ak.GetModuleAccount(ctx, types.ModuleName)
}
