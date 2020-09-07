package types

// user module event types
const (
	EventTypePost             = "post"
	EventTypeUpvote           = "upvote"
	EventTypeModerate         = "moderate"
	EventTypeCurationComplete = "curation_complete"

	AttributeKeyVendorID      = "vendor_id"
	AttributeKeyPostID        = "post_id"
	AttributeKeyCreator       = "creator"
	AttributeKeyCurator       = "curator"
	AttributeKeyModerator     = "moderator"
	AttributeKeyRewardAccount = "reward_account"
	AttributeKeyDeposit       = "deposit"
	AttributeKeyBody          = "body"
	AttributeCurationEndTime  = "curation_end_time"
	AttributeKeyVoteNumber    = "vote_number"
	AttributeKeyVoteAmount    = "vote_amount"
	AttributeKeyVoteDenom     = "vote_denom"

	AttributeValueCategory = ModuleName
)
