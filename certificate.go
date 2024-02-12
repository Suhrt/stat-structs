package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Certificate struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	CertificateId            string             `json:"certificate_id" bson:"certificate_id"`
	DoctorId                string             `json:"doctor_id" bson:"doctor_id"`
	PatientId               string             `json:"patient_id" bson:"patient_id"`
	Password                string             `json:"password" bson:"password"`
	Title                   string             `json:"title" bson:"title"`
	Type                    string             `json:"type" bson:"type"`
	TemplateFields          []TemplateField    `json:"template_fields" bson:"template_fields"`
	IncludePatientSignature bool               `json:"include_patient_signature" bson:"include_patient_signature"`
	PatientSignature        string             `json:"patient_signature" bson:"patient_signature"`
	PatientAddress          string             `json:"patient_address" bson:"patient_address"`
	PatientIdentifiers      []string           `json:"patient_identifiers" bson:"patient_identifiers"`
	Charges                 []Charges          `json:"charges" bson:"charges"`
	Order                   *Order             `json:"order" bson:"order"`
}
