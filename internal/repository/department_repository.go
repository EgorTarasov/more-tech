package repository

import (
	"context"
	"more-tech/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type departmentMongoRepository struct {
	db         *mongo.Database
	collection string
}

func NewDepartmentMongoRepository(mongoDb *mongo.Database) model.DepartmentRepository {
	return &departmentMongoRepository{
		db:         mongoDb,
		collection: "departments",
	}
}

func (dr *departmentMongoRepository) FindOne(c context.Context, filter bson.M) (*model.Department, error) {
	department := model.Department{}

	err := dr.db.Collection(dr.collection).FindOne(c, filter).Decode(&department)
	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (dr *departmentMongoRepository) FindMany(c context.Context, departmentData model.DepartmentRangeRequest) ([]model.Department, error) {
	var departments []model.Department

	cursor, err := dr.db.Collection(dr.collection).Find(c, bson.M{
		"location": bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": bson.A{
					[]float64{departmentData.Latitude, departmentData.Longitude},
					departmentData.Radius / 6380.752,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &departments)
	if err != nil {
		return nil, err
	}

	return departments, nil
}
