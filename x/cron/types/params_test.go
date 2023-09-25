package types_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v12/x/globalfee/types"
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
