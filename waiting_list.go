package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Booking struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PtId        string             `bson:"pt_id"`
	Name        string             `bson:"name"`
	Age         string             `bson:"age"`
	Sex         string             `bson:"sex"`
	Complaints  string             `bson:"complaint"`
	PhoneNumber string             `bson:"phone_number"`
	ReportId    []string           `bson:"report_ids"`
}

type PracticeLocation struct {
	Name     string    `bson:"name"`
	IsOnline bool      `bson:"is_online"`
	Bookings []Booking `bson:"bookings"`
}

type WaitingList struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	StatId            string             `bson:"stat_id"`
	Day               int                `bson:"day"`
	Month             int                `bson:"month"`
	Year              int                `bson:"year"`
	PracticeLocations []PracticeLocation `bson:"practice_locations"`
}
