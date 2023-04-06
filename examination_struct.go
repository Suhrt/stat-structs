package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Examination struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Examination  string             `json:"examination" bson:"examination"`
	PulseRate    string             `json:"pulse_rate" bson:"pulse_rate"`
	PulsePattern string             `json:"pulse_patter" bson:"pulse_patter"`
	SystolicBP   string             `json:"systolic_bp" bson:"systolic_bp"`
	DiastolicBP  string             `json:"diastolic_bp" bson:"diastolic_bp"`
	RespRate     string             `json:"resp_rate" bson:"resp_rate"`
	RespPattern  string             `json:"resp_pattern" bson:"resp_pattern"`
	Temp         string             `json:"temp" bson:"temp"`
	TempUnit     string             `json:"temp_unit" bson:"temp_unit"`
	Sp02         string             `json:"sp02" bson:"sp02"`
	Rbs          string             `json:"rbs" bson:"rbs"`
	HasVitals    bool               `json:"has_vitals" bson:"has_vitals"`
}
