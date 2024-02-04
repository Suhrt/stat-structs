package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type CertificateTemplate struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Template       string             `json:"template" bson:"template"`
	TemplateFields []templateField    `json:"template_fields" bson:"template_fields"`
	Title          string             `json:"title" bson:"title"`
	Type           string             `json:"type" bson:"type"`
}

type templateField struct {
	Name  string `json:"name" bson:"name"`
	Type  string `json:"type" bson:"type"`
	Value string `json:"value" bson:"value"`
}

