package contract

// SudoMsg callback message sent to a contract.
type SudoMsg struct {
	// This will be delivered every block if the contract is currently registered for End Block
	// Block height and time is already in Env
	EndBlock *struct{} `json:"end_block,omitempty"`
}
