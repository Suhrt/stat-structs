package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID             string              `json:"id" bson:"_id"`
	PatientID      string              `json:"patient_id" bson:"patient_id"`
	PatientName    string              `json:"patient_name" bson:"patient_name"`
	PatientAge     string              `json:"patient_age" bson:"patient_age"`
	PatientSex     string              `json:"patient_sex" bson:"patient_sex"`
	PatientPhone   string              `json:"patient_phone" bson:"patient_phone"`
	DoctorName     string              `json:"doctor_name" bson:"doctor_name"`
	DoctorID       string              `json:"doctor_id" bson:"doctor_id"`
	Date           time.Time           `json:"date" bson:"date"`
	TokenNumber    int                 `json:"token_number" bson:"token_number"`
	StartTime      time.Time           `json:"start_time" bson:"start_time"`
	EndTime        time.Time           `json:"end_time" bson:"end_time"`
	Reason         string              `json:"reason" bson:"reason"`
	Status         string              `json:"status" bson:"status"`
	Notes          string              `json:"notes,omitempty" bson:"notes,omitempty"`
	UpdatedAt      time.Time           `json:"updated_at" bson:"updated_at"`
	OrganisationID string              `json:"organisation_id" bson:"organisation_id"`
	LastUpdated    primitive.Timestamp `json:"last_updated" bson:"last_updated"`
}
