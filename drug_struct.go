package structs

type Drug struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	BrandName    string `json:"brand_name" bson:"brand_name"`
	GenericName  string `json:"generic_name" bson:"generic_name"`
	Dose         string `json:"dose" bson:"dose"`
	DoseUnit     string `json:"dose_unit" bson:"dose_unit"`
	Frequency    string `json:"frequency" bson:"frequency"`
	Duration     string `json:"duration" bson:"duration"`
	DurationUnit string `json:"duration_unit" bson:"duration_unit"`
	Instructions string `json:"instructions" bson:"instructions"`
	Quantity     string `json:"quantity" bson:"quantity"`
	Rxcui        string `json:"rxcui" bson:"rxcui"`
	Custom       string `json:"custom" bson:"custom"`

	PackQuantity string `json:"pack_quantity" bson:"pack_quantity"`
	Manufacturer string `json:"manufacturer" bson:"manufacturer"`
	Price        string `json:"price" bson:"price"`
	Country      string `json:"country" bson:"country"`

	Source    string   `json:"source" bson:"source"`
	Reviews   int      `json:"reviews" bson:"reviews"`
	Reviewers []string `json:"reviewers" bson:"reviewers"`
}
