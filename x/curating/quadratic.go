package curating

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/keeper"
)

// QVFData holds vars for quadratic voting+funding calculations
type QVFData struct {
	ctx        sdk.Context
	keeper     keeper.Keeper
	voterCount int64
	VotingPool sdk.Int
	rootSum    sdk.Dec
}

// NewQVFData returns an instance of QVFData
func NewQVFData(ctx sdk.Context, k keeper.Keeper) QVFData {
	return QVFData{
		ctx:        ctx,
		keeper:     k,
		VotingPool: sdk.ZeroInt(),
		rootSum:    sdk.ZeroDec(),
	}
}

// TallyVote tallies a vote
func (q QVFData) TallyVote(amount sdk.Int) (QVFData, error) {
	q.voterCount++
	q.VotingPool = q.VotingPool.Add(amount)

	sqrt, err := amount.
		QuoRaw(1_000_000). // divide by 10^6, the default denom unit
		ToDec().           // convert to decimal
		ApproxSqrt()       // deterministic square root
	if err != nil {
		return q, err
	}
	q.rootSum = q.rootSum.Add(sqrt)

	return q, nil
}

// MatchPool calculates and returns the quadratic match pool
func (q QVFData) MatchPool() sdk.Dec {
	idealPoolSize := q.rootSum.
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
	if q.voterCount == 0 {
		return sdk.ZeroInt()
	}
	return q.VotingPool.QuoRaw(q.voterCount)
}

// MatchReward returns the funding match for a voter
func (q QVFData) MatchReward() sdk.Dec {
	if q.voterCount == 0 {
		return sdk.ZeroDec()
	}
	return q.MatchPool().QuoInt64(q.voterCount)
}
