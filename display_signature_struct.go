package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type DisplaySignature struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	SignatureUrl  string             `json:"signature_url" bson:"signature_url"`
	SignatureName string             `json:"signature_name" bson:"signature_name"`
	SignatureFont string             `json:"signature_font" bson:"signature_font"`
}
