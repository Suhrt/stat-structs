package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type DigitalSignature struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Hash      string             `json:"hash" bson:"hash"`
	PublicKey string             `json:"publicKey" bson:"publicKey"`
}
