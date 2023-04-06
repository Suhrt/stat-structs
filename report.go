package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	ID            primitive.ObjectID  `bson:"_id,omitempty"`
	ReportId      string              `bson:"report_id"`
	PatientId     string              `bson:"patient_id"`
	DoctorsAccess []string            `bson:"doctors_access"`
	RxId          string              `bson:"rx_id"`
	UploadDate    primitive.Timestamp `bson:"uplaod_date"`
	Seen          bool                `bson:"seen"`
	Format        int                 `bson:"format"`
	Url           string              `bson:"url"`
	Remarks       []Remark            `bson:"remarks"`
}

type Remark struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserId   string             `bson:"user_id"`
	IsDoctor bool               `bson:"is_doctor"`
	Comment  string             `bson:"comment"`
}
