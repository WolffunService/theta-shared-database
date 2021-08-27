package event

type ItemTransferredSuccessEvent struct {
	BaseBlockchainEvent `json:",inline"`
	TokenId             string `json:"TokenId"`
	FromAddress         string `json:"FromAddress"`
	ToAddress           string `json:"ToAddress"`
	Timestamp           int64  `json:"Timestamp"`
}
