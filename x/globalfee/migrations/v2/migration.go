package v2

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v17/x/globalfee/exported"
	"github.com/public-awesome/stargaze/v17/x/globalfee/types"
)

// MigrateStore migrates the x/globalfee module state from the consensus version 1 to
// version 2. Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/globalfee
// module state.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, legacySubspace exported.Subspace, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	var currParams types.Params
	legacySubspace.GetParamSet(ctx, &currParams)

	if err := currParams.Validate(); err != nil {
		return err
	}

	bz, err := cdc.Marshal(&currParams)
	if err != nil {
		return err
	}

	store.Set(types.ParamsKey, bz)

	return nil
}
