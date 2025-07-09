package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Replace struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	StatId string             `json:"stat_id" bson:"stat_id"`
	From   string             `json:"from" bson:"from"`
	To     string             `json:"to" bson:"to"`
}
