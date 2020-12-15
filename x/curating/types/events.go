package types

// curating module event types
const (
	EventTypePost             = "post"
	EventTypeUpvote           = "upvote"
	EventTypeModerate         = "moderate"
	EventTypeCurationComplete = "curation_complete"
	EventTypeProtocolReward   = "protocol_reward"
	EventTypeVotingPoolReturn = "voting_pool_return"

	AttributeKeyVendorID           = "vendor_id"
	AttributeKeyPostID             = "post_id"
	AttributeKeyCreator            = "creator"
	AttributeKeyCurator            = "curator"
	AttributeKeyRewardAccount      = "reward_account"
	AttributeKeyDeposit            = "deposit"
	AttributeKeyBody               = "body"
	AttributeCurationEndTime       = "curation_end_time"
	AttributeKeyVoteNumber         = "vote_number"
	AttributeKeyVoteAmount         = "vote_amount"
	AttributeKeyVoteDenom          = "vote_denom"
	AttributeKeyProtocolRewardType = "reward_type"

	AttributeRewardTypeCreator = "creator"
	AttributeRewardTypeCurator = "curator"

	AttributeKeyDelegator   = "delegator"
	AttributeKeyValidator   = "validator"
	AttributeKeyStakeAmount = "stake_amount"

	AttributeValueCategory = ModuleName
)
