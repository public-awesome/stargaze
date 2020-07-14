package curating

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QVFData holds vars for quadratic voting+funding calculations
type QVFData struct {
	VoterCount int64
	VotingPool sdk.Int
	RootSum    sdk.Dec
}

func NewQVFData() QVFData {
	return QVFData{
		VotingPool: sdk.ZeroInt(),
		RootSum:    sdk.ZeroDec(),
	}
}

// TallyVote tallies a vote
func (q QVFData) TallyVote(amount sdk.Int) (QVFData, error) {
	q.VoterCount++
	q.VotingPool = q.VotingPool.Add(amount)

	sqrt, err := amount.ToDec().ApproxSqrt()
	if err != nil {
		return q, err
	}
	q.RootSum = q.RootSum.Add(sqrt)

	return q, nil
}

// MatchPool calculates and returns the quadratic match pool
func (q QVFData) MatchPool() sdk.Dec {
	return q.RootSum.Power(2).Sub(q.VotingPool.ToDec())
}

// VoterReward returns the distribution for a voter
func (q QVFData) VoterReward() sdk.Int {
	return q.VotingPool.QuoRaw(q.VoterCount)
}

// MatchReward returns the funding match for a voter
func (q QVFData) MatchReward() sdk.Dec {
	return q.MatchPool().QuoInt64(q.VoterCount)
}
