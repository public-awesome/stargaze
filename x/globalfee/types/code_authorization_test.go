package types_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
	"github.com/stretchr/testify/require"
)

func TestCodeAuthorizationValidate(t *testing.T) {
	testCases := []struct {
		testCase    string
		ca          types.CodeAuthorization
		expectError bool
	}{
		{
			"fail: empty methods",
			types.CodeAuthorization{
				CodeId:  1,
				Methods: []string{},
			},
			true,
		},
		{
			"fail: invalid method names",
			types.CodeAuthorization{
				CodeId:  1,
				Methods: []string{"^&()"},
			},
			true,
		},
		{
			"ok: valid name",
			types.CodeAuthorization{
				CodeId:  1,
				Methods: []string{"mint"},
			},
			false,
		},
		{
			"ok: wildcard",
			types.CodeAuthorization{
				CodeId:  1,
				Methods: []string{"*"},
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testCase, func(t *testing.T) {
			err := tc.ca.Validate()
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}