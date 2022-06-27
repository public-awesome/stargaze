package types

import (
	"errors"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stargaze/v6/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateVestingAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateVestingAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVestingAccount{
				FromAddress: "invalid_address",
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("stake", 10)),
				StartTime:   time.Now().Unix(),
				EndTime:     time.Now().Unix() + 1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid to address",
			msg: MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   "invalid_address",
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("stake", 10)),
				StartTime:   time.Now().Unix(),
				EndTime:     time.Now().Unix() + 1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid start time",
			msg: MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("stake", 10)),
				StartTime:   0,
				EndTime:     time.Now().Unix() + 1,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid end time",
			msg: MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("stake", 10)),
				StartTime:   time.Now().Unix(),
				EndTime:     0,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "star time < end time",
			msg: MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("stake", 10)),
				StartTime:   time.Now().Unix(),
				EndTime:     time.Now().Unix(),
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid amount",
			msg: MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),

				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 1,
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "valid address",
			msg: MsgCreateVestingAccount{
				FromAddress: sample.AccAddress(),
				ToAddress:   sample.AccAddress(),
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("stake", 10)),
				StartTime:   time.Now().Unix(),
				EndTime:     time.Now().Unix() + 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {

				require.EqualError(t, errors.Unwrap(err), tt.err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}
