package event

//mua box thanh cong => chuyen status cua box => BUY_SUCCESS + plus box amount cua user len 1
type ThetanBoxPaidEvent struct {
	BaseBlockchainEvent `json:",inline"`
	BuyerAddress        string `json:"Buyer"`
	BoxId               int64  `json:"BoxId"`
	BoxType             int    `json:"BoxType"`
	PriceInWei          string `json:"PriceInWei"`
	Price               int64  `json:"Price"` //not use?
	PaymentToken        string `json:"PaymentToken"`
	Timestamp           int64  `json:"Timestamp"`
}
