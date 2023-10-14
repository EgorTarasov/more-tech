package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type TicketCreate struct {
	UserId         string  `json:"-" bson:"userId"`
	DepartmentId   string  `json:"departmentId" bson:"departmentId" example:"5f9e3b4e1d9eaf1f3125fc3f" binding:"required"`
	TimeSlot       string  `json:"timeSlot" bson:"timeSlot" example:"12:00-13:00"`
	Duration       float64 `json:"duration" bson:"duration"`
	Description    string  `json:"description" bson:"description"`
	StartLongitude float64 `json:"startLongitude"`
	StartLatitude  float64 `json:"startLatitude"`
}

type Ticket struct {
	Id           string    `json:"_id,omitempty" bson:"_id,omitempty" example:"5f9e3b4e1d9eaf1f3asdfc3f"`
	UserId       string    `json:"userId" bson:"userId" example:"5f9e3b4e1d9jnh1f3125fc3f"`
	DepartmentId string    `json:"departmentId" bson:"departmentId" example:"5f9e3b4eknjeaf1f3125fc3f"`
	TimeSlot     string    `json:"timeSlot" bson:"timeSlot" example:"12:00-13:00"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt" example:"2021-01-01T00:00:00Z"`
	Duration     float64   `json:"duration" bson:"duration"`
	Description  string    `json:"description" bson:"description"`
}

type TicketRepository interface {
	InsertOne(c context.Context, ticketData TicketCreate) (string, error)
	FindOne(c context.Context, ticketId string) (*Ticket, error)
	FindMany(c context.Context, filter bson.M) ([]Ticket, error)
	Count(c context.Context, filter bson.M) (int64, error)
	DeleteOne(c context.Context, ticketId string) error
}
