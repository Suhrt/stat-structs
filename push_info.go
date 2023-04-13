package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type PushInfo struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	StatId   string             `bson:"stat_id"`
	Token    string             `bson:"token"`
	Platform int              `bson:"platform"`
}
