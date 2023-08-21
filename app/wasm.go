package app

import "strings"

var wasmCapabilities = []string{
	"iterator",
	"staking",
	"stargate",
	"stargaze",
	"cosmwasm_1_1",
	"cosmwasm_1_2",
	"cosmwasm_1_3",
	"token_factory",
}

func GetWasmCapabilities() string {
	return strings.Join(wasmCapabilities, ",")
}
