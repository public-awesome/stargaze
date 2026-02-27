package keeper_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	keepertest "github.com/public-awesome/stargaze/v18/testutil/keeper"
	"github.com/public-awesome/stargaze/v18/x/pauser/keeper"
	"github.com/public-awesome/stargaze/v18/x/pauser/types"
	"github.com/stretchr/testify/require"
)

func seedPausedContracts(t *testing.T, ctx sdk.Context, k keeper.Keeper, addrs []string) {
	t.Helper()
	for _, addr := range addrs {
		err := k.SetPausedContract(ctx, types.PausedContract{
			ContractAddress: addr,
			PausedBy:        "pauser",
			PausedAt:        time.Now(),
		})
		require.NoError(t, err)
	}
}

func seedPausedCodeIDs(t *testing.T, ctx sdk.Context, k keeper.Keeper, ids []uint64) {
	t.Helper()
	for _, id := range ids {
		err := k.SetPausedCodeID(ctx, types.PausedCodeID{
			CodeID:   id,
			PausedBy: "pauser",
			PausedAt: time.Now(),
		})
		require.NoError(t, err)
	}
}

// generateAddresses creates N deterministic bech32 addresses for testing pagination.
func generateAddresses(n int) []string {
	addrs := make([]string, n)
	for i := 0; i < n; i++ {
		addr := sdk.AccAddress(fmt.Appendf(nil, "addr%020d", i))
		addrs[i] = addr.String()
	}
	return addrs
}

func TestQueryPausedContractsPagination(t *testing.T) {
	k, ctx := keepertest.PauserKeeper(t)
	qs := keeper.NewQueryServer(k)

	addrs := generateAddresses(5)
	seedPausedContracts(t, ctx, k, addrs)

	t.Run("all results with no pagination", func(t *testing.T) {
		resp, err := qs.PausedContracts(ctx, &types.QueryPausedContractsRequest{})
		require.NoError(t, err)
		require.Len(t, resp.PausedContracts, 5)
	})

	t.Run("limit returns correct page size", func(t *testing.T) {
		resp, err := qs.PausedContracts(ctx, &types.QueryPausedContractsRequest{
			Pagination: &query.PageRequest{Limit: 2},
		})
		require.NoError(t, err)
		require.Len(t, resp.PausedContracts, 2)
		require.NotEmpty(t, resp.Pagination.NextKey)
	})

	t.Run("cursor-based pagination walks all pages", func(t *testing.T) {
		var allContracts []types.PausedContract
		var nextKey []byte

		for {
			resp, err := qs.PausedContracts(ctx, &types.QueryPausedContractsRequest{
				Pagination: &query.PageRequest{Limit: 2, Key: nextKey},
			})
			require.NoError(t, err)
			allContracts = append(allContracts, resp.PausedContracts...)
			if len(resp.Pagination.NextKey) == 0 {
				break
			}
			nextKey = resp.Pagination.NextKey
		}
		require.Len(t, allContracts, 5)
	})

	t.Run("offset pagination", func(t *testing.T) {
		resp, err := qs.PausedContracts(ctx, &types.QueryPausedContractsRequest{
			Pagination: &query.PageRequest{Offset: 3, Limit: 10},
		})
		require.NoError(t, err)
		require.Len(t, resp.PausedContracts, 2) // 5 total - 3 offset = 2
	})

	t.Run("count_total returns total count", func(t *testing.T) {
		resp, err := qs.PausedContracts(ctx, &types.QueryPausedContractsRequest{
			Pagination: &query.PageRequest{Limit: 2, CountTotal: true},
		})
		require.NoError(t, err)
		require.Len(t, resp.PausedContracts, 2)
		require.Equal(t, uint64(5), resp.Pagination.Total)
	})

	t.Run("empty collection returns empty results", func(t *testing.T) {
		emptyK, emptyCtx := keepertest.PauserKeeper(t)
		emptyQs := keeper.NewQueryServer(emptyK)
		resp, err := emptyQs.PausedContracts(emptyCtx, &types.QueryPausedContractsRequest{
			Pagination: &query.PageRequest{Limit: 10},
		})
		require.NoError(t, err)
		require.Empty(t, resp.PausedContracts)
	})

	t.Run("reverse order", func(t *testing.T) {
		resp, err := qs.PausedContracts(ctx, &types.QueryPausedContractsRequest{
			Pagination: &query.PageRequest{Limit: 5, Reverse: true},
		})
		require.NoError(t, err)
		require.Len(t, resp.PausedContracts, 5)

		fwdResp, err := qs.PausedContracts(ctx, &types.QueryPausedContractsRequest{
			Pagination: &query.PageRequest{Limit: 5},
		})
		require.NoError(t, err)

		// first of forward should be last of reverse
		require.Equal(t,
			fwdResp.PausedContracts[0].ContractAddress,
			resp.PausedContracts[len(resp.PausedContracts)-1].ContractAddress,
		)
	})
}

func TestQueryPausedCodeIDsPagination(t *testing.T) {
	k, ctx := keepertest.PauserKeeper(t)
	qs := keeper.NewQueryServer(k)

	ids := []uint64{10, 20, 30, 40, 50}
	seedPausedCodeIDs(t, ctx, k, ids)

	t.Run("all results with no pagination", func(t *testing.T) {
		resp, err := qs.PausedCodeIDs(ctx, &types.QueryPausedCodeIDsRequest{})
		require.NoError(t, err)
		require.Len(t, resp.PausedCodeIds, 5)
	})

	t.Run("limit returns correct page size", func(t *testing.T) {
		resp, err := qs.PausedCodeIDs(ctx, &types.QueryPausedCodeIDsRequest{
			Pagination: &query.PageRequest{Limit: 2},
		})
		require.NoError(t, err)
		require.Len(t, resp.PausedCodeIds, 2)
		require.NotEmpty(t, resp.Pagination.NextKey)
	})

	t.Run("cursor-based pagination walks all pages", func(t *testing.T) {
		var allCodeIDs []types.PausedCodeID
		var nextKey []byte

		for {
			resp, err := qs.PausedCodeIDs(ctx, &types.QueryPausedCodeIDsRequest{
				Pagination: &query.PageRequest{Limit: 2, Key: nextKey},
			})
			require.NoError(t, err)
			allCodeIDs = append(allCodeIDs, resp.PausedCodeIds...)
			if len(resp.Pagination.NextKey) == 0 {
				break
			}
			nextKey = resp.Pagination.NextKey
		}
		require.Len(t, allCodeIDs, 5)

		// verify all original IDs are present
		foundIDs := make(map[uint64]bool)
		for _, pc := range allCodeIDs {
			foundIDs[pc.CodeID] = true
		}
		for _, id := range ids {
			require.True(t, foundIDs[id], "expected code ID %d in paginated results", id)
		}
	})

	t.Run("offset pagination", func(t *testing.T) {
		resp, err := qs.PausedCodeIDs(ctx, &types.QueryPausedCodeIDsRequest{
			Pagination: &query.PageRequest{Offset: 3, Limit: 10},
		})
		require.NoError(t, err)
		require.Len(t, resp.PausedCodeIds, 2) // 5 total - 3 offset = 2
	})

	t.Run("count_total returns total count", func(t *testing.T) {
		resp, err := qs.PausedCodeIDs(ctx, &types.QueryPausedCodeIDsRequest{
			Pagination: &query.PageRequest{Limit: 2, CountTotal: true},
		})
		require.NoError(t, err)
		require.Len(t, resp.PausedCodeIds, 2)
		require.Equal(t, uint64(5), resp.Pagination.Total)
	})

	t.Run("empty collection returns empty results", func(t *testing.T) {
		emptyK, emptyCtx := keepertest.PauserKeeper(t)
		emptyQs := keeper.NewQueryServer(emptyK)
		resp, err := emptyQs.PausedCodeIDs(emptyCtx, &types.QueryPausedCodeIDsRequest{
			Pagination: &query.PageRequest{Limit: 10},
		})
		require.NoError(t, err)
		require.Empty(t, resp.PausedCodeIds)
	})

	t.Run("reverse order", func(t *testing.T) {
		resp, err := qs.PausedCodeIDs(ctx, &types.QueryPausedCodeIDsRequest{
			Pagination: &query.PageRequest{Limit: 5, Reverse: true},
		})
		require.NoError(t, err)
		require.Len(t, resp.PausedCodeIds, 5)
		// uint64 keys are ordered ascending by default, so reverse should give 50, 40, 30, 20, 10
		require.Equal(t, uint64(50), resp.PausedCodeIds[0].CodeID)
		require.Equal(t, uint64(10), resp.PausedCodeIds[4].CodeID)
	})
}
