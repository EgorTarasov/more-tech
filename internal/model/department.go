package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// Есть еще записи на время, предполагается, что не может быть больше 4 записей на один час (в качестве прототипа)
type Department struct {
	MongoId           string      `json:"_id,omitempty" bson:"_id,omitempty" example:"65298f171d9eaf1f3125fc41"`
	ID                int         `json:"id" example:"29000262"`
	BiskvitID         string      `json:"Biskvit_id" example:"5010"`
	ShortName         string      `json:"shortName" example:"ДО «ЦИК «Химки-Правобережный» Филиала № 7701 Банка ВТБ (ПАО)"`
	Address           string      `json:"address" example:"Московская область, г. Химки, ул. Пролетарская, д. 8, стр. 1"`
	City              string      `json:"city" example:"Химки"`
	ScheduleFl        string      `json:"scheduleFl" example:"пн-пт: 10:00-20:00 сб: 10:00-17:00 вс: выходной"`
	ScheduleJurL      string      `json:"scheduleJurL" example:"пн-чт: 10:00-19:00 пт: 10:00-18:00 сб, вс: выходной"`
	Special           Special     `json:"special"`
	Coordinates       Coordinates `json:"coordinates"`
	Location          Location    `json:"location"`
	Workload          []Workload  `json:"workload"` // историческое
	Favourite         bool        `json:"favourite" example:"false"`
	EstimatedTimeCar  float64     `json:"estimatedTimeCar"`
	EstimatedTimeWalk float64     `json:"estimatedTimeWalk"`
	AvailableNow      bool        `json:"availableNow"`
}

type Special struct {
	VipZone   int `json:"vipZone" example:"1"`
	VipOffice int `json:"vipOffice" example:"0"`
	Ramp      int `json:"ramp" example:"1"`
	Person    int `json:"person" example:"1"`
	Juridical int `json:"juridical" example:"1"`
	Prime     int `json:"Prime" example:"0"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude" example:"55.892334"`
	Longitude float64 `json:"longitude" example:"37.44055"`
}

type Workload struct {
	Day       string         `json:"day" example:"пн"`
	LoadHours []HourWorkload `json:"loadHours"`
}

type Location struct {
	Type        string      `json:"type" example:"Point"`
	Coordinates Coordinates `json:"coordinates"`
}

type HourWorkload struct {
	Hour string  `json:"hour" example:"10:0-11:0"`
	Load float64 `json:"load" example:"0.3256373598976446"`
}

type DepartmentRangeRequest struct {
	Latitude  float64 `json:"latitude" binding:"required" example:"55.892334"`
	Longitude float64 `json:"longitude" binding:"required" example:"37.44055"`
	Radius    float64 `json:"radius" binding:"required" example:"10"` // in km
}

type DepartmentRangeResponse struct {
	MongoId      string   `json:"_id,omitempty" bson:"_id,omitempty" example:"65298f171d9eaf1f3125fc41"`
	ShortName    string   `json:"shortName" example:"ДО «ЦИК «Химки-Правобережный» Филиала № 7701 Банка ВТБ (ПАО)"`
	Addresss     string   `json:"address" bson:"address"`
	Schedulefl   string   `json:"schedulefl" example:"пн-пт: 10:00-20:00 сб: 10:00-17:00 вс: выходной"`
	Schedulejurl string   `json:"schedulejurl" example:"пн-чт: 10:00-19:00 пт: 10:00-18:00 сб, вс: выходной"`
	Distance     float64  `json:"distance" example:"0.3256373598976446"`
	Rating       float64  `json:"rating" example:"4.2"`
	Special      Special  `json:"special"`
	Location     Location `json:"location"`
	Favourite    bool     `json:"favourite" example:"true" default:"false"`
}

type DepartmentRepository interface {
	FindOne(c context.Context, filter bson.M) (*Department, error)
	FindMany(c context.Context, filter bson.M) ([]Department, error)
	FindRange(c context.Context, departmentData DepartmentRangeRequest) ([]DepartmentRangeResponse, error)
}
