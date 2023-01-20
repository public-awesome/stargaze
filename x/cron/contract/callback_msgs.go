package contract

// TgradeSudoMsg callback message sent to a contract.
// See https://github.com/confio/tgrade-contracts/blob/main/packages/bindings/src/sudo.rs
type TgradeSudoMsg struct {
	// This will be delivered every block if the contract is currently registered for End Block
	// Block height and time is already in Env
	EndBlock *struct{} `json:"end_block,omitempty"`
}
