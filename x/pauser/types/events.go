package types

const (
	EventTypeContractPaused   = "contract_paused"
	EventTypeContractUnpaused = "contract_unpaused"
	EventTypeCodeIDPaused     = "code_id_paused"
	EventTypeCodeIDUnpaused   = "code_id_unpaused"

	AttributeKeyContractAddress = "contract_address"
	AttributeKeyCodeID          = "code_id"
	AttributeKeyPausedBy        = "paused_by"
)
