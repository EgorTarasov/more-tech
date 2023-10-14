package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type SearchCreateRequest struct {
	Text string `json:"text" bson:"text" binding:"required" example:"текст запроса"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates" binding:"required"`
}

type SearchResponse struct {
	Id        string    `json:"id" bson:"_id" example:"5f9e9b9b9b9b9b9b9b9b9b9b"`
	Text      string    `json:"text" bson:"text" binding:"required" example:"текст запроса"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates" binding:"required"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt" binding:"required" example:"2021-01-01T00:00:00Z"`
}

type SearchFullResponse struct {
	Id        string    `json:"id" bson:"_id" example:"5f9e9b9b9b9b9b9b9b9b9b9b"`
	Text      string    `json:"text" bson:"text" binding:"required" example:"текст запроса"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates" binding:"required"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt" binding:"required" example:"2021-01-01T00:00:00Z"`
	Special Special `json:"special" bson:"special" binding:"required"`
}

type SearchRepository interface {
	InsertOne(c context.Context, searchData SearchCreateRequest) (string, error)
	FindOne(c context.Context, searchId string) (*SearchFullResponse, error)
	FindMany(c context.Context, filter bson.M) ([]SearchResponse, error)
}
