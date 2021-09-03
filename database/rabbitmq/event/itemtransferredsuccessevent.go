package event

//mint item => TransactionHash = empty
//transfer success => ref id, reftype = empty
type ItemTransferredSuccessEvent struct {
	BaseBlockchainEvent `json:",inline"`
	ItemId              string `json:"ItemId"`
	RefId               string `json:"RefId"`
	RefType             string `json:"RefType"`
	TokenId             string `json:"TokenId"`
	FromAddress         string `json:"FromAddress"`
	ToAddress           string `json:"ToAddress"`
	Timestamp           int64  `json:"Timestamp"`
}
