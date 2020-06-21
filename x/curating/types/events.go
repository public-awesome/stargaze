package types

// stake module event types
const (
	EventTypeCuratingEndTime = "curating_end_time"
	EventTypeDelegate        = "delegate"
	EventTypePost            = "post"

	AttributeKeyVendorID       = "vendor_id"
	AttributeKeyPostID         = "post_id"
	AttributeKeyCreator        = "creator"
	AttributeKeyStake          = "stake"
	AttributeKeyHash           = "hash"
	AttributeKeyCurationWindow = "curation_window"

	AttributeValueCategory = ModuleName
)
