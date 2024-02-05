package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Certificate struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	CertifcateId            string             `json:"certificate_id" bson:"certifcate_id"`
	DoctorId                string             `json:"doctor_id" bson:"doctor_id"`
	PatientId               string             `json:"patient_id" bson:"patient_id"`
	Password                string             `json:"password" bson:"password"`
	Date                    string             `json:"date" bson:"date"`
	Doctor                  *Doctor            `json:"doctor,omitempty" bson:"doctor,omitempty"`
	Patient                 *Patient           `json:"patient,omitempty" bson:"patient,omitempty"`
	DigitalSignature        DigitalSignature   `json:"digital_signature" bson:"digital_signature"`
	Charges                 []Charges          `json:"charges" bson:"charges"`
	Order                   *Order             `json:"order" bson:"order"`
	Template                string             `json:"template" bson:"template"`
	TemplateFields          []templateField    `json:"template_fields" bson:"template_fields"`
	Title                   string             `json:"title" bson:"title"`
	Type                    string             `json:"type" bson:"type"`
	IncludePatientSignature bool               `json:"include_patient_signature" bson:"include_patient_signature"`
	IncludePatientAddress   bool               `json:"include_patient_address" bson:"include_address"`
}
