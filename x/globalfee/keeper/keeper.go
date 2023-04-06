package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
)

// Keeper provides module state operations.
type Keeper struct {
	cdc        codec.Codec
	paramStore paramTypes.Subspace
}

// NewKeeper creates a new Keeper instance.
func NewKeeper(cdc codec.Codec, ps paramTypes.Subspace) Keeper {
	return Keeper{
		cdc:        cdc,
		paramStore: ps,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
