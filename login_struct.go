package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Login struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	Password    string             `json:"password" bson:"password,omitempty"`
	StatId      string             `json:"stat_id" bson:"stat_id,omitempty"`
	OtpKey      string             `json:"otp_key" bson:"otp_key,omitempty"`
}
