package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/stretchr/testify/assert"
)

func TestMsgJsonSignBytes(t *testing.T) {
	goodAddress := sdk.AccAddress(make([]byte, 20)).String()
	specs := map[string]struct {
		src legacytx.LegacyMsg
		exp string
	}{
		"MsgInitialClaim": {
			src: &MsgInitialClaim{Sender: goodAddress},
			exp: `
{
	"type":"claim/InitialClaim",
	"value": {"sender": "cosmos1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqnrql8a"}
}`,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			bz := spec.src.GetSignBytes()
			assert.JSONEq(t, spec.exp, string(bz), "raw: %s", string(bz))
		})
	}
}
