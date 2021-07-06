package curating_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/simapp"
	"github.com/public-awesome/stargaze/x/curating"
	curatingtypes "github.com/public-awesome/stargaze/x/curating/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// 1 vote  = 1 vote credit
// 2 votes = 4 vote credits
// 3 votes = 9 vote credits
// 4 votes = 16 vote credits
// 5 votes = 25 vote credits
func TestQVF(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, fakedenom)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// add funds to reward pool
	funds := sdk.NewInt64Coin(fakedenom, 10_000_000)
	err := app.BankKeeper.MintCoins(ctx, curatingtypes.RewardPoolName, sdk.NewCoins(funds))
	require.NoError(t, err)

	q := curating.NewQVFData(ctx, app.CuratingKeeper)
	q, err = q.TallyVote(sdk.NewInt(1_000_000))
	require.NoError(t, err)
	q, err = q.TallyVote(sdk.NewInt(9_000_000))
	require.NoError(t, err)

	// for reference...
	require.Equal(t, int64(2), q.VoterCount)
	require.Equal(t, "10000000", q.VotingPool.String())
	require.Equal(t, "4.000000000000000000", q.RootSum.String())
	require.Equal(t, "6000000.000000000000000000", q.MatchPool().String())
	require.Equal(t, "1500000.000000000000000000", q.MatchPoolPerVote().String())
	require.Equal(t, "5000000", q.VoterReward().String())
	require.Equal(t, "3000000.000000000000000000", q.MatchReward().String())
}

func TestQVFZeroVotes(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, fakedenom)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// add funds to reward pool
	funds := sdk.NewInt64Coin(fakedenom, 10_000_000)
	err := app.BankKeeper.MintCoins(ctx, curatingtypes.RewardPoolName, sdk.NewCoins(funds))
	require.NoError(t, err)

	q := curating.NewQVFData(ctx, app.CuratingKeeper)

	// for reference...
	require.Equal(t, int64(0), q.VoterCount)
	require.Equal(t, "0", q.VotingPool.String())
	require.Equal(t, "0.000000000000000000", q.RootSum.String())
	require.Equal(t, "0.000000000000000000", q.MatchPool().String())
	require.Equal(t, "0.000000000000000000", q.MatchPoolPerVote().String())
	require.Equal(t, "0", q.VoterReward().String())
	require.Equal(t, "0.000000000000000000", q.MatchReward().String())
}

func TestQVFOneVote(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// add funds to reward pool
	funds := sdk.NewInt64Coin("ustarx", 10_000_000)
	err := app.BankKeeper.MintCoins(ctx, curatingtypes.RewardPoolName, sdk.NewCoins(funds))
	require.NoError(t, err)

	q := curating.NewQVFData(ctx, app.CuratingKeeper)
	q, err = q.TallyVote(sdk.NewInt(1_000_000))
	require.NoError(t, err)

	require.Equal(t, int64(1), q.VoterCount)
	require.Equal(t, "1000000", q.VotingPool.String())
	require.Equal(t, "1.000000000000000000", q.RootSum.String())
	require.Equal(t, "0.000000000000000000", q.MatchPool().String())
	require.Equal(t, "0.000000000000000000", q.MatchPoolPerVote().String())
	require.Equal(t, "1000000", q.VoterReward().String())
	require.Equal(t, "0.000000000000000000", q.MatchReward().String())
}
