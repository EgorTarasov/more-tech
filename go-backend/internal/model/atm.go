package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Atm struct {
	Id                string                       `json:"_id,omitempty" bson:"_id,omitempty"`
	Address           string                       `json:"address" bson:"address" binding:"required" example:"ул. Богородский Вал, д. 6, корп. 1"`
	Latitude          float64                      `json:"latitude" bson:"latitude" binding:"required" example:"55.802432"`
	Longitude         float64                      `json:"longitude" bson:"longitude" binding:"required" example:"37.704547"`
	AllDay            bool                         `json:"allDay" bson:"allDay" binding:"required" example:"true"`
	Services          map[string]map[string]string `json:"services" bson:"services" binding:"required"`
	Location          Location                     `json:"location"`
	EstimatedTimeCar  float64                      `json:"estimatedTimeCar"`
	EstimatedTimeWalk float64                      `json:"estimatedTimeWalk"`
}

type AtmRangeRequest struct {
	Latitude  float64 `json:"latitude" binding:"required" example:"55.802432"`
	Longitude float64 `json:"longitude" binding:"required" example:"37.704547"`
	Radius    float64 `json:"radius" binding:"required" example:"10"` // in km
}

type AtmRangeResponse struct {
	Id        string                       `json:"_id,omitempty" bson:"_id,omitempty"`
	Address   string                       `json:"address" bson:"address" binding:"required" example:"ул. Богородский Вал, д. 6, корп. 1"`
	Latitude  float64                      `json:"latitude" bson:"latitude" binding:"required" example:"55.802432"`
	Longitude float64                      `json:"longitude" bson:"longitude" binding:"required" example:"37.704547"`
	AllDay    bool                         `json:"allDay" bson:"allDay" binding:"required" example:"true"`
	Services  map[string]map[string]string `json:"services" bson:"services" binding:"required"`
	Location  Location                     `json:"location"`
	Distance  float64                      `json:"distance"`
}

type AtmRepository interface {
	FindOne(c context.Context, filter bson.M) (*Atm, error)
	FindMany(c context.Context, filter bson.M) ([]Atm, error)
	FindRange(c context.Context, atmData AtmRangeRequest) ([]AtmRangeResponse, error)
}
