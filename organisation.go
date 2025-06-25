package structs

// Organisation represents an organization entity
type Organisation struct {
	ID             string `json:"id" bson:"_id"`
	Name           string `json:"name" bson:"name"`
	Description    string `json:"description" bson:"description"`
	Address        string `json:"address" bson:"address"`
	Email          string `json:"email" bson:"email"`
	Phone          string `json:"phone" bson:"phone"`
	Website        string `json:"website" bson:"website"`
	ContactPerson  string `json:"contact_person" bson:"contact_person"`
	CreatedAt      string `json:"created_at" bson:"created_at"`
	UpdatedAt      string `json:"updated_at" bson:"updated_at"`
	Active         bool   `json:"active" bson:"active"`
	OrganisationID string `json:"organisation_id" bson:"organisation_id"`
}
