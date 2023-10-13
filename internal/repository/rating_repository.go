package repository

import (
	"context"
	"more-tech/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ratingMongoRepository struct {
	db         *mongo.Database
	collection string
}

func NewRatingMongoRepository(mongoDb *mongo.Database) *ratingMongoRepository {
	return &ratingMongoRepository{
		db:         mongoDb,
		collection: "rating",
	}
}

func (rr *ratingMongoRepository) InsertOne(c context.Context, data model.DepartmentRating) error {
	data.CreatedAt = time.Now()
	_, err := rr.db.Collection(rr.collection).InsertOne(c, data)
	if err != nil {
		return err
	}
	return nil
}

func (rr *ratingMongoRepository) FindOne(c context.Context, filter bson.M) (*model.DepartmentRating, error) {
	rating := model.DepartmentRating{}

	err := rr.db.Collection(rr.collection).FindOne(c, filter).Decode(&rating)
	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (rr *ratingMongoRepository) FindMany(c context.Context, filter bson.M) ([]model.DepartmentRating, error) {
	var ratings []model.DepartmentRating

	cursor, err := rr.db.Collection(rr.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &ratings)
	if err != nil {
		return nil, err
	}

	return ratings, nil
}
