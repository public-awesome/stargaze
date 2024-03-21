package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v14/x/globalfee/types"
	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	testCases := []struct {
		testCase    string
		p           types.Params
		expectError bool
	}{
		{
			"ok: empty params",
			types.Params{},
			false,
		},
		{
			"fail: invalid addr",
			types.Params{PrivilegedAddresses: []string{"👻"}},
			true,
		},
		{
			"ok: valid addr",
			types.Params{PrivilegedAddresses: []string{"cosmos1c4k24jzduc365kywrsvf5ujz4ya6mwymy8vq4q"}},
			false,
		},
		{
			"ok: zero fees",
			types.Params{
				MinimumGasPrices: sdk.NewDecCoinsFromCoins(sdk.NewCoin("stars", sdkmath.ZeroInt())),
			},
			false,
		},
		{
			"fail: duplicate denom fees",
			types.Params{
				MinimumGasPrices: sdk.DecCoins{sdk.NewDecCoin("stars", sdkmath.OneInt()), sdk.NewDecCoin("stars", sdkmath.OneInt())},
			},
			true,
		},
		{
			"fail: unordered by denom",
			types.Params{
				MinimumGasPrices: sdk.DecCoins{sdk.NewDecCoin("stars", sdkmath.OneInt()), sdk.NewDecCoin("atom", sdkmath.OneInt())},
			},
			true,
		},
		{
			"ok: valid min gas fees",
			types.Params{
				MinimumGasPrices: sdk.DecCoins{sdk.NewDecCoin("atom", sdkmath.OneInt()), sdk.NewDecCoin("stars", sdkmath.OneInt())},
			},
			false,
		},
		{
			"ok: default params",
			types.DefaultParams(),
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testCase, func(t *testing.T) {
			err := tc.p.Validate()
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
