package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payouts struct {
	ID               primitive.ObjectID  `bson:"_id,omitempty"`
	PayoutId         string              `json:"payout_id" bson:"payout_id"`
	StatId           string              `json:"stat_id" bson:"stat_id"`
	Amount           int                 `json:"amount" bson:"amount"`
	SettlementAmount int                 `json:"settlement_amount" bson:"settlement_amount"`
	RxIds            []string            `json:"rx_ids" bson:"rx_ids"`
	TimeStamp        primitive.Timestamp `json:"time_stamp" bson:"time_stamp"`
	Status           string              `json:"status" bson:"status"`
}

