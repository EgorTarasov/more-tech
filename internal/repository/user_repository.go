package repository

import (
	"context"
	"more-tech/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongoRepository struct {
	db         *mongo.Database
	collection string
}

func NewUserMongoRepository(mongoDb *mongo.Database) model.UserRepository {
	return &userMongoRepository{
		db:         mongoDb,
		collection: "users",
	}
}

func (ur *userMongoRepository) InsertOne(c context.Context, userData model.UserCreateRequest) error {
	encoded, err := bson.Marshal(userData)
	if err != nil {
		return err
	}
	_, err = ur.db.Collection(ur.collection).InsertOne(c, encoded)
	return err
}

func (ur *userMongoRepository) FindOne(c context.Context, filter bson.M) (*model.UserResponse, error) {
	user := model.UserResponse{}

	err := ur.db.Collection(ur.collection).FindOne(c, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
