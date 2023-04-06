package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Membership struct {
	ProStatus int32               `json:"pro_status" bson:"pro_status"`
	StartDate primitive.Timestamp `json:"start_date" bson:"start_date"`
	EndDate   primitive.Timestamp `json:"end_date" bson:"end_date"`
	OrderId   string              `json:"order_id" bson:"order_id"`
	PaymentId string              `json:"payment_id" bson:"payment_id"`
	Amount    int                 `json:"amount" bson:"amount"`
	RefundIds  []string              `json:"refund_id" bson:"refund_id"`
}
