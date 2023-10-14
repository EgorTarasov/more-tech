package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type FavouriteDepartment struct {
	UserId        string   `json:"userId" bson:"userId"`
	DepartmentIds []string `json:"departmentIds" bson:"departmentIds"`
}

type FavouriteRepository interface {
	InsertOne(c context.Context, favouriteData FavouriteDepartment) error
	FindOne(c context.Context, filter bson.M) (*FavouriteDepartment, error)
	FindMany(c context.Context, filter bson.M) ([]FavouriteDepartment, error)
	FindOneAndUpdate(c context.Context, filter bson.M, update bson.M) (*FavouriteDepartment, error)
	UpdateOne(c context.Context, filter bson.M, update bson.M) error
	DeleteOne(c context.Context, filter bson.M) error
}
