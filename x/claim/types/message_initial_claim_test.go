package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stargaze/v6/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgInitialClaim_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgInitialClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgInitialClaim{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgInitialClaim{
				Sender: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
