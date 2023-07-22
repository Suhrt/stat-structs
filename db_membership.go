package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type DbMembership struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	StatId    string              `bson:"stat_id" json:"stat_id"`
	Amount    int                 `bson:"amount" json:"amount"`
	TrxIds    []string            `bson:"trx_ids" json:"trx_ids"`
	StartDate primitive.Timestamp `bson:"stat_date" json:"stat_date"`
	EndDate   primitive.Timestamp `bson:"end_date" json:"end_date"`
	RefundIds []string            `bson:"refund_ids" json:"refund_ids"`
}
