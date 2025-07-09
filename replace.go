package structs

type Replace struct {
	ID     string `json:"id" bson:"_id"`
	StatId string `json:"stat_id" bson:"stat_id"`
	From   string `json:"from" bson:"from"`
	To     string `json:"to" bson:"to"`
}
