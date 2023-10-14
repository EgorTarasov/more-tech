package repository

import (
	"context"
	"more-tech/internal/model"

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

func (sr *searchMongoRepository) InsertOne(c context.Context, searchData model.Search) (string, error) {
	res, err := sr.db.Collection(sr.collection).InsertOne(c, searchData)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (sr *searchMongoRepository) FindOne(c context.Context, searchId string) (*model.SearchResponse, error) {
	hex_id, err := primitive.ObjectIDFromHex(searchId)
	if err != nil {
		return nil, err
	}

	search := model.SearchResponse{}

	err = sr.db.Collection(sr.collection).FindOne(c, bson.M{"_id": hex_id}).Decode(&search)
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
