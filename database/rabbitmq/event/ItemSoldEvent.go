package event

type ItemSoldEvent struct {
	ItemId          string `json:"ItemId"`
	TokenId         string `json:"TokenId"`
	PaymentTokenId  string `json:"PaymentTokenId"`
	Price           int64  `json:"Price"`
	OldUserId       string `json:"OldUserId"`
	OldOwnerAddress string `json:"OldOwnerAddress"`
	NewUserId       string `json:"NewUserId"`
	NewOwnerAddress string `json:"NewOwnerAddress"`
	Timestamp       int64  `json:"Timestamp"`
}
