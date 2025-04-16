package keeper_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v17/testutil/keeper"
	"github.com/public-awesome/stargaze/v17/x/globalfee/types"
	"github.com/stretchr/testify/require"
)

func Test_CodeAuthorization(t *testing.T) {
	k, ctx := keeper.GlobalFeeKeeper(t)
	ca := types.CodeAuthorization{
		CodeID:  1,
		Methods: []string{"mint", "list"},
	}

	t.Run("store invalid ca", func(t *testing.T) {
		err := k.SetCodeAuthorization(ctx, types.CodeAuthorization{
			CodeID:  1,
			Methods: []string{},
		})
		require.Error(t, err)
	})

	t.Run("code id does not exist", func(t *testing.T) {
		err := k.SetCodeAuthorization(ctx, types.CodeAuthorization{
			CodeID:  10,
			Methods: []string{"*"},
		})
		require.Error(t, err)
	})

	t.Run("authorization doesn't exist", func(t *testing.T) {
		found := k.HasCodeAuthorization(ctx, ca.CodeID)
		require.False(t, found)
	})

	t.Run("store authorization", func(t *testing.T) {
		err := k.SetCodeAuthorization(ctx, ca)
		require.NoError(t, err)

		found := k.HasCodeAuthorization(ctx, ca.CodeID)
		require.True(t, found)
	})

	t.Run("delete authorization", func(t *testing.T) {
		found := k.HasCodeAuthorization(ctx, ca.CodeID)
		require.True(t, found)

		err := k.DeleteCodeAuthorization(ctx, ca.CodeID)
		require.NoError(t, err)

		found = k.HasCodeAuthorization(ctx, ca.CodeID)
		require.False(t, found)
	})

	t.Run("iterate code authorization", func(t *testing.T) {
		count := 0
		k.IterateCodeAuthorizations(ctx, func(ca types.CodeAuthorization) bool {
			count++
			return false
		})
		require.Equal(t, 0, count)

		err := k.SetCodeAuthorization(ctx, ca)
		require.NoError(t, err)
		err = k.SetCodeAuthorization(ctx, types.CodeAuthorization{
			CodeID:  2,
			Methods: ca.GetMethods(),
		})
		require.NoError(t, err)
		err = k.SetCodeAuthorization(ctx, types.CodeAuthorization{
			CodeID:  3,
			Methods: ca.GetMethods(),
		})
		require.NoError(t, err)

		count = 0
		k.IterateCodeAuthorizations(ctx, func(ca types.CodeAuthorization) bool {
			count++
			return false
		})
		require.Equal(t, 3, count)
	})
}
