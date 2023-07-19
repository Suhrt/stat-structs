package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Login struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Password    string             `json:"password" bson:"password"`
}