package event

type TransferItemEvent struct {
	To      string `json:"To"`
	TokenId string `json:"TokenId"`
	NftId   string `json:"NftId"`
}
