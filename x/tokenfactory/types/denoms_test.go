package types_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v18/x/tokenfactory/types"
	"github.com/stretchr/testify/require"
)

func TestDeconstructDenom(t *testing.T) {
	for _, tc := range []struct {
		desc             string
		denom            string
		expectedSubdenom string
		err              error
	}{
		{
			desc:  "empty is invalid",
			denom: "",
			err:   types.ErrInvalidDenom,
		},
		{
			desc:             "normal",
			denom:            "factory/cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f/testy",
			expectedSubdenom: "testy",
		},
		{
			desc:             "multiple slashes in subdenom",
			denom:            "factory/cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f/testy/1",
			expectedSubdenom: "testy/1",
		},
		{
			desc:             "no subdenom",
			denom:            "factory/cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f/",
			expectedSubdenom: "",
		},
		{
			desc:  "incorrect prefix",
			denom: "ibc/cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f/testy",
			err:   types.ErrInvalidDenom,
		},
		{
			desc:             "subdenom of only slashes",
			denom:            "factory/cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f/////",
			expectedSubdenom: "////",
		},
		{
			desc:  "too long name",
			denom: "factory/cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f/adsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsf",
			err:   types.ErrInvalidDenom,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			expectedCreator := "cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f"
			creator, subdenom, err := types.DeconstructDenom(tc.denom)
			if tc.err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, expectedCreator, creator)
				require.Equal(t, tc.expectedSubdenom, subdenom)
			}
		})
	}
}

func TestGetTokenDenom(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		creator  string
		subdenom string
		valid    bool
	}{
		{
			desc:     "normal",
			creator:  "cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f",
			subdenom: "testy",
			valid:    true,
		},
		{
			desc:     "multiple slashes in subdenom",
			creator:  "cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f",
			subdenom: "testy/1",
			valid:    true,
		},
		{
			desc:     "no subdenom",
			creator:  "cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f",
			subdenom: "",
			valid:    true,
		},
		{
			desc:     "subdenom of only slashes",
			creator:  "cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f",
			subdenom: "/////",
			valid:    true,
		},
		{
			desc:     "too long name",
			creator:  "cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f",
			subdenom: "adsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsf",
			valid:    false,
		},
		{
			desc:     "subdenom is exactly max length",
			creator:  "cosmos1rf3renzcj8m2pav74758lj7wm8z98yky20x64f",
			subdenom: "testyfsadfsdfeadfsafwefsefsefsdfsdafasefsf",
			valid:    true,
		},
		{
			desc:     "creator is exactly max length",
			creator:  "cosmos1t7egva48prqmzl59x5ngv4zx0dtrwewcdqdjr8jhgjhgkhjklhkjhkhgjhgjgjgheugt",
			subdenom: "testy",
			valid:    true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := types.GetTokenDenom(tc.creator, tc.subdenom)
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
