package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Patient struct {
	ID                  primitive.ObjectID  `bson:"_id,omitempty"`
	PatientId           string              `json:"patient_id" bson:"patient_id"`
	DoctorAccess        []string            `json:"doctor_access" bson:"doctor_access"`
	Name                string              `json:"name" bson:"name"`
	Age                 string              `json:"age" bson:"age"`
	Sex                 string              `json:"sex" bson:"sex"`
	Address             string              `json:"address" bson:"address"`
	Allergies           string              `json:"allergies" bson:"allergies"`
	Comorbidities       []string            `json:"comorbidities" bson:"comorbidities"`
	PastSurgeries       string              `json:"pastSurgeries" bson:"pastSurgeries"`
	Other               string              `json:"other" bson:"other"`
	HtWtRecords         []HtWt              `json:"ht_wt_records" bson:"ht_wt_records"`
	CountryCode         string              `json:"countryCode" bson:"countryCode"`
	Note                string              `json:"note" bson:"note"`
	PhoneNumber         string              `json:"phone_number" bson:"phone_number"`
	Email               string              `json:"email" bson:"email"`
	Dependent           int32               `json:"dependent" bson:"dependent"`
	LastUpdated         primitive.Timestamp `json:"last_updated" bson:"last_updated"`
	Diagnosis           string              `json:"diagnosis" bson:"diagnosis"`
	ProfileComplete     bool                `json:"profile_complete" bson:"profile_complete"`
	IdentificationMarks []string            `json:"identification_marks" bson:"identification_marks"`
}

type HtWt struct {
	ID         primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	TimeStamp  primitive.Timestamp `json:"time_stamp" bson:"time_stamp"`
	Height     string              `json:"height" bson:"height"`
	Weight     string              `json:"weight" bson:"weight"`
	HeightUnit string              `json:"heightUnit" bson:"heightUnit"`
	WeightUnit string              `json:"weightUnit" bson:"weightUnit"`
}
