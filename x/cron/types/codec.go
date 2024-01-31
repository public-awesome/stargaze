package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// RegisterLegacyAminoCodec registers the account types and interface
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPromoteToPrivilegedContract{}, "cron/MsgPromoteToPrivilegedContract", nil)
	cdc.RegisterConcrete(&MsgDemoteFromPrivilegedContract{}, "cron/MsgDemoteFromPrivilegedContract", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "cron/MsgUpdateParams", nil)

	// support for legacy proposals
	cdc.RegisterConcrete(&PromoteToPrivilegedContractProposal{}, "cron/PromoteToPrivilegedContractProposal", nil)
	cdc.RegisterConcrete(&DemotePrivilegedContractProposal{}, "cron/DemotePrivilegedContractProposal", nil)
}

func RegisterCodec(_ *codec.LegacyAmino) {
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPromoteToPrivilegedContract{},
		&MsgDemoteFromPrivilegedContract{},
		&MsgUpdateParams{},
	)

	// support for legacy proposals
	registry.RegisterImplementations(
		(*v1beta1.Content)(nil),
		&PromoteToPrivilegedContractProposal{},
		&DemotePrivilegedContractProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
