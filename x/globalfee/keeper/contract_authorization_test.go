package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v16/testutil/keeper"
	"github.com/public-awesome/stargaze/v16/testutil/sample"
	"github.com/public-awesome/stargaze/v16/x/globalfee/types"
	"github.com/stretchr/testify/require"
)

func Test_ContractAuthorization(t *testing.T) {
	k, ctx := keeper.GlobalFeeKeeper(t)
	ca := types.ContractAuthorization{
		ContractAddress: "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du",
		Methods:         []string{"mint", "list"},
	}

	t.Run("store invalid ca", func(t *testing.T) {
		err := k.SetContractAuthorization(ctx, types.ContractAuthorization{
			ContractAddress: "👻",
			Methods:         []string{"mint", "list"},
		})
		require.Error(t, err)
	})

	t.Run("non existing contract address", func(t *testing.T) {
		err := k.SetContractAuthorization(ctx, types.ContractAuthorization{
			ContractAddress: sample.AccAddress().String(),
			Methods:         []string{"mint", "list"},
		})
		require.Error(t, err)
	})

	t.Run("authorization doesn't exist", func(t *testing.T) {
		found := k.HasContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))
		require.False(t, found)
	})

	t.Run("store authorization", func(t *testing.T) {
		err := k.SetContractAuthorization(ctx, ca)
		require.NoError(t, err)

		found := k.HasContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))
		require.True(t, found)
	})

	t.Run("delete authorization", func(t *testing.T) {
		found := k.HasContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))
		require.True(t, found)

		err := k.DeleteContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))
		require.NoError(t, err)

		found = k.HasContractAuthorization(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress))
		require.False(t, found)
	})

	t.Run("iterate contract authorization", func(t *testing.T) {
		count := 0
		k.IterateContractAuthorizations(ctx, func(ca types.ContractAuthorization) bool {
			count++
			return false
		})
		require.Equal(t, 0, count)

		err := k.SetContractAuthorization(ctx, ca)
		require.NoError(t, err)
		err = k.SetContractAuthorization(ctx, types.ContractAuthorization{
			ContractAddress: "cosmos1hfml4tzwlc3mvynsg6vtgywyx00wfkhrtpkx6t",
			Methods:         ca.GetMethods(),
		})
		require.NoError(t, err)
		err = k.SetContractAuthorization(ctx, types.ContractAuthorization{
			ContractAddress: "cosmos144sh8vyv5nqfylmg4mlydnpe3l4w780jsrmf4k",
			Methods:         []string{"test"},
		})
		require.NoError(t, err)

		count = 0
		k.IterateContractAuthorizations(ctx, func(ca types.ContractAuthorization) bool {
			count++
			return false
		})
		require.Equal(t, 3, count)
	})
}
