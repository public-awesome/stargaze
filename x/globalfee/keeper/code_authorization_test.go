package keeper_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v9/testutil/keeper"
	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
	"github.com/stretchr/testify/require"
)

func Test_CodeAuthorization(t *testing.T) {
	k, ctx := keeper.GlobalFeeKeeper(t)
	ca := types.CodeAuthorization{
		CodeId:  1,
		Methods: []string{"mint", "list"},
	}

	t.Run("store invalid ca", func(t *testing.T) {
		err := k.SetCodeAuthorization(ctx, types.CodeAuthorization{
			CodeId:  1,
			Methods: []string{},
		})
		require.Error(t, err)
	})

	t.Run("authorization doesnt exist", func(t *testing.T) {
		_, found := k.GetCodeAuthorization(ctx, ca.CodeId)
		require.False(t, found)
	})

	t.Run("store authorization", func(t *testing.T) {
		err := k.SetCodeAuthorization(ctx, ca)
		require.NoError(t, err)

		_, found := k.GetCodeAuthorization(ctx, ca.CodeId)
		require.True(t, found)
	})

	t.Run("delete authorization", func(t *testing.T) {
		_, found := k.GetCodeAuthorization(ctx, ca.CodeId)
		require.True(t, found)

		k.DeleteCodeAuthorization(ctx, ca.CodeId)

		_, found = k.GetCodeAuthorization(ctx, ca.CodeId)
		require.False(t, found)
	})
}
