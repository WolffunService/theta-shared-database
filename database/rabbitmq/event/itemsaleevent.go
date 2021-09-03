package event

//1 item được dua len market place de sale
type ItemSaleEvent struct {
	ItemId         string `json:"ItemId"`
	TokenId        string `json:"TokenId"`
	PaymentTokenId string `json:"PaymentTokenId"`
	Price          int64  `json:"Price"`
}
