package event

type ItemMintedEvent struct {
	BaseBlockchainEvent `json:",inline"`
	ItemId          string `json:"ItemId"`
	TokenId         string `json:"TokenId"`
	Timestamp       int64  `json:"Timestamp"`
}
