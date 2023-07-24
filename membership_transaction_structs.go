package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mtransaction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Pg              string             `bson:"pg" json:"pg"`
	TrxId           string             `bson:"trx_id" json:"trx_id"`
	Status          string             `bson:"status" json:"status"`
	StatId          string             `bson:"stat_id" json:"stat_id"`
	Amount          int                `bson:"amount" json:"amount"`
	Url             string             `bson:"url" json:"url"`
	Plan            string             `bson:"plan" json:"plan"`
	Duration        int                `bson:"duration" json:"duration"`
	Type            string             `bson:"type" json:"type"` // renew, new, upgrade
	PrevTransaction string             `bson:"prev_transaction" json:"prev_transaction"`
}
