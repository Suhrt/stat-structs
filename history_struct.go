package structs

type History struct {
	StatId    string   `json:"stat_id" bson:"stat_id"`
	Diagnosis string   `json:"diagnosis" bson:"diagnosis"`
	DrugIds   []string `json:"drug_ids" bson:"drug_ids"`
}