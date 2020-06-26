package types

// curating module event types
const (
	EventTypeCuratingEndTime = "curating_end_time"
	EventTypeUpvote          = "upvote"
	EventTypePost            = "post"

	AttributeKeyVendorID      = "vendor_id"
	AttributeKeyPostID        = "post_id"
	AttributeKeyCreator       = "creator"
	AttributeKeyCurator       = "curator"
	AttributeKeyModerator     = "moderator"
	AttributeKeyRewardAccount = "reward_account"
	AttributeKeyDeposit       = "deposit"
	AttributeKeyBody          = "body"

	AttributeValueCategory = ModuleName
)
