package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Department struct {
	MongoId      string `json:"_id,omitempty" bson:"_id,omitempty"`
	Id           int    `json:"id"`
	BiskvitId    string `json:"biskvitId"`
	ShortName    string `json:"shortName"`
	Address      string `json:"address"`
	City         string `json:"city"`
	ScheduleFl   string `json:"scheduleFl"`
	ScheduleJurL string `json:"scheduleJurL"`
	Special      struct {
		VipZone   int `json:"vipZone"`
		VipOffice int `json:"vipOffice"`
		Ramp      int `json:"ramp"`
		Person    int `json:"person"`
		Juridical int `json:"juridical"`
		Prime     int `json:"prime"`
	} `json:"special"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
}

type DepartmentRangeRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
}

type DepartmentRepository interface {
	FindOne(c context.Context, filter bson.M) (*Department, error)
	FindMany(c context.Context, departmentData DepartmentRangeRequest) ([]Department, error)
}
