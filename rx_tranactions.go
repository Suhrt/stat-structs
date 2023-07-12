package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type RxTransactions struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Pg     string             `bson:"pg"`
	TrxId  string             `bson:"trx_id"`
	Status string             `bson:"status"`
	RxId   string             `bson:"rx_id"`
	Amount int                `bson:"amount"`
	Url    string             `bson:"url"`
}
