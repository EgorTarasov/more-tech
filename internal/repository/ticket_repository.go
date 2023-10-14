package repository

import (
	"context"
	"more-tech/internal/logging"
	"more-tech/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ticketMongoRepository struct {
	db         *mongo.Database
	collection string
}

func NewTicketMongoRepository(mongoDb *mongo.Database) model.TicketRepository {
	return &ticketMongoRepository{
		db:         mongoDb,
		collection: "tickets",
	}
}

func (tr *ticketMongoRepository) InsertOne(c context.Context, ticketData model.TicketCreate) (string, error) {
	res, err := tr.db.Collection(tr.collection).InsertOne(c, bson.M{
		"userId":       ticketData.UserId,
		"departmentId": ticketData.DepartmentId,
		"timeSlot":     ticketData.TimeSlot,
		"createdAt":    time.Now(),
	})
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (tr *ticketMongoRepository) FindOne(c context.Context, ticketId string) (*model.Ticket, error) {
	hex_id, err := primitive.ObjectIDFromHex(ticketId)
	if err != nil {
		return nil, err
	}

	ticket := model.Ticket{}

	err = tr.db.Collection(tr.collection).FindOne(c, bson.M{"_id": hex_id}).Decode(&ticket)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (tr *ticketMongoRepository) FindMany(c context.Context, filter bson.M) ([]model.Ticket, error) {
	var tickets []model.Ticket

	cursor, err := tr.db.Collection(tr.collection).Find(c, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &tickets)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (tr *ticketMongoRepository) Count(c context.Context, filter bson.M) (int64, error) {
	count, err := tr.db.Collection(tr.collection).CountDocuments(c, filter)
	logging.Log.Debug(count)
	return count, err
}

func (tr *ticketMongoRepository) DeleteOne(c context.Context, ticketId string) error {
	hex_id, err := primitive.ObjectIDFromHex(ticketId)
	if err != nil {
		return err
	}

	res, err := tr.db.Collection(tr.collection).DeleteOne(c, bson.M{"_id": hex_id})
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return err
}
