package event

type ItemSaleEvent struct {
	ItemId         string `json:"ItemId"`
	TokenId        string `json:"TokenId"`
	PaymentTokenId string `json:"PaymentTokenId"`
	Price          int64  `json:"Price"`
}
