package event

type ItemTransferredEvent struct {
	BaseBlockchainEvent `json:",inline"`
	TokenId             string `json:"TokenId"`
}
