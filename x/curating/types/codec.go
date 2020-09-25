package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
}

// RegisterInterfaces register the curating module interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPost{},
		&MsgUpvote{},
	)
}

var (
	// ModuleCdc references the global x/curating module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding.
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)
