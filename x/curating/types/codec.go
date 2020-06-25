package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgPost{}, "curating/MsgPost", nil)
}

// var (
// 	amino = codec.New()

// 	// ModuleCdc references the global x/distribution module codec. Note, the codec
// 	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
// 	// is still used for that purpose.
// 	//
// 	// The actual codec used for serialization should be provided to x/distribution and
// 	// defined at the application level.
// 	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
// )

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
