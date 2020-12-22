package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stakebird/x/stake/types"
)

// Keeper of the x/curating store
type Keeper struct {
	storeKey       sdk.StoreKey
	cdc            codec.BinaryMarshaler
	accountKeeper  types.AccountKeeper
	bankKeeper     types.BankKeeper
	curationKeeper types.CurationKeeper
	paramstore     paramtypes.Subspace
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, ak types.AccountKeeper,
	ck types.CurationKeeper, bk types.BankKeeper, ps paramtypes.Subspace) Keeper {

	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(ParamKeyTable())
	}

	// ensure reward pool module account is set
	if addr := ak.GetModuleAddress(types.RewardPoolName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.RewardPoolName))
	}

	keeper := Keeper{
		storeKey:       key,
		cdc:            cdc,
		accountKeeper:  ak,
		bankKeeper:     bk,
		curatingKeeper: ck,
		paramstore:     ps,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) validateVendorID(ctx sdk.Context, vendorID uint32) error {
	if vendorID > k.GetParams(ctx).MaxVendors {
		return sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid vendor_id %d", vendorID))
	}

	return nil
}
