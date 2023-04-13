package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Doctor struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	StatID           string             `json:"stat_id" bson:"stat_id"`
	PhoneNumber      string             `json:"phone_number" bson:"phone_number"`
	CountryCode      string             `json:"country_code" bson:"country_code"`
	Email            string             `json:"email" bson:"email"`
	Name             string             `json:"name" bson:"name"`
	Qualifications   string             `json:"qualifications" bson:"qualifications"`
	MedicalCouncil   string             `json:"medical_council" bson:"medical_council"`
	RegNo            string             `json:"reg_no" bson:"reg_no"`
	Work             []Work             `json:"work" bson:"work"`
	TaxInfo          string             `json:"tax_info" bson:"tax_info"`
	IsDoctor         bool               `json:"is_doctor" bson:"is_doctor"`
	Customizations   Customizations     `json:"customizations" bson:"customizations"`
	DisplaySignature DisplaySignature   `json:"display_signature" bson:"display_signature"`
	ProfileComplete  bool               `json:"profile_complete" bson:"profile_complete"`
	Membership       Membership         `json:"membership" bson:"membership"`
	Account          Account            `json:"account" bson:"account"`
	IsVerified       bool               `json:"is_verified" bson:"is_verified"`
	IsOldUser        bool               `json:"is_old_user" bson:"is_old_user"`
	OtpKey           string             `json:"otp_key" bson:"otp_key"`
}
