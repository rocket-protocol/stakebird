package types

// curating module event types
const (
	EventTypeStake          = "stake"
	EventTypeUnstake        = "unstake"
	EventTypeBuyCreatorCoin = "buy_creator_coin"

	AttributeKeyVendorID  = "vendor_id"
	AttributeKeyPostID    = "post_id"
	AttributeKeyDelegator = "delegator"
	AttributeKeyValidator = "validator"
	AttributeKeyAmount    = "amount"

	AttributeValueCategory = ModuleName
)
