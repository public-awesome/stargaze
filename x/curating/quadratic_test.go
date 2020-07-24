package curating_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating"
	"github.com/stretchr/testify/require"
)

// 1 vote  = 1 vote credit
// 2 votes = 4 vote credits
// 3 votes = 9 vote credits
// 4 votes = 16 vote credits
// 5 votes = 25 vote credits
func TestQVF(t *testing.T) {
	q := curating.NewQVFData()
	q, err := q.TallyVote(sdk.NewInt(1))
	require.NoError(t, err)
	q, err = q.TallyVote(sdk.NewInt(9))
	require.NoError(t, err)

	require.Equal(t, int64(2), q.VoterCount)
	require.Equal(t, "10", q.VotingPool.String())
	require.Equal(t, "4.000000000000000000", q.RootSum.String())
	require.Equal(t, "6.000000000000000000", q.MatchPool().String())
	require.Equal(t, "5", q.VoterReward().String())
	require.Equal(t, "3.000000000000000000", q.MatchReward().String())
}
