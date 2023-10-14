package repository

import (
	"context"
	"more-tech/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type favouriteMongoRepository struct {
	db         *mongo.Database
	collection string
}

func NewFavouriteMongoRepository(mongoDb *mongo.Database) model.FavouriteRepository {
	return &favouriteMongoRepository{
		db:         mongoDb,
		collection: "favourites",
	}
}

func (fr *favouriteMongoRepository) InsertOne(c context.Context, favouriteData model.FavouriteDepartment) error {
	_, err := fr.db.Collection(fr.collection).InsertOne(c, favouriteData)
	return err
}

func (fr *favouriteMongoRepository) FindOne(c context.Context, filter bson.M) (*model.FavouriteDepartment, error) {
	favourite := model.FavouriteDepartment{}

	err := fr.db.Collection(fr.collection).FindOne(c, filter).Decode(&favourite)
	if err != nil {
		return nil, err
	}

	return &favourite, nil
}

func (fr *favouriteMongoRepository) FindMany(c context.Context, filter bson.M) ([]model.FavouriteDepartment, error) {

	var favourites []model.FavouriteDepartment

	cursor, err := fr.db.Collection(fr.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &favourites)
	if err != nil {
		return nil, err
	}

	return favourites, nil
}

func (fr *favouriteMongoRepository) UpdateOne(c context.Context, filter bson.M, update bson.M) error {
	_, err := fr.db.Collection(fr.collection).UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (fr *favouriteMongoRepository) FindOneAndUpdate(c context.Context, filter bson.M, update bson.M) (*model.FavouriteDepartment, error) {
	var favourite model.FavouriteDepartment
	after := options.After
	upset := true
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upset,
	}
	result := fr.db.Collection(fr.collection).FindOneAndUpdate(c, filter, update, &opt)
	if result.Err() != nil {
		return nil, result.Err()
	}
	err := result.Decode(&favourite)
	if err != nil {
		return nil, err
	}
	return &favourite, nil

}

func (fr *favouriteMongoRepository) DeleteOne(c context.Context, filter bson.M) error {
	res := fr.db.Collection(fr.collection).FindOneAndDelete(c, filter)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}


