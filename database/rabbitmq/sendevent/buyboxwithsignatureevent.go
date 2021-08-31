package sendevent

type BuyBoxWithSignatureEvent struct {
	BoxId        int64    `json:"BoxId"`
	BoxType      int    `json:"BoxType"`
	PaymentToken string `json:"PaymentToken"`
	UserAddress  string `json:"UserAddress"`
	Signature    string `json:"Signature"`
	Price        int64  `json:"Price"`
	Decimals     int    `json:"Decimals"`
}
