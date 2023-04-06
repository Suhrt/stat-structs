package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customizations struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	StatId        string             `json:"stat_id" bson:"stat_id"`
	Color         []int              `json:"color" bson:"color"`
	PhoneNumber   bool               `json:"phone_number" bson:"phone_number"`
	Email         bool               `json:"email" bson:"email"`
	Address       bool               `json:"address" bson:"address"`
	PastSurgeries bool               `json:"pastSurgeries" bson:"pastSurgeries"`
	Comorbidities bool               `json:"comorbidities" bson:"comorbidities"`
	Allergies     bool               `json:"allergies" bson:"allergies"`
	HtWt          bool               `json:"htWt" bson:"htWt"`
	LogoUrl       string             `json:"logo_url" bson:"logo_url"`
	LetterHeadUrl string             `json:"letter_head_url" bson:"letter_head_url"`
	FooterUrl     string             `json:"footer_url" bson:"footer_url"`
	Disclaimer    string             `json:"disclaimer" bson:"disclaimer"`
	RawLetterHead string             `json:"raw_letter_head_url" bson:"raw_letter_head_url"`
}
