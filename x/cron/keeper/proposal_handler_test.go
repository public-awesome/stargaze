package keeper_test

import (
	"testing"

	"github.com/cometbft/cometbft/libs/rand"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	legacygovtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/public-awesome/stargaze/v12/x/cron/keeper"
	"github.com/public-awesome/stargaze/v12/x/cron/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGovHandler(t *testing.T) {
	var (
		myAddr                sdk.AccAddress = rand.Bytes(address.Len)
		capturedContractAddrs []sdk.AccAddress
	)
	notHandler := func(ctx sdk.Context, content legacygovtypes.Content) error {
		return sdkerrors.ErrUnknownRequest
	}

	specs := map[string]struct {
		wasmHandler           legacygovtypes.Handler
		setupGovKeeper        func(*MockGovKeeper)
		srcProposal           legacygovtypes.Content
		expErr                *sdkerrors.Error
		expCapturedAddrs      []sdk.AccAddress
		expCapturedGovContent []legacygovtypes.Content
	}{
		"promote proposal": {
			wasmHandler: notHandler,
			setupGovKeeper: func(m *MockGovKeeper) {
				m.SetPrivilegedFn = func(ctx sdk.Context, contractAddr sdk.AccAddress) error {
					capturedContractAddrs = append(capturedContractAddrs, contractAddr)
					return nil
				}
			},
			srcProposal: types.PromoteProposalFixture(func(proposal *types.PromoteToPrivilegedContractProposal) {
				proposal.Contract = myAddr.String()
			}),
			expCapturedAddrs: []sdk.AccAddress{myAddr},
		},
		"invalid promote proposal rejected": {
			wasmHandler: notHandler,
			srcProposal: &types.PromoteToPrivilegedContractProposal{},
			expErr:      govtypes.ErrInvalidProposalContent,
		},
		"demote proposal": {
			wasmHandler: notHandler,
			setupGovKeeper: func(m *MockGovKeeper) {
				m.UnsetPrivilegedFn = func(ctx sdk.Context, contractAddr sdk.AccAddress) error {
					capturedContractAddrs = append(capturedContractAddrs, contractAddr)
					return nil
				}
			},
			srcProposal: types.DemoteProposalFixture(func(proposal *types.DemotePrivilegedContractProposal) {
				proposal.Contract = myAddr.String()
			}),
			expCapturedAddrs: []sdk.AccAddress{myAddr},
		},
		"invalid demote proposal rejected": {
			wasmHandler: notHandler,
			srcProposal: &types.DemotePrivilegedContractProposal{},
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
			assert.Equal(t, spec.expCapturedAddrs, capturedContractAddrs)
			assert.Equal(t, spec.expCapturedGovContent, router.captured)
		})
	}
}

type MockGovKeeper struct {
	SetPrivilegedFn   func(ctx sdk.Context, contractAddr sdk.AccAddress) error
	UnsetPrivilegedFn func(ctx sdk.Context, contractAddr sdk.AccAddress) error
}

func (m MockGovKeeper) SetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	if m.SetPrivilegedFn == nil {
		panic("not expected to be called")
	}
	return m.SetPrivilegedFn(ctx, contractAddr)
}

func (m MockGovKeeper) UnsetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	if m.UnsetPrivilegedFn == nil {
		panic("not expected to be called")
	}
	return m.UnsetPrivilegedFn(ctx, contractAddr)
}

type CapturingGovRouter struct {
	legacygovtypes.Router
	captured []legacygovtypes.Content
}

func (m CapturingGovRouter) HasRoute(_ string) bool {
	return true
}

func (m *CapturingGovRouter) GetRoute(_ string) (h legacygovtypes.Handler) {
	return func(ctx sdk.Context, content legacygovtypes.Content) error {
		m.captured = append(m.captured, content)
		return nil
	}
}
