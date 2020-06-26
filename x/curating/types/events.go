package types

// curating module event types
const (
	EventTypePost            = "post"
	EventTypeUpvote          = "upvote"
	EventTypeCuratingEndTime = "curating_end_time"

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
