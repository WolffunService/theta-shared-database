package sendevent

type TransferGameTokenEvent struct {
	TransactionId int64  `json:"TransactionId"`
	To            string `json:"To"`
	AmountInWei   string `json:"AmountInWei"`
	PaymentToken  string `json:"PaymentToken"`
}