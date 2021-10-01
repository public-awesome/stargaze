package alloc_test

import (
	"testing"

	keepertest "github.com/public-awesome/stargaze/testutil/keeper"
	"github.com/public-awesome/stargaze/x/alloc"
	"github.com/public-awesome/stargaze/x/alloc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AllocKeeper(t)
	alloc.InitGenesis(ctx, *k, genesisState)
	got := alloc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
