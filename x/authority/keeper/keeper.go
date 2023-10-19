package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/v13/x/authority/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		paramstore paramtypes.Subspace
		router     *baseapp.MsgServiceRouter
		authority  string // this should be the x/gov module account
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	router *baseapp.MsgServiceRouter,
	authority string,
) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramstore: ps,
		router:     router,
		authority:  authority,
	}
}

// GetAuthority returns the x/wasm module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Router returns the adminmodule Keeper's Router
func (k Keeper) Router() *baseapp.MsgServiceRouter {
	return k.router
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ModuleLogger(ctx)
}

func ModuleLogger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
