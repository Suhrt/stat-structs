package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payouts struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	StatId string             `json:"stat_id" bson:"stat_id"`
	Amount int32             `json:"amount" bson:"amount"`
	Status string             `json:"status" bson:"status"`
}
