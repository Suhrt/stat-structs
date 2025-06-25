package structs

import "time"

type Appointment struct {
	ID          string    `json:"id" bson:"_id"`
	PatientID   string    `json:"patientId" bson:"patientId"`
	PatientName string    `json:"patientName" bson:"patientName"`
	PatientAge  int       `json:"patientAge" bson:"patientAge"`
	PatientSex  string    `json:"patientSex" bson:"patientSex"`
	DoctorID    string    `json:"doctorId" bson:"doctorId"`
	Date        time.Time `json:"date" bson:"date"`
	TokenNumber int       `json:"tokenNumber" bson:"tokenNumber"`
	StartTime   time.Time `json:"startTime" bson:"startTime"`
	EndTime     time.Time `json:"endTime" bson:"endTime"`
	Reason      string    `json:"reason" bson:"reason"`
	Status      string    `json:"status" bson:"status"`
	Notes       string    `json:"notes,omitempty" bson:"notes,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
