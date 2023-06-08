package types_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v11/x/globalfee/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "invalid params state",
			genState: &types.GenesisState{
				Params: types.Params{
					PrivilegedAddresses: []string{"👻"},
				},
				CodeAuthorizations:     types.DefaultGenesis().CodeAuthorizations,
				ContractAuthorizations: types.DefaultGenesis().ContractAuthorizations,
			},
			valid: false,
		},
		{
			desc: "invalid code authorization state",
			genState: &types.GenesisState{
				Params: types.DefaultGenesis().Params,
				CodeAuthorizations: []types.CodeAuthorization{
					{
						CodeID:  2,
						Methods: []string{},
					},
				},
				ContractAuthorizations: types.DefaultGenesis().ContractAuthorizations,
			},
			valid: false,
		},
		{
			desc: "invalid contract authorization state",
			genState: &types.GenesisState{
				Params:             types.DefaultGenesis().Params,
				CodeAuthorizations: types.DefaultGenesis().CodeAuthorizations,
				ContractAuthorizations: []types.ContractAuthorization{
					{
						ContractAddress: "👻",
						Methods:         []string{"tests"},
					},
				},
			},
			valid: false,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.Params{
					PrivilegedAddresses: []string{"cosmos1hfml4tzwlc3mvynsg6vtgywyx00wfkhrtpkx6t"},
				},
			},
			valid: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
