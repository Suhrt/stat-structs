package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Charges struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ChargeID     string             `json:"charge_id" bson:"charge_id"`
	StatId       string             `json:"stat_id" bson:"stat_id"`
	Name         string             `json:"name" bson:"name"`
	Amount       string             `json:"amount" bson:"amount"`
	Currency     string             `json:"currency" bson:"currency"`
	Taxed        bool               `json:"taxed" bson:"taxed"`
	TaxName      string             `json:"tax_name" bson:"tax_name"`
	TaxPercent   string             `json:"tax_percent" bson:"tax_percent"`
	TaxInclusive bool               `json:"tax_inclusive" bson:"tax_inclusive"`
	Quantity     int32              `json:"quantity" bson:"quantity"`
	Total        string             `json:"total" bson:"total"`
	CurrencyCode string             `json:"currency_code" bson:"currency_code"`
}
