package types

// stake module event types
const (
	EventTypeVoteEnd  = "voting_period_end"
	EventTypeDelegate = "delegate"
	EventTypePost     = "post"

	AttributeKeyVendorID     = "vendor_id"
	AttributeKeyPostID       = "post_id"
	AttributeKeyDelegator    = "delegator"
	AttributeKeyAmount       = "amount"
	AttributeKeyBody         = "body"
	AttributeKeyVotingPeriod = "voting_period"

	AttributeValueCategory = ModuleName
)
