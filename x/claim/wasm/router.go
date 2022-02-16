package wasm

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Encoder(sender sdk.AccAddress, msg json.RawMessage) ([]sdk.Msg, error) {
	return nil, fmt.Errorf("not implemented")
}
