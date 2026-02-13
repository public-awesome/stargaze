package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary interfaces and concrete types on the provided LegacyAmino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPauseContract{}, "pauser/MsgPauseContract", nil)
	cdc.RegisterConcrete(&MsgUnpauseContract{}, "pauser/MsgUnpauseContract", nil)
	cdc.RegisterConcrete(&MsgPauseCodeID{}, "pauser/MsgPauseCodeID", nil)
	cdc.RegisterConcrete(&MsgUnpauseCodeID{}, "pauser/MsgUnpauseCodeID", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "pauser/MsgUpdateParams", nil)
}

// RegisterInterfaces registers interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPauseContract{},
		&MsgUnpauseContract{},
		&MsgPauseCodeID{},
		&MsgUnpauseCodeID{},
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var Amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(Amino)
	cryptoCodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
	Amino.Seal()
}
