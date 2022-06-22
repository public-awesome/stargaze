package wasm

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sgwasm "github.com/public-awesome/stargaze/v6/internal/wasm"
	claimtypes "github.com/public-awesome/stargaze/v6/x/claim/types"
)

var _ sgwasm.Encoder = Encoder

type ClaimAction string

const (
	ClaimActionMintNFT = "mint_nft"
	ClaimActionBidNFT  = "bid_nft"
)

type ClaimFor struct {
	Address string      `json:"address"`
	Action  ClaimAction `json:"action"`
}

func (a ClaimAction) ToAction() (claimtypes.Action, error) {
	if a == ClaimActionMintNFT {
		return claimtypes.ActionMintNFT, nil
	}

	// rebranding the action
	if a == ClaimActionBidNFT {
		return claimtypes.ActionBidNFT, nil
	}

	return 0, fmt.Errorf("invalid action")
}

type ClaimMsg struct {
	ClaimFor *ClaimFor `json:"claim_for,omitempty"`
}

func (c ClaimFor) Encode(contract sdk.AccAddress) ([]sdk.Msg, error) {
	action, err := c.Action.ToAction()
	if err != nil {
		return nil, err
	}
	msg := claimtypes.NewMsgClaimFor(contract.String(), c.Address, action)
	return []sdk.Msg{msg}, nil
}

func Encoder(contract sdk.AccAddress, data json.RawMessage, version string) ([]sdk.Msg, error) {
	msg := &ClaimMsg{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	if msg.ClaimFor != nil {
		return msg.ClaimFor.Encode(contract)
	}
	return nil, fmt.Errorf("wasm: invalid custom claim message")
}
