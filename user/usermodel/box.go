package usermodel

import "thetansm/models/currencymodel"

type BoxEvent struct {
	ID               interface{}       `json:"id" bson:"_id,omitempty"`
	UserId           string            `json:"userId" bson:"userId"`
	BoxId            int64             `json:"boxId" bson:"boxId"`
	BoxType          int32             `json:"boxType" bson:"boxType"`
	Timestamp        int64             `json:"timestamp" bson:"timestamp"`
	BoxPurchaseEvent *BoxPurchaseEvent `json:"boxPurchaseEvent,omitempty" bson:"boxPurchaseEvent,omitempty"`
}

type BoxPurchaseEvent struct {
	Price         currencymodel.SystemCurrency `json:"price" bson:"price"`
	Amount        int32                        `json:"amount" bson:"amount"`
	TransactionId string                       `json:"-" bson:"transactionId"`
}
