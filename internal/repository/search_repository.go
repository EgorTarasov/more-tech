package repository

import (
	"context"
	"more-tech/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type searchMongoRepository struct {
	db         *mongo.Database
	collection string
}

func NewSearchMongoRepository(mongoDb *mongo.Database) model.SearchRepository {
	return &searchMongoRepository{
		db:         mongoDb,
		collection: "searchHistory",
	}
}

func (sr *searchMongoRepository) InsertOne(c context.Context, searchData model.SearchCreateRequest) (string, error) {
	res, err := sr.db.Collection(sr.collection).InsertOne(c, bson.M{
		"text": searchData.Text,
		"coordinates": bson.M{
			"latitude": searchData.Coordinates.Latitude,
			"longitude": searchData.Coordinates.Longitude,
		},
		"createdAt": time.Now(),
	})
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (sr *searchMongoRepository) FindOne(c context.Context, searchId string) (*model.SearchFullResponse, error) {
	search := model.SearchFullResponse{}

	err := sr.db.Collection(sr.collection).FindOne(c, searchId).Decode(&search)
	if err != nil {
		return nil, err
	}

	return &search, nil
}

func (dr *searchMongoRepository) FindMany(c context.Context, filter bson.M) ([]model.SearchResponse, error) {
	var searches []model.SearchResponse

	cursor, err := dr.db.Collection(dr.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &searches)
	if err != nil {
		return nil, err
	}

	return searches, nil
}
