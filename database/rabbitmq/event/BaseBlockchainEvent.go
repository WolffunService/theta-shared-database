package event

type BaseBlockchainEvent struct {
	Block           string `json:"Block"`
	ErrorMessage    string `json:"ErrorMessage"`
	TransactionHash string `json:"TransactionHash"`
}
