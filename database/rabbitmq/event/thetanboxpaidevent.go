package event

type ThetanBoxPaidEvent struct {
	BaseBlockchainEvent `json:",inline"`
	BuyerAddress    string `json:"Buyer"`
	BoxType         int    `json:"BoxType"`
	Price           int64  `json:"Price"`
	PaymentToken    string `json:"PaymentToken"`
	Timestamp       int64  `json:"Timestamp"`
}
