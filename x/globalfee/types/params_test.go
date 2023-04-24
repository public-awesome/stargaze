package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
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
			types.Params{PrivilegedAddress: []string{"👻"}},
			true,
		},
		{
			"ok: valid addr",
			types.Params{PrivilegedAddress: []string{"cosmos1c4k24jzduc365kywrsvf5ujz4ya6mwymy8vq4q"}},
			false,
		},
		{
			"fail: zero fees",
			types.Params{
				MinimumGasPrices: sdk.NewDecCoinsFromCoins(sdk.NewCoin("stars", sdk.ZeroInt())),
			},
			true,
		},
		{
			"fail: duplicate denom fees",
			types.Params{
				MinimumGasPrices: sdk.DecCoins{sdk.NewDecCoin("stars", sdk.OneInt()), sdk.NewDecCoin("stars", sdk.OneInt())},
			},
			true,
		},
		{
			"fail: unordered by denom",
			types.Params{
				MinimumGasPrices: sdk.DecCoins{sdk.NewDecCoin("stars", sdk.OneInt()), sdk.NewDecCoin("atom", sdk.OneInt())},
			},
			true,
		},
		{
			"ok: valid min gas fees",
			types.Params{
				MinimumGasPrices: sdk.DecCoins{sdk.NewDecCoin("atom", sdk.OneInt()), sdk.NewDecCoin("stars", sdk.OneInt())},
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
