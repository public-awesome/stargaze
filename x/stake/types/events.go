package types

// curating module event types
const (
	EventTypeStake           = "stake"
	EventTypeUnstake         = "unstake"
	EventTypeBuyCreatorCoin  = "buy_creator_coin"
	EventTypeSellCreatorCoin = "sell_creator_coin"

	AttributeKeyVendorID  = "vendor_id"
	AttributeKeyPostID    = "post_id"
	AttributeKeyUsername  = "username"
	AttributeKeyCreator   = "creator"
	AttributeKeyBuyer     = "buyer"
	AttributeKeySeller    = "seller"
	AttributeKeyDelegator = "delegator"
	AttributeKeyValidator = "validator"
	AttributeKeyAmount    = "amount"

	AttributeValueCategory = ModuleName
)
