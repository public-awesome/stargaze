package v3

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/v16/x/alloc/types"
)

// MigrateStore performs in-place store migrations from v2 to v3
// The migration includes:
//
// - Setting the KeyIncentiveRewardsReceiver param in the paramstore
func MigrateStore(ctx sdk.Context, _ storetypes.StoreKey, _ codec.BinaryCodec, paramstore paramtypes.Subspace) error {
	migrateParamsStore(ctx, paramstore)
	return nil
}

func migrateParamsStore(ctx sdk.Context, paramstore paramtypes.Subspace) {
	defaultParams := types.DefaultParams()
	if paramstore.HasKeyTable() {
		paramstore.Set(ctx, types.KeyIncentiveRewardsReceiver, defaultParams.WeightedIncentivesRewardsReceivers)
		paramstore.Set(ctx, types.KeySupplementAmount, defaultParams.SupplementAmount)
	} else {
		paramstore.WithKeyTable(types.ParamKeyTable())
		paramstore.Set(ctx, types.KeyIncentiveRewardsReceiver, defaultParams.WeightedIncentivesRewardsReceivers)
		paramstore.Set(ctx, types.KeySupplementAmount, defaultParams.SupplementAmount)
	}
}
