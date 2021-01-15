package types

// curating module event types
const (
	EventTypePost             = "post"
	EventTypeUpvote           = "upvote"
	EventTypeCurationComplete = "curation_complete"
	EventTypeProtocolReward   = "protocol_reward"

	AttributeKeyVendorID           = "vendor_id"
	AttributeKeyPostID             = "post_id"
	AttributeKeyCreator            = "creator"
	AttributeKeyCurator            = "curator"
	AttributeKeyRewardAccount      = "reward_account"
	AttributeKeyBody               = "body"
	AttributeCurationEndTime       = "curation_end_time"
	AttributeKeyVoteNumber         = "vote_number"
	AttributeKeyVoteAmount         = "vote_amount"
	AttributeKeyVoteDenom          = "vote_denom"
	AttributeKeyProtocolRewardType = "reward_type"

	AttributeValueCategory = ModuleName
)
