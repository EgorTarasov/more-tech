package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Ticket struct {
	UserId       string `json:"userId" example:"5f9e3b4e1d9eaf1f3125fc3f"`
	DepartmentId string `json:"departmentId" example:"5f9e3b4e1d9eaf1f3125fc3f"`
	TimeSlot     string `json:"timeSlot" example:"2020-11-02T10:00:00.000Z"`
}

type TicketRepository interface {
	InsertOne(c context.Context, ticketData Ticket) error
	FindOne(c context.Context, ticketId string) (*Ticket, error)
	FindMany(c context.Context, filter bson.M) ([]Ticket, error)
	DeleteOne(c context.Context, ticketId string) error
}
