package event

//event xác nhận việc chuyển gTHC -> THC thành công
type GameTokenTransferredSuccessEvent struct {
	BaseBlockchainEvent `json:",inline"`
	TransactionId       int64  `json:"TransactionId"`
	From                string `json:"From"`
	To                  string `json:"To"`
	AmountInWei         string `json:"AmountInWei"`
}
