package types

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the account types and interface
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&PromoteToPrivilegedContractProposal{}, "cron/PromoteToPrivilegedContractProposal", nil)
	cdc.RegisterConcrete(&DemotePrivilegedContractProposal{}, "cron/DemotePrivilegedContractProposal", nil)
}

func RegisterCodec(_ *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	wasmtypes.RegisterInterfaces(registry)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&PromoteToPrivilegedContractProposal{},
		&DemotePrivilegedContractProposal{},
	)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
