package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

var amino = codec.NewLegacyAmino()

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// RegisterInterfaces registers interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	//msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
