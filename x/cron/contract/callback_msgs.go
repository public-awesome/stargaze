package contract

// SudoMsg callback message sent to a contract.
type SudoMsg struct {
	// This will be delivered every block for all the contracts which have been marked as privileged
	EndBlock *struct{} `json:"end_block,omitempty"`
}
