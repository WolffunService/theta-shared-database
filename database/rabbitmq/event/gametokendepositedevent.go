package event

//Khi user chuyen tien vao vi admin de nhan gTHC hoac gTHG
type GameTokenDepositedEvent struct {
	BaseBlockchainEvent `json:",inline"`
	From                string `json:"From"`
	PaymentToken        string `json:"PaymentToken"`
	AmountInWei         string `json:"AmountInWei"`
}
