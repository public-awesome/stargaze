package app

import "strings"

var wasmCapabilities = []string{
	"iterator", "staking", "stargate", "stargaze", "cosmwasm_1_1",
}

func GetWasmCapabilities() string {
	return strings.Join(wasmCapabilities, ",")
}
