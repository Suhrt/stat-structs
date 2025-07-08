package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type PhraseStruct struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Phrase string             `json:"phrase" bson:"phrase"`
	StatId string             `json:"stat_id" bson:"stat_id"`
}
