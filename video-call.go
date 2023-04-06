package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type VideoCall struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	RoomID string `bson:"room_id"`
	PatientToken string `bson:"patient_token"`
	DoctorToken string `bson:"doctor_token"`
}