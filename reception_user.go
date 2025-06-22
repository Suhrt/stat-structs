package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

// Receptionist represents a receptionist working at an organization
type Receptionist struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	StaffID        string             `json:"staff_id" bson:"staff_id"`
	Name           string             `json:"name" bson:"name"`
	Email          string             `json:"email" bson:"email"`
	Phone          string             `json:"phone" bson:"phone"`
	OrganisationID string             `json:"organisation_id" bson:"organisation_id"`
	PasswordHash   string             `json:"-" bson:"password_hash"` // Bcrypt hash of the password
	Designation    string             `json:"designation" bson:"designation"`
	Status         string             `json:"status" bson:"status"`
	IsActive       bool               `json:"is_active" bson:"is_active"`
	CreatedAt      string             `json:"created_at" bson:"created_at"`
	UpdatedAt      string             `json:"updated_at" bson:"updated_at"`
}
