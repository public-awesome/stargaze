package v2_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"

	storetypes "cosmossdk.io/store/types"
	"github.com/public-awesome/stargaze/v16/x/mint"
	"github.com/public-awesome/stargaze/v16/x/mint/exported"
	v2 "github.com/public-awesome/stargaze/v16/x/mint/migrations/v2"
	"github.com/public-awesome/stargaze/v16/x/mint/types"
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
	encCfg := moduletestutil.MakeTestEncodingConfig(mint.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	store := ctx.KVStore(storeKey)

	legacySubspace := newMockSubspace(types.DefaultParams())
	require.NoError(t, v2.MigrateStore(ctx, storeKey, legacySubspace, cdc))

	var res types.Params
	bz := store.Get(types.ParamsKey)
	require.NoError(t, cdc.Unmarshal(bz, &res))
	require.Equal(t, legacySubspace.ps.BlocksPerYear, res.BlocksPerYear)
	require.Equal(t, legacySubspace.ps.InitialAnnualProvisions, res.InitialAnnualProvisions)
	require.Equal(t, legacySubspace.ps.MintDenom, res.MintDenom)
	require.Equal(t, legacySubspace.ps.ReductionFactor, res.ReductionFactor)
}
