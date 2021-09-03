package event

//xac nhan transfer
type ItemTransferredEvent struct {
	BaseBlockchainEvent `json:",inline"`
	TokenId             string `json:"TokenId"`
}
