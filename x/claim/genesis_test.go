package claim_test

import (
	"testing"

	keepertest "github.com/public-awesome/stargaze/testutil/keeper"
	"github.com/public-awesome/stargaze/x/claim"
	"github.com/public-awesome/stargaze/x/claim/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ClaimKeeper(t)
	claim.InitGenesis(ctx, *k, genesisState)
	got := claim.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
