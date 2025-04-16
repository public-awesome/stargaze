package v4_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"

	storetypes "cosmossdk.io/store/types"
	"github.com/public-awesome/stargaze/v17/x/alloc"
	"github.com/public-awesome/stargaze/v17/x/alloc/exported"
	v4 "github.com/public-awesome/stargaze/v17/x/alloc/migrations/v4"
	"github.com/public-awesome/stargaze/v17/x/alloc/types"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(_ sdk.Context, ps exported.ParamSet) {
	*ps.(*types.Params) = ms.ps
}

func TestMigrateStore(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(alloc.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	store := ctx.KVStore(storeKey)

	legacySubspace := newMockSubspace(types.DefaultParams())
	require.NoError(t, v4.MigrateStore(ctx, storeKey, legacySubspace, cdc))

	var res types.Params
	bz := store.Get(types.ParamsKey)
	require.NoError(t, cdc.Unmarshal(bz, &res))
	require.Equal(t, legacySubspace.ps.DistributionProportions, res.DistributionProportions)
	require.True(t, legacySubspace.ps.SupplementAmount.Equal(res.SupplementAmount))
	require.Len(t, res.WeightedDeveloperRewardsReceivers, len(legacySubspace.ps.WeightedDeveloperRewardsReceivers))
	require.Len(t, res.WeightedIncentivesRewardsReceivers, len(legacySubspace.ps.WeightedIncentivesRewardsReceivers))
}
