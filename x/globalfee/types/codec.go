package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary interfaces and concrete types on the provided LegacyAmino codec.
// These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetCodeAuthorization{}, "globalfee/MsgSetCodeAuthorization", nil)
	cdc.RegisterConcrete(&MsgRemoveCodeAuthorization{}, "globalfee/MsgRemoveCodeAuthorization", nil)
	cdc.RegisterConcrete(&MsgSetContractAuthorization{}, "globalfee/MsgSetContractAuthorization", nil)
	cdc.RegisterConcrete(&MsgRemoveContractAuthorization{}, "globalfee/MsgRemoveContractAuthorization", nil)

	cdc.RegisterConcrete(&SetCodeAuthorizationProposal{}, "globalfee/SetCodeAuthorizationProposal", nil)
	cdc.RegisterConcrete(&RemoveCodeAuthorizationProposal{}, "globalfee/RemoveCodeAuthorizationProposal", nil)
	cdc.RegisterConcrete(&SetContractAuthorizationProposal{}, "globalfee/SetContractAuthorizationProposal", nil)
	cdc.RegisterConcrete(&RemoveContractAuthorizationProposal{}, "globalfee/RemoveContractAuthorizationProposal", nil)
}

// RegisterInterfaces registers interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCodeAuthorization{},
		&MsgRemoveCodeAuthorization{},
		&MsgSetContractAuthorization{},
		&MsgRemoveContractAuthorization{},

		&SetCodeAuthorizationProposal{},
		&RemoveCodeAuthorizationProposal{},
		&SetContractAuthorizationProposal{},
		&RemoveContractAuthorizationProposal{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	ModuleCdc = codec.NewAminoCodec(Amino)
	Amino     = codec.NewLegacyAmino()
)

func init() {
	RegisterLegacyAminoCodec(Amino)
	cryptoCodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
	Amino.Seal()
}
