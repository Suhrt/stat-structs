package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Prescription struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	PrescriptionId   string             `json:"prescription_id" bson:"prescription_id"`
	Doctor           *Doctor            `json:"doctor,omitempty" bson:"doctor,omitempty"`
	Patient          *Patient           `json:"patient,omitempty" bson:"patient,omitempty"`
	DoctorId         string             `json:"doctor_id" bson:"doctor_id"`
	PatientId        string             `json:"patient_id" bson:"patient_id"`
	Date             string             `json:"date" bson:"date"`
	Complaints       string             `json:"complaints" bson:"complaints"`
	History          string             `json:"history" bson:"history"`
	Examination      Examination        `json:"examination" bson:"examination"`
	Diagnosis        string             `json:"diagnosis" bson:"diagnosis"`
	Drugs            []Drug             `json:"drugs" bson:"drugs"`
	Investigations   []Investigation    `json:"investigations" bson:"investigations"`
	Advice           string             `json:"advice" bson:"advice"`
	Signature        *DisplaySignature  `json:"signature,omitempty" bson:"signature,omitempty"`
	Password         string             `json:"password" bson:"password"`
	Notes            string             `json:"notes" bson:"notes"`
	Location         Location           `json:"location" bson:"location"`
	Charges          []Charges          `json:"charges" bson:"charges"`
	Order            *Order             `json:"order" bson:"order"`
	DigitalSignature DigitalSignature   `json:"digital_signature" bson:"digital_signature"`
	HtWt             HtWt               `json:"htw" bson:"htw"`
	IsOld            bool               `json:"is_old" bson:"is_old"`
	OldUrl           string             `json:"old_url" bson:"old_url"`
}
