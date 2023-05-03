package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/public-awesome/stargaze/v10/x/globalfee/keeper"
	"github.com/public-awesome/stargaze/v10/x/globalfee/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGovHandler(t *testing.T) {
	var (
		capturedContractAddrs []string
		capturedCodeIds       []uint64
	)
	notHandler := func(ctx sdk.Context, content govtypes.Content) error {
		return sdkerrors.ErrUnknownRequest
	}

	specs := map[string]struct {
		wasmHandler           govtypes.Handler
		setupGovKeeper        func(*MockGovKeeper)
		srcProposal           govtypes.Content
		expErr                *sdkerrors.Error
		expCapturedAddrs      []string
		expCapturedCodeIds    []uint64
		expCapturedGovContent []govtypes.Content
	}{
		"set code auth proposal": {
			wasmHandler: notHandler,
			setupGovKeeper: func(m *MockGovKeeper) {
				m.SetCodeAuthorizationFn = func(ctx sdk.Context, ca types.CodeAuthorization) error {
					capturedCodeIds = append(capturedCodeIds, ca.GetCodeId())
					return nil
				}
			},
			srcProposal: types.SetCodeAuthorizationProposalFixture(func(proposal *types.SetCodeAuthorizationProposal) {
				proposal.CodeAuthorization = &types.CodeAuthorization{
					CodeId:  1,
					Methods: []string{"*"},
				}
			}),
			expCapturedCodeIds: []uint64{1},
		},
		"invalid set code auth proposal rejected": {
			wasmHandler: notHandler,
			srcProposal: &types.SetCodeAuthorizationProposal{},
			expErr:      govtypes.ErrInvalidProposalContent,
		},
		"remove code auth proposal": {
			wasmHandler: notHandler,
			setupGovKeeper: func(m *MockGovKeeper) {
				m.DeleteCodeAuthorizationFn = func(ctx sdk.Context, codeId uint64) {
					// removing the codeIds
					for i, v := range capturedCodeIds {
						if v == codeId {
							capturedCodeIds = append(capturedCodeIds[:i], capturedCodeIds[i+1:]...)
							return
						}
					}
				}
			},
			srcProposal: types.RemoveCodeAuthorizationProposalFixture(func(proposal *types.RemoveCodeAuthorizationProposal) {
				proposal.CodeId = 1
			}),
			expCapturedCodeIds: nil,
		},
		"invalid remove code auth proposal rejected": {
			wasmHandler: notHandler,
			srcProposal: &types.RemoveCodeAuthorizationProposal{},
			expErr:      govtypes.ErrInvalidProposalContent,
		},
		"set contract auth proposal": {
			wasmHandler: notHandler,
			setupGovKeeper: func(m *MockGovKeeper) {
				m.SetContractAuthorizationFn = func(ctx sdk.Context, ca types.ContractAuthorization) error {
					capturedContractAddrs = append(capturedContractAddrs, ca.GetContractAddress())
					return nil
				}
			},
			srcProposal: types.SetContractAuthorizationProposalFixture(func(proposal *types.SetContractAuthorizationProposal) {
				proposal.ContractAuthorization = &types.ContractAuthorization{
					ContractAddress: "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du",
					Methods:         []string{"*"},
				}
			}),
			expCapturedAddrs: []string{"cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du"},
		},
		"invalid set contract auth proposal rejected": {
			wasmHandler: notHandler,
			srcProposal: &types.SetContractAuthorizationProposal{},
			expErr:      govtypes.ErrInvalidProposalContent,
		},
		"remove contract auth proposal": {
			wasmHandler: notHandler,
			setupGovKeeper: func(m *MockGovKeeper) {
				m.DeleteContractAuthorizationFn = func(ctx sdk.Context, contractAddr sdk.AccAddress) {
					// removing the contractAddresses
					for i, v := range capturedContractAddrs {
						if v == contractAddr.String() {
							capturedContractAddrs = append(capturedContractAddrs[:i], capturedContractAddrs[i+1:]...)
							return
						}
					}
				}
			},
			srcProposal: types.RemoveContractAuthorizationProposalFixture(func(proposal *types.RemoveContractAuthorizationProposal) {
				proposal.ContractAddress = "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du"
			}),
			expCapturedAddrs: nil,
		},
		"invalid remove contract auth proposal rejected": {
			wasmHandler: notHandler,
			srcProposal: &types.RemoveContractAuthorizationProposal{},
			expErr:      govtypes.ErrInvalidProposalContent,
		},
		"nil content": {
			wasmHandler: notHandler,
			expErr:      sdkerrors.ErrUnknownRequest,
		},
	}
	var ctx sdk.Context
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			capturedContractAddrs = nil
			var mock MockGovKeeper
			if spec.setupGovKeeper != nil {
				spec.setupGovKeeper(&mock)
			}
			// when
			router := &CapturingGovRouter{}
			h := keeper.NewProposalHandlerX(mock)
			gotErr := h(ctx, spec.srcProposal)
			// then
			require.True(t, spec.expErr.Is(gotErr), "exp %v but got #+v", spec.expErr, gotErr)
			if spec.expCapturedCodeIds != nil {
				assert.Equal(t, spec.expCapturedCodeIds, capturedCodeIds)
			} else {
				assert.Equal(t, spec.expCapturedAddrs, capturedContractAddrs)
			}
			assert.Equal(t, spec.expCapturedGovContent, router.captured)
		})
	}
}

type MockGovKeeper struct {
	SetCodeAuthorizationFn        func(ctx sdk.Context, ca types.CodeAuthorization) error
	DeleteCodeAuthorizationFn     func(ctx sdk.Context, codeId uint64)
	SetContractAuthorizationFn    func(ctx sdk.Context, ca types.ContractAuthorization) error
	DeleteContractAuthorizationFn func(ctx sdk.Context, contractAddr sdk.AccAddress)
}

func (m MockGovKeeper) SetCodeAuthorization(ctx sdk.Context, ca types.CodeAuthorization) error {
	if m.SetCodeAuthorizationFn == nil {
		panic("not expected to be called")
	}
	return m.SetCodeAuthorizationFn(ctx, ca)
}

func (m MockGovKeeper) DeleteCodeAuthorization(ctx sdk.Context, codeId uint64) {
	if m.DeleteCodeAuthorizationFn == nil {
		panic("not expected to be called")
	}
	m.DeleteCodeAuthorizationFn(ctx, codeId)
	return
}

func (m MockGovKeeper) SetContractAuthorization(ctx sdk.Context, ca types.ContractAuthorization) error {
	if m.SetContractAuthorizationFn == nil {
		panic("not expected to be called")
	}
	return m.SetContractAuthorizationFn(ctx, ca)
}

func (m MockGovKeeper) DeleteContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress) {
	if m.DeleteContractAuthorizationFn == nil {
		panic("not expected to be called")
	}
	m.DeleteContractAuthorizationFn(ctx, contractAddr)
	return
}

type CapturingGovRouter struct {
	govtypes.Router
	captured []govtypes.Content
}

func (m CapturingGovRouter) HasRoute(r string) bool {
	return true
}

func (m *CapturingGovRouter) GetRoute(path string) (h govtypes.Handler) {
	return func(ctx sdk.Context, content govtypes.Content) error {
		m.captured = append(m.captured, content)
		return nil
	}
}
