package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `bson:"email,omitempty"`
	PasswordHash string             `bson:"password_hash,omitempty"`
	AccessLevel  int                `bson:"access_level,omitempty"`
}
