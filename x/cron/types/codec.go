package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

// RegisterLegacyAminoCodec registers the account types and interface
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&PromoteToPrivilegedContractProposal{}, "cron/PromoteToPrivilegedContractProposal", nil)
	cdc.RegisterConcrete(&DemotePrivilegedContractProposal{}, "cron/DemotePrivilegedContractProposal", nil)
	cdc.RegisterConcrete(&MsgPromoteToPrivilegedContract{}, "cron/MsgPromoteToPrivilegedContract", nil)
	cdc.RegisterConcrete(&MsgDemoteFromPrivilegedContract{}, "cron/MsgDemoteFromPrivilegedContract", nil)
}

func RegisterCodec(_ *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPromoteToPrivilegedContract{},
		&MsgDemoteFromPrivilegedContract{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
