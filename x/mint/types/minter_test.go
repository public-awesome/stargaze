package types

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestNextInflation(t *testing.T) {
	minter := DefaultInitialMinter()
	params := DefaultParams()

	tests := []struct {
		blockTime    time.Time
		expInflation sdk.Dec
	}{
		// inflation start time = in 1 year from now
		// { blockTime, inflation }
		// before inflation start time
		{time.Now(), sdk.ZeroDec()},
		// year 1 after start, inflation 100%
		{time.Now().AddDate(1, 0, 0), sdk.NewDecWithPrec(100, 2)},
		// year 2 after start, inflation 67%
		{time.Now().AddDate(2, 0, 0), sdk.NewDecWithPrec(67, 2)},
		// year 3 after start, inflation 44%
		{time.Now().AddDate(3, 0, 0), sdk.NewDecWithPrec(4489, 4)},
		// year 4 after start, inflation 30%
		{time.Now().AddDate(4, 0, 0), sdk.NewDecWithPrec(300763, 6)},
		// year 5 after start, inflation 20%
		{time.Now().AddDate(5, 0, 0), sdk.NewDecWithPrec(20151121, 8)},
	}
	for i, tc := range tests {
		inflation := minter.NextInflationRate(tc.blockTime, params)

		require.True(t, inflation.Equal(tc.expInflation),
			"Test Index: %v\nInflation:  %v\nExpected: %v\n", i, inflation, tc.expInflation)
	}
}

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
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
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
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	params := DefaultParams()

	s1 := rand.NewSource(100)
	r1 := rand.New(s1)
	minter.AnnualProvisions = sdk.NewDec(r1.Int63n(1000000))

	// run the BlockProvision function b.N times
	for n := 0; n < b.N; n++ {
		minter.BlockProvision(params)
	}
}

// Next inflation benchmarking
// BenchmarkNextInflation-4 1000000 1828 ns/op
func BenchmarkNextInflation(b *testing.B) {
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	params := DefaultParams()

	// run the NextInflationRate function b.N times
	for n := 0; n < b.N; n++ {
		minter.NextInflationRate(time.Now(), params)
	}
}

// Next annual provisions benchmarking
// BenchmarkNextAnnualProvisions-4 5000000 251 ns/op
func BenchmarkNextAnnualProvisions(b *testing.B) {
	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
	params := DefaultParams()
	totalSupply := sdk.NewInt(100000000000000)

	// run the NextAnnualProvisions function b.N times
	for n := 0; n < b.N; n++ {
		minter.NextAnnualProvisions(params, totalSupply)
	}

}
