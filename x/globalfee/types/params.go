package types

import (
	fmt "fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyPrivilegedAddresses = []byte("PrivilegedAddresses")
	KeyMinGasPrices        = []byte("MinimumGasPricesParam")
)

var _ paramtypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// default module parameters
func DefaultParams() Params {
	return Params{
		PrivilegedAddresses: []string{},
		MinimumGasPrices:    sdk.NewDecCoins(),
	}
}

func NewParams(addresses []string) Params {
	return Params{
		PrivilegedAddresses: addresses,
		MinimumGasPrices:    sdk.NewDecCoins(),
	}
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPrivilegedAddresses, &p.PrivilegedAddresses, validatePriviligedAddresses),
		paramtypes.NewParamSetPair(KeyMinGasPrices, &p.MinimumGasPrices, validateMinimumGasPrices),
	}
}

func (p Params) Validate() error {
	if err := validatePriviligedAddresses(p.PrivilegedAddresses); err != nil {
		return err
	}
	return validateMinimumGasPrices(p.MinimumGasPrices)
}

func validatePriviligedAddresses(i interface{}) error {
	privilegedAddress, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, addr := range privilegedAddress {
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateMinimumGasPrices(i interface{}) error {
	v, ok := i.(sdk.DecCoins)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "type: %T, expected sdk.DecCoins", i)
	}

	return v.Validate()
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

var _ sdk.Msg = &MsgUpdateParams{}

// msg types
const (
	TypeMsgUpdateParams = "update_params"
)

func NewMsgUpdateParams(sender string, minimumGasPrices sdk.DecCoins, privilegedAddresses []string) *MsgUpdateParams {
	return &MsgUpdateParams{
		Sender: sender,
		Params: Params{
			MinimumGasPrices:    minimumGasPrices,
			PrivilegedAddresses: privilegedAddresses,
		},
	}
}

func (msg MsgUpdateParams) Route() string {
	return RouterKey
}

func (msg MsgUpdateParams) Type() string {
	return TypeMsgUpdateParams
}

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgUpdateParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return msg.Params.Validate()
}
