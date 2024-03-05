package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	OrderId         string             `json:"order_id" bson:"order_id"`
	Amount          int                `json:"amount" bson:"amount"`
	DisplayAmount   string             `json:"displayAmount" bson:"displayAmount"`
	ByCash          bool               `json:"by_cash" bson:"by_cash"`
	PaymentComplete bool               `json:"paymentComplete" bson:"paymentComplete"`
	PaymentId       string             `json:"paymentId" bson:"paymentId"`
	Fee             int64              `json:"fee" bson:"fee"`
	PayedOut        string             `json:"payedOut" bson:"payedOut"`
	RefundId        string             `json:"refundId" bson:"refundId"`
	EnableOnlinePay bool               `json:"enableOnlinePay" bson:"enableOnlinePay"`
	AccessBeforePay bool               `json:"accessBeforePay" bson:"accessBeforePay"`
}
