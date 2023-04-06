package structs


type Investigation struct {
	ID           string `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string `json:"name" bson:"name"`
	Components   string `json:"components" bson:"components"`
	Instructions string `json:"instructions" bson:"instructions"`
	Reviews      int    `json:"reviews" bson:"reviews"`
	Reviewers    []string `json:"reviewers" bson:"reviewers"`
}
