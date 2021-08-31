package event

//khi 1 user sale 1 item
//=> change status => saleing
type ItemMintedEvent struct {
	BaseBlockchainEvent `json:",inline"`
	ItemId          string `json:"ItemId"`
	TokenId         string `json:"TokenId"`
	Timestamp       int64  `json:"Timestamp"`
}
