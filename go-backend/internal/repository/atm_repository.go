package repository

import (
	"context"
	"more-tech/internal/model"
	"more-tech/internal/service"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type atmMongoRepository struct {
	db         *mongo.Database
	collection string
}

func NewAtmMongoRepository(mongoDb *mongo.Database) model.AtmRepository {
	return &atmMongoRepository{
		db:         mongoDb,
		collection: "atms",
	}
}

func (ar *atmMongoRepository) FindOne(c context.Context, filter bson.M) (*model.Atm, error) {
	atm := model.Atm{}

	err := ar.db.Collection(ar.collection).FindOne(c, filter).Decode(&atm)
	if err != nil {
		return nil, err
	}

	return &atm, nil
}

func (ar *atmMongoRepository) FindMany(c context.Context, filter bson.M) ([]model.Atm, error) {
	var atms []model.Atm

	cursor, err := ar.db.Collection(ar.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(c, &atms)
	if err != nil {
		return nil, err
	}

	return atms, nil
}

func (ar *atmMongoRepository) FindRange(c context.Context, atmData model.AtmRangeRequest) ([]model.AtmRangeResponse, error) {
	var atms []model.AtmRangeResponse

	cursor, err := ar.db.Collection(ar.collection).Find(c, bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": bson.A{
					[]float64{atmData.Latitude, atmData.Longitude},
					atmData.Radius / 6380.752,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	err = cursor.All(c, &atms)
	for i := range atms {
		atms[i].Distance = service.Haversine(atmData.Longitude, atmData.Latitude, atms[i].Location.Coordinates.Longitude, atms[i].Location.Coordinates.Latitude)
	}

	sort.Slice(atms, func(i, j int) bool {
		return atms[i].Distance < atms[j].Distance
	})

	if err != nil {
		return nil, err
	}

	return atms, nil
}
