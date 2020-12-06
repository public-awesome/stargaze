package curating

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/keeper"
)

// QVFData holds vars for quadratic voting+funding calculations
type QVFData struct {
	ctx        sdk.Context
	keeper     keeper.Keeper
	VoterCount int64
	VotingPool sdk.Int
	RootSum    sdk.Dec
}

// NewQVFData returns an instance of QVFData
func NewQVFData(ctx sdk.Context, k keeper.Keeper) QVFData {
	return QVFData{
		ctx:        ctx,
		keeper:     k,
		VotingPool: sdk.ZeroInt(),
		RootSum:    sdk.ZeroDec(),
	}
}

// TallyVote tallies a vote
func (q QVFData) TallyVote(amount sdk.Int) (QVFData, error) {
	q.VoterCount++
	q.VotingPool = q.VotingPool.Add(amount)

	sqrt, err := amount.
		QuoRaw(1_000_000). // divide by 10^6, the default denom unit
		ToDec().           // convert to decimal
		ApproxSqrt()       // deterministic square root
	if err != nil {
		return q, err
	}
	q.RootSum = q.RootSum.Add(sqrt)

	return q, nil
}

// MatchPool calculates and returns the quadratic match pool
func (q QVFData) MatchPool() sdk.Dec {
	idealPoolSize := q.RootSum.
		Power(2).                 // increase quadratically
		MulInt64(1_000_000).      // multiply by 10^6, the default denom unit
		Sub(q.VotingPool.ToDec()) // subtract the voting pool

	rewardPool := q.keeper.GetRewardPoolBalance(q.ctx).Amount.ToDec()
	maxPoolPercent := q.keeper.GetParams(q.ctx).RewardPoolCurationMaxAlloc
	maxPoolSize := rewardPool.MulTruncate(maxPoolPercent)

	if idealPoolSize.GT(maxPoolSize) {
		return maxPoolSize
	}

	return idealPoolSize
}

// VoterReward returns the distribution for a voter
func (q QVFData) VoterReward() sdk.Int {
	if q.VoterCount == 0 {
		return sdk.ZeroInt()
	}
	return q.VotingPool.QuoRaw(q.VoterCount)
}

// MatchReward returns the funding match for a voter
func (q QVFData) MatchReward() sdk.Dec {
	if q.VoterCount == 0 {
		return sdk.ZeroDec()
	}
	return q.MatchPool().QuoInt64(q.VoterCount)
}

// MatchPoolPerVote calculates the portion of the match pool per vote
func (q QVFData) MatchPoolPerVote() sdk.Dec {
	if q.VoterCount == 0 {
		return sdk.ZeroDec()
	}
	fmt.Println(q.RootSum)
	return q.MatchPool().Quo(q.RootSum)
}
