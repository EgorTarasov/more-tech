package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type DepartmentRating struct {
	Rating       float64   `json:"rating" bson:"rating" binding:"required"`
	DepartmentId string    `json:"departmentId" bson:"departmentId" binding:"required"`
	UserId       string    `json:"-" bson:"userId"`
	Text         string    `json:"text" bson:"text" binding:"required"`
	CreatedAt    time.Time `json:"-" bson:"createdAt"`
}

type RatingRepository interface {
	InsertOne(c context.Context, data DepartmentRating) error
	FindOne(c context.Context, filter bson.M) (*DepartmentRating, error)
	FindMany(c context.Context, filter bson.M) ([]DepartmentRating, error)
}
