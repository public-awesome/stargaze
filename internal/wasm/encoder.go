package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Encoder describes behavior for Stargaze smart contract message encoding.
// The contract address must ALWAYS be set as the Msg signer.
type Encoder func(contract sdk.AccAddress, data json.RawMessage, version string) ([]sdk.Msg, error)

// MessageEncoders provides stargaze custom encoder for contracts
func MessageEncoders(registry *EncoderRegistry) *wasmkeeper.MessageEncoders {
	return &wasmkeeper.MessageEncoders{
		Custom: customEncoders(registry),
	}
}

type MessageEncodeRequest struct {
	Route   string          `json:"route"`
	MsgData json.RawMessage `json:"msg_data"`
	Version string          `json:"version"`
}

func customEncoders(registry *EncoderRegistry) wasmkeeper.CustomEncoder {
	return func(sender sdk.AccAddress, m json.RawMessage) ([]sdk.Msg, error) {
		encodeRequest := &MessageEncodeRequest{}
		err := json.Unmarshal(m, encodeRequest)
		if err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
		}
		encode, exists := registry.encoders[encodeRequest.Route]
		if !exists {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "encoder not found for route: %s", encodeRequest.Route)
		}

		msgs, err := encode(sender, encodeRequest.MsgData, encodeRequest.Version)
		if err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		for _, msg := range msgs {
			m, ok := msg.(sdk.HasValidateBasic)
			if !ok {
				continue
			}
			if err := m.ValidateBasic(); err != nil {
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
			}
		}
		return msgs, nil
	}
}
