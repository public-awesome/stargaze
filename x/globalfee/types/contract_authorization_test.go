package types_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v11/x/globalfee/types"
	"github.com/stretchr/testify/require"
)

func TestContractAuthorizationValidate(t *testing.T) {
	testCases := []struct {
		testCase    string
		ca          types.ContractAuthorization
		expectError bool
	}{
		{
			"fail: invalid contract address",
			types.ContractAuthorization{
				ContractAddress: "👻",
				Methods:         []string{"mint"},
			},
			true,
		},
		{
			"fail: empty methods",
			types.ContractAuthorization{
				ContractAddress: "cosmos19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseueszehj9h",
				Methods:         []string{},
			},
			true,
		},
		{
			"fail: invalid method names",
			types.ContractAuthorization{
				ContractAddress: "cosmos19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseueszehj9h",
				Methods:         []string{"^&()"},
			},
			true,
		},
		{
			"ok: valid name",
			types.ContractAuthorization{
				ContractAddress: "cosmos19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseueszehj9h",
				Methods:         []string{"mint"},
			},
			false,
		},
		{
			"ok: wildcard",
			types.ContractAuthorization{
				ContractAddress: "cosmos19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseueszehj9h",
				Methods:         []string{"*"},
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
