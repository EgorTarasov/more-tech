package repository

import (
	"context"
	"more-tech/internal/model"

	"go.mongodb.org/mongo-driver/bson"
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

func (tr *ticketMongoRepository) InsertOne(c context.Context, ticketData model.Ticket) error {
	_, err := tr.db.Collection(tr.collection).InsertOne(c, ticketData)
	return err
}

func (tr *ticketMongoRepository) FindOne(c context.Context, ticketId string) (*model.Ticket, error) {
	ticket := model.Ticket{}

	err := tr.db.Collection(tr.collection).FindOne(c, bson.M{"_id": ticketId}).Decode(&ticket)
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

func (tr *ticketMongoRepository) DeleteOne(c context.Context, ticketId string) error {
	_, err := tr.db.Collection(tr.collection).DeleteOne(c, bson.M{"_id": ticketId})
	return err
}
