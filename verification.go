package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Verification struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	StatId          string             `json:"stat_id" bson:"stat_id"`
	Certs           []string           `json:"certs" bson:"certs"`
	Status          string             `json:"status" bson:"status"`
	RejectedMessage string             `json:"rejected_message" bson:"rejected_message"`
}
