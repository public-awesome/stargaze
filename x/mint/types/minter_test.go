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
		// year 1, inflation 100%
		{time.Now(), sdk.NewDecWithPrec(100, 2)},
		// year 2, inflation 67%
		{time.Now(), sdk.NewDecWithPrec(67, 2)},
	}
	for i, tc := range tests {
		inflation := minter.NextInflationRate(tc.blockTime, params)

		require.True(t, inflation.Equal(tc.expInflation),
			"Test Index: %v\nInflation:  %v\nExpected: %v\n", i, inflation, tc.expInflation)
	}
}

func TestYear(t *testing.T) {
	minter := DefaultInitialMinter()
	genesisTime := time.Now().AddDate(-2, 0, 0)
	actualYear := minter.CurrentYear(time.Now(), genesisTime)
	require.Equal(t, int64(2), actualYear)
}

func TestYearBoundry(t *testing.T) {
	minter := DefaultInitialMinter()
	genesisTime := time.Now().AddDate(-2, -11, 0)
	actualYear := minter.CurrentYear(time.Now(), genesisTime)
	require.Equal(t, int64(2), actualYear)
}

func TestYearBoundry1(t *testing.T) {
	minter := DefaultInitialMinter()
	genesisTime := time.Now().AddDate(-2, 0, 2)
	actualYear := minter.CurrentYear(time.Now(), genesisTime)
	require.Equal(t, int64(2), actualYear)
}

func TestYearBoundry2(t *testing.T) {
	minter := DefaultInitialMinter()
	genesisTime := time.Now().AddDate(-1, -1, 0)
	actualYear := minter.CurrentYear(time.Now(), genesisTime)
	require.Equal(t, int64(1), actualYear)
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

// // Next inflation benchmarking
// // BenchmarkNextInflation-4 1000000 1828 ns/op
// func BenchmarkNextInflation(b *testing.B) {
// 	minter := InitialMinter(sdk.NewDecWithPrec(1, 1))
// 	params := DefaultParams()
// 	bondedRatio := sdk.NewDecWithPrec(1, 1)

// 	// run the NextInflationRate function b.N times
// 	for n := 0; n < b.N; n++ {
// 		minter.NextInflationRate(params, bondedRatio)
// 	}

// }

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
