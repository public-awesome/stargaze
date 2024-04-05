package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "tokenfactory/create-denom", nil)
	cdc.RegisterConcrete(&MsgMint{}, "tokenfactory/mint", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "tokenfactory/burn", nil)
	cdc.RegisterConcrete(&MsgChangeAdmin{}, "tokenfactory/change-admin", nil)
	cdc.RegisterConcrete(&MsgForceTransfer{}, "tokenfactory/force-transfer", nil)
	cdc.RegisterConcrete(&MsgSetBeforeSendHook{}, "osmosis/tokenfactory/set-beforesend-hook", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgMint{},
		&MsgBurn{},
		&MsgChangeAdmin{},
		&MsgForceTransfer{},
		&MsgSetBeforeSendHook{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyCodec(amino)
	sdk.RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
