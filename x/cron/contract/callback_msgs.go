package contract

// SudoMsg callback message sent to a contract.
type SudoMsg struct {
	// This will be delivered beginning of every block for all the contracts which have been marked as privileged
	BeginBlock *struct{} `json:"begin_block,omitempty"`
	// This will be delivered end of every block for all the contracts which have been marked as privileged
	EndBlock *struct{} `json:"end_block,omitempty"`
}
