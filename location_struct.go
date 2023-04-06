package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`
}