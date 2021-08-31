package event

//When event buyboxwithsignature call block chain success
type BoxPaymentProcessingEvent struct {
	BaseBlockchainEvent `json:",inline"`
	BoxId        int64    `json:"BoxId"`
	BoxType      int    `json:"BoxType"`
}
