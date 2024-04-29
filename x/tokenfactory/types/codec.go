package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "tokenfactory/create-denom", nil)
	cdc.RegisterConcrete(&MsgMint{}, "tokenfactory/mint", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "tokenfactory/burn", nil)
	cdc.RegisterConcrete(&MsgChangeAdmin{}, "tokenfactory/change-admin", nil)
	cdc.RegisterConcrete(&MsgSetDenomMetadata{}, "tokenfactory/set-denom-metadata", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "tokenfactory/update-params", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgMint{},
		&MsgBurn{},
		&MsgChangeAdmin{},
		&MsgSetDenomMetadata{},
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	sdk.RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
