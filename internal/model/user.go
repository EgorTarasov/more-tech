package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserCreateRequest struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Fio      string `json:"fio" bson:"fio" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Fio      string `json:"fio"`
	Password string `json:"-"`
}

type UserRepository interface {
	InsertOne(c context.Context, userData UserCreateRequest) (string, error)
	FindOne(c context.Context, filter bson.M) (*UserResponse, error)
}
