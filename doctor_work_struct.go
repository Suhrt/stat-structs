package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Session struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StartHour int                `json:"start_hour" bson:"start_hour"`
	StartMin  int                `json:"start_min" bson:"start_min"`
	EndHour   int                `json:"end_hour" bson:"end_hour"`
	EndMin    int                `json:"end_min" bson:"end_min"`
}

type Timing struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Day      string             `json:"day" bson:"day"`
	Selected bool               `json:"selected" bson:"selected"` //true if the doctor is workbng on that day
	Sessions []Session          `json:"sessions" bson:"sessions"`
}

type Work struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	WorkId        string             `json:"work_id" bson:"work_id"`
	Name          string             `json:"name" bson:"name"`
	Address       string             `json:"address" bson:"address"`
	Designation   string             `json:"designation" bson:"designation"`
	PhoneNumber   string             `json:"phone_number" bson:"phone_number"`
	Latitude      string             `json:"latitude" bson:"latitude"`
	Longitude     string             `json:"longitude" bson:"longitude"`
	IsOnline      bool               `json:"is_online" bson:"is_online"` //if practice is online
	PlaceId       string             `json:"place_id" bson:"place_id"`
	Timings       []Timing           `json:"timings" bson:"timings"`
	IsActive      bool               `json:"is_active" bson:"is_active"` //for logo templates
}
