package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/public-awesome/stargaze/v17/testutil/keeper"
	"github.com/public-awesome/stargaze/v17/testutil/sample"
	"github.com/public-awesome/stargaze/v17/x/pauser/keeper"
	"github.com/public-awesome/stargaze/v17/x/pauser/types"
	"github.com/stretchr/testify/require"
)

func TestPauseContract(t *testing.T) {
	testCases := []struct {
		testCase    string
		prepare     func(ctx sdk.Context, keeper keeper.Keeper) *types.MsgPauseContract
		expectError bool
	}{
		{
			"invalid sender address",
			func(_ sdk.Context, _ keeper.Keeper) *types.MsgPauseContract {
				return &types.MsgPauseContract{
					Sender:          "invalid",
					ContractAddress: keepertest.TestContract1,
				}
			},
			true,
		},
		{
			"unauthorized sender",
			func(_ sdk.Context, _ keeper.Keeper) *types.MsgPauseContract {
				sender := sample.AccAddress()
				return &types.MsgPauseContract{
					Sender:          sender.String(),
					ContractAddress: keepertest.TestContract1,
				}
			},
			true,
		},
		{
			"non-existent contract",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgPauseContract {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)
				return &types.MsgPauseContract{
					Sender:          sender.String(),
					ContractAddress: sample.AccAddress().String(),
				}
			},
			true,
		},
		{
			"already paused contract",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgPauseContract {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)

				// Pre-pause the contract
				err = k.SetPausedContract(ctx, types.PausedContract{
					ContractAddress: keepertest.TestContract1,
					PausedBy:        sender.String(),
				})
				require.NoError(t, err)

				return &types.MsgPauseContract{
					Sender:          sender.String(),
					ContractAddress: keepertest.TestContract1,
				}
			},
			true,
		},
		{
			"valid pause",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgPauseContract {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)
				return &types.MsgPauseContract{
					Sender:          sender.String(),
					ContractAddress: keepertest.TestContract1,
				}
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testCase, func(t *testing.T) {
			k, c := keepertest.PauserKeeper(t)
			msgSrvr := keeper.NewMsgServerImpl(k)

			msg := tc.prepare(c, k)
			_, err := msgSrvr.PauseContract(c, msg)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				contractAddr := sdk.MustAccAddressFromBech32(msg.ContractAddress)
				require.True(t, k.IsContractPaused(c, contractAddr))
			}
		})
	}
}

func TestUnpauseContract(t *testing.T) {
	testCases := []struct {
		testCase    string
		prepare     func(ctx sdk.Context, keeper keeper.Keeper) *types.MsgUnpauseContract
		expectError bool
	}{
		{
			"invalid sender address",
			func(_ sdk.Context, _ keeper.Keeper) *types.MsgUnpauseContract {
				return &types.MsgUnpauseContract{
					Sender:          "invalid",
					ContractAddress: keepertest.TestContract1,
				}
			},
			true,
		},
		{
			"unauthorized sender",
			func(_ sdk.Context, _ keeper.Keeper) *types.MsgUnpauseContract {
				sender := sample.AccAddress()
				return &types.MsgUnpauseContract{
					Sender:          sender.String(),
					ContractAddress: keepertest.TestContract1,
				}
			},
			true,
		},
		{
			"not paused contract",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgUnpauseContract {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)
				return &types.MsgUnpauseContract{
					Sender:          sender.String(),
					ContractAddress: keepertest.TestContract1,
				}
			},
			true,
		},
		{
			"valid unpause",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgUnpauseContract {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)

				// Pre-pause the contract
				err = k.SetPausedContract(ctx, types.PausedContract{
					ContractAddress: keepertest.TestContract1,
					PausedBy:        sender.String(),
				})
				require.NoError(t, err)

				return &types.MsgUnpauseContract{
					Sender:          sender.String(),
					ContractAddress: keepertest.TestContract1,
				}
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testCase, func(t *testing.T) {
			k, c := keepertest.PauserKeeper(t)
			msgSrvr := keeper.NewMsgServerImpl(k)

			msg := tc.prepare(c, k)
			_, err := msgSrvr.UnpauseContract(c, msg)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				contractAddr := sdk.MustAccAddressFromBech32(msg.ContractAddress)
				require.False(t, k.IsContractPaused(c, contractAddr))
			}
		})
	}
}

func TestPauseCodeID(t *testing.T) {
	testCases := []struct {
		testCase    string
		prepare     func(ctx sdk.Context, keeper keeper.Keeper) *types.MsgPauseCodeID
		expectError bool
	}{
		{
			"invalid sender address",
			func(_ sdk.Context, _ keeper.Keeper) *types.MsgPauseCodeID {
				return &types.MsgPauseCodeID{
					Sender: "invalid",
					CodeID: 1,
				}
			},
			true,
		},
		{
			"unauthorized sender",
			func(_ sdk.Context, _ keeper.Keeper) *types.MsgPauseCodeID {
				sender := sample.AccAddress()
				return &types.MsgPauseCodeID{
					Sender: sender.String(),
					CodeID: 1,
				}
			},
			true,
		},
		{
			"non-existent code ID",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgPauseCodeID {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)
				return &types.MsgPauseCodeID{
					Sender: sender.String(),
					CodeID: 999,
				}
			},
			true,
		},
		{
			"already paused code ID",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgPauseCodeID {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)

				err = k.SetPausedCodeID(ctx, types.PausedCodeID{
					CodeID:   1,
					PausedBy: sender.String(),
				})
				require.NoError(t, err)

				return &types.MsgPauseCodeID{
					Sender: sender.String(),
					CodeID: 1,
				}
			},
			true,
		},
		{
			"valid pause",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgPauseCodeID {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)
				return &types.MsgPauseCodeID{
					Sender: sender.String(),
					CodeID: 1,
				}
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testCase, func(t *testing.T) {
			k, c := keepertest.PauserKeeper(t)
			msgSrvr := keeper.NewMsgServerImpl(k)

			msg := tc.prepare(c, k)
			_, err := msgSrvr.PauseCodeID(c, msg)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.True(t, k.IsCodeIDPaused(c, msg.CodeID))
			}
		})
	}
}

func TestUnpauseCodeID(t *testing.T) {
	testCases := []struct {
		testCase    string
		prepare     func(ctx sdk.Context, keeper keeper.Keeper) *types.MsgUnpauseCodeID
		expectError bool
	}{
		{
			"invalid sender address",
			func(_ sdk.Context, _ keeper.Keeper) *types.MsgUnpauseCodeID {
				return &types.MsgUnpauseCodeID{
					Sender: "invalid",
					CodeID: 1,
				}
			},
			true,
		},
		{
			"unauthorized sender",
			func(_ sdk.Context, _ keeper.Keeper) *types.MsgUnpauseCodeID {
				sender := sample.AccAddress()
				return &types.MsgUnpauseCodeID{
					Sender: sender.String(),
					CodeID: 1,
				}
			},
			true,
		},
		{
			"not paused code ID",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgUnpauseCodeID {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)
				return &types.MsgUnpauseCodeID{
					Sender: sender.String(),
					CodeID: 1,
				}
			},
			true,
		},
		{
			"valid unpause",
			func(ctx sdk.Context, k keeper.Keeper) *types.MsgUnpauseCodeID {
				sender := sample.AccAddress()
				params := types.Params{PrivilegedAddresses: []string{sender.String()}}
				err := k.SetParams(ctx, params)
				require.NoError(t, err)

				err = k.SetPausedCodeID(ctx, types.PausedCodeID{
					CodeID:   1,
					PausedBy: sender.String(),
				})
				require.NoError(t, err)

				return &types.MsgUnpauseCodeID{
					Sender: sender.String(),
					CodeID: 1,
				}
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testCase, func(t *testing.T) {
			k, c := keepertest.PauserKeeper(t)
			msgSrvr := keeper.NewMsgServerImpl(k)

			msg := tc.prepare(c, k)
			_, err := msgSrvr.UnpauseCodeID(c, msg)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.False(t, k.IsCodeIDPaused(c, msg.CodeID))
			}
		})
	}
}

func TestUpdateParams(t *testing.T) {
	authority := "cosmos1a48wdtjn3egw7swhfkeshwdtjvs6hq9nlyrwut"

	testCases := []struct {
		testCase    string
		prepare     func() *types.MsgUpdateParams
		expectError bool
	}{
		{
			"non-authority sender rejected",
			func() *types.MsgUpdateParams {
				sender := sample.AccAddress()
				return &types.MsgUpdateParams{
					Sender: sender.String(),
					Params: types.Params{PrivilegedAddresses: []string{}},
				}
			},
			true,
		},
		{
			"valid params update",
			func() *types.MsgUpdateParams {
				privAddr := sample.AccAddress()
				return &types.MsgUpdateParams{
					Sender: authority,
					Params: types.Params{PrivilegedAddresses: []string{privAddr.String()}},
				}
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testCase, func(t *testing.T) {
			k, c := keepertest.PauserKeeper(t)
			msgSrvr := keeper.NewMsgServerImpl(k)

			msg := tc.prepare()
			_, err := msgSrvr.UpdateParams(c, msg)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				params, err := k.GetParams(c)
				require.NoError(t, err)
				require.Equal(t, msg.Params.PrivilegedAddresses, params.PrivilegedAddresses)
			}
		})
	}
}

func TestIsExecutionPaused(t *testing.T) {
	t.Run("direct contract pause", func(t *testing.T) {
		k, ctx := keepertest.PauserKeeper(t)
		contractAddr := sdk.MustAccAddressFromBech32(keepertest.TestContract1)

		require.False(t, k.IsExecutionPaused(ctx, contractAddr))

		err := k.SetPausedContract(ctx, types.PausedContract{
			ContractAddress: keepertest.TestContract1,
			PausedBy:        "someone",
		})
		require.NoError(t, err)

		require.True(t, k.IsExecutionPaused(ctx, contractAddr))
	})

	t.Run("code ID pause", func(t *testing.T) {
		k, ctx := keepertest.PauserKeeper(t)
		// TestContract1 has CodeID 1 per the mock
		contractAddr := sdk.MustAccAddressFromBech32(keepertest.TestContract1)

		require.False(t, k.IsExecutionPaused(ctx, contractAddr))

		err := k.SetPausedCodeID(ctx, types.PausedCodeID{
			CodeID:   1,
			PausedBy: "someone",
		})
		require.NoError(t, err)

		require.True(t, k.IsExecutionPaused(ctx, contractAddr))
	})

	t.Run("unpause restores execution", func(t *testing.T) {
		k, ctx := keepertest.PauserKeeper(t)
		contractAddr := sdk.MustAccAddressFromBech32(keepertest.TestContract1)

		// Pause and then unpause contract
		err := k.SetPausedContract(ctx, types.PausedContract{
			ContractAddress: keepertest.TestContract1,
			PausedBy:        "someone",
		})
		require.NoError(t, err)
		require.True(t, k.IsExecutionPaused(ctx, contractAddr))

		err = k.DeletePausedContract(ctx, contractAddr)
		require.NoError(t, err)
		require.False(t, k.IsExecutionPaused(ctx, contractAddr))

		// Pause and then unpause code ID
		err = k.SetPausedCodeID(ctx, types.PausedCodeID{
			CodeID:   1,
			PausedBy: "someone",
		})
		require.NoError(t, err)
		require.True(t, k.IsExecutionPaused(ctx, contractAddr))

		err = k.DeletePausedCodeID(ctx, 1)
		require.NoError(t, err)
		require.False(t, k.IsExecutionPaused(ctx, contractAddr))
	})
}
