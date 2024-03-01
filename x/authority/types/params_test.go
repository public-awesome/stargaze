package types_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v13/x/authority/types"
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
			"fail: no addrs",
			types.Params{
				Authorizations: []*types.Authorization{
					{
						Addresses:  []string{},
						MsgTypeUrl: "/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
					},
				},
			},
			true,
		},
		{
			"fail: invalid addr",
			types.Params{
				Authorizations: []*types.Authorization{
					{
						Addresses: []string{
							"👻",
						},
						MsgTypeUrl: "/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
					},
				},
			},
			true,
		},
		{
			"fail: valid addr but empty msg type url",
			types.Params{
				Authorizations: []*types.Authorization{
					{
						Addresses: []string{
							"cosmos1c4k24jzduc365kywrsvf5ujz4ya6mwymy8vq4q",
						},
						MsgTypeUrl: "",
					},
				},
			},
			true,
		},
		{
			"ok: all valid input",
			types.Params{
				Authorizations: []*types.Authorization{
					{
						Addresses: []string{
							"cosmos1c4k24jzduc365kywrsvf5ujz4ya6mwymy8vq4q",
						},
						MsgTypeUrl: "/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
					},
				},
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
