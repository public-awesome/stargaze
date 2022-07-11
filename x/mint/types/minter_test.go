package types

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCurrentYear(t *testing.T) {
	genesisTime := time.Now()
	actualYear := currentYear(time.Now().AddDate(0, 1, 0), genesisTime)
	require.Equal(t, uint64(0), actualYear)
}

func TestCurrentYear1(t *testing.T) {
	genesisTime := time.Now()
	actualYear := currentYear(time.Now().AddDate(2, 1, 0), genesisTime)
	require.Equal(t, uint64(2), actualYear)
}

func TestCurrentYear2(t *testing.T) {
	genesisTime := time.Now()
	actualYear := currentYear(time.Now().AddDate(2, 0, 2), genesisTime)
	require.Equal(t, uint64(2), actualYear)
}

func TestCurrentYear3(t *testing.T) {
	genesisTime := time.Now()
	actualYear := currentYear(time.Now().AddDate(1, 1, 0), genesisTime)
	require.Equal(t, uint64(1), actualYear)
}

func TestBlockProvision(t *testing.T) {
	minter := InitialMinter()
	params := DefaultParams()

	secondsPerYear := int64(60 * 60 * 8766)

	tests := []struct {
		annualProvisions int64
		expProvisions    int64
	}{
		{secondsPerYear / 5, 1},
		{secondsPerYear/5 + 1, 1},
		{(secondsPerYear / 5) * 2, 2},
		{(secondsPerYear / 5) / 2, 0},
	}
	for i, tc := range tests {
		minter.AnnualProvisions = sdk.NewDec(tc.annualProvisions)
		provisions := minter.BlockProvision(params)

		expProvisions := sdk.NewCoin(params.MintDenom,
			sdk.NewInt(tc.expProvisions))

		require.True(t, expProvisions.IsEqual(provisions),
			"test: %v\n\tExp: %v\n\tGot: %v\n",
			i, tc.expProvisions, provisions)
	}
}

// Benchmarking :)
// previously using sdk.Int operations:
// BenchmarkBlockProvision-4 5000000 220 ns/op
//
// using sdk.Dec operations: (current implementation)
// BenchmarkBlockProvision-4 3000000 429 ns/op
func BenchmarkBlockProvision(b *testing.B) {
	minter := InitialMinter()
	params := DefaultParams()

	s1 := rand.NewSource(100)
	r1 := rand.New(s1)
	minter.AnnualProvisions = sdk.NewDec(r1.Int63n(1000000))

	// run the BlockProvision function b.N times
	for n := 0; n < b.N; n++ {
		minter.BlockProvision(params)
	}
}

// Next annual provisions benchmarking
// BenchmarkNextAnnualProvisions-4 5000000 251 ns/op
func BenchmarkNextAnnualProvisions(b *testing.B) {
	minter := InitialMinter()
	params := DefaultParams()

	// run the NextAnnualProvisions function b.N times
	for n := 0; n < b.N; n++ {
		minter.NextAnnualProvisions(time.Time{}, params)
	}
}
