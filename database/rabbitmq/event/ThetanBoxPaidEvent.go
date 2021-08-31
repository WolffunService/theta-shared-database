package event

type ThetanBoxPaidEvent struct {
	BuyerAddress    string `json:"Buyer"`
	BoxType         int    `json:"BoxType"`
	Price           int64  `json:"Price"`
	PaymentToken    string `json:"PaymentToken"`
	Block           string `json:"Block"`
	TransactionHash string `json:"TransactionHash"`
	Timestamp       int64  `json:"Timestamp"`
}
