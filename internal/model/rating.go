package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type DepartmentRating struct {
	Rating       float64 `json:"rating"`
	DepartmentId string  `json:"departmentId"`
	UserId       string  `json:"userId"`
	Text         string  `json:"text"`
	CreatedAt    time.Time
}
type RatingRepository interface {
	InsertOne(c context.Context, data DepartmentRating) error
	FindOne(c context.Context, filter bson.M) (*DepartmentRating, error)
	FindMany(c context.Context, filter bson.M) ([]DepartmentRating, error)
}
