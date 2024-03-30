package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type ManualOtp struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	OTP         string             `json:"otp" bson:"otp,omitempty"`
}
