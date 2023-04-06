package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type DrugHistory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	HistoryId string             `json:"history_id" bson:"history_id"`
	DrugIds   []string           `json:"drug_ids" bson:"drug_ids"`
}
