package v046

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/v11/x/alloc/types"
)

// MigrateStore performs in-place store migrations from v2 to v3
// The migration includes:
//
// - Setting the KeyIcentiveRewardsReceiver param in the paramstore
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, paramstore paramtypes.Subspace) error {
	migrateParamsStore(ctx, paramstore)
	return nil
}

func migrateParamsStore(ctx sdk.Context, paramstore paramtypes.Subspace) {
	defaultParams := types.DefaultParams()
	if paramstore.HasKeyTable() {
		paramstore.Set(ctx, types.KeyIcentiveRewardsReceiver, defaultParams.WeightedIncentivesRewardsReceivers)
	} else {
		paramstore.WithKeyTable(types.ParamKeyTable())
		paramstore.Set(ctx, types.KeyIcentiveRewardsReceiver, defaultParams.WeightedIncentivesRewardsReceivers)
	}
}
