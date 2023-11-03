package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/gov/codec"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	_ sdk.Msg                            = &MsgExecuteProposal{}
	_ codectypes.UnpackInterfacesMessage = &MsgExecuteProposal{}
	_ sdk.Msg                            = &MsgUpdateParams{}
)

// NewMsgExecuteProposal creates a new MsgExecuteProposal.
//
//nolint:interfacer
func NewMsgExecuteProposal(messages []sdk.Msg, authority string) (*MsgExecuteProposal, error) {
	m := &MsgExecuteProposal{
		Authority: authority,
	}

	anys, err := sdktx.SetMsgs(messages)
	if err != nil {
		return nil, err
	}

	m.Messages = anys

	return m, nil
}

// GetMsgs unpacks m.Messages Any's into sdk.Msg's
func (m *MsgExecuteProposal) GetMsgs() ([]sdk.Msg, error) {
	return sdktx.GetMsgs(m.Messages, "sdk.MsgProposal")
}

// SetMsgs packs sdk.Msg's into m.Messages Any's
// NOTE: this will overwrite any existing messages
func (m *MsgExecuteProposal) SetMsgs(msgs []sdk.Msg) error {
	anys, err := sdktx.SetMsgs(msgs)
	if err != nil {
		return err
	}

	m.Messages = anys
	return nil
}

// Type implements the sdk.Msg interface.
func (m MsgExecuteProposal) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic implements the sdk.Msg interface.
func (m MsgExecuteProposal) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.GetAuthority()); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid proposer address: %s", err)
	}

	// Check that either metadata or Msgs length is non nil.
	if len(m.Messages) == 0 {
		return errorsmod.Wrap(govtypes.ErrNoProposalMsgs, "Msgs length must be non-nil")
	}

	msgs, err := m.GetMsgs()
	if err != nil {
		return err
	}

	for idx, msg := range msgs {
		if err := msg.ValidateBasic(); err != nil {
			return errorsmod.Wrap(govtypes.ErrInvalidProposalMsg,
				fmt.Sprintf("msg: %d, err: %s", idx, err.Error()))
		}
	}

	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (m MsgExecuteProposal) GetSignBytes() []byte {
	bz := codec.ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the expected signers for a MsgExecuteProposal.
func (m MsgExecuteProposal) GetSigners() []sdk.AccAddress {
	proposer, _ := sdk.AccAddressFromBech32(m.GetAuthority())
	return []sdk.AccAddress{proposer}
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgExecuteProposal) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return sdktx.UnpackInterfaces(unpacker, m.Messages)
}

// NewMsgUpdateParams creates a new MsgUpdateParams.
//
//nolint:interfacer
func NewMsgUpdateParams(params Params, authority string) (*MsgUpdateParams, error) {
	m := &MsgUpdateParams{
		Authority: authority,
		Params:    params,
	}

	return m, nil
}

// Type implements the sdk.Msg interface.
func (m MsgUpdateParams) Type() string { return sdk.MsgTypeURL(&m) }

// ValidateBasic implements the sdk.Msg interface.
func (m MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.GetAuthority()); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid proposer address: %s", err)
	}

	return m.Params.Validate()
}

// GetSignBytes returns the message bytes to sign over.
func (m MsgUpdateParams) GetSignBytes() []byte {
	bz := codec.ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the expected signers for a MsgExecuteProposal.
func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	proposer, _ := sdk.AccAddressFromBech32(m.GetAuthority())
	return []sdk.AccAddress{proposer}
}
