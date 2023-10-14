package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type SearchCreate struct {
	Text        string      `json:"text" bson:"text" binding:"required" example:"текст запроса"`
	Coordinates Coordinates `json:"coordinates" bson:"coordinates" binding:"required"`
	Test        bool        `json:"test" bson:"test" example:"true"`
}

type Search struct {
	Text        string        `json:"text" bson:"text" binding:"required" example:"текст запроса"`
	UserId      string        `json:"userId" bson:"userId" binding:"required" example:"5f9e9b9b9b9b9b9b9b889b9b"`
	Coordinates Coordinates   `json:"coordinates" bson:"coordinates" binding:"required"`
	CreatedAt   time.Time     `json:"createdAt,omitempty" bson:"createdAt" example:"2021-01-01T00:00:00Z"`
	Special     SearchSpecial `json:"special" bson:"special" binding:"required"`
	Atm         bool          `json:"atm" bson:"atm" binding:"required"`
	Online      bool          `json:"online" bson:"online" binding:"required"`
}

type SearchResponse struct {
	Id          string        `json:"_id" bson:"_id" example:"5f9e9b9b9b9b9b9b9b9b9b9b"`
	Text        string        `json:"text" bson:"text" binding:"required" example:"текст запроса"`
	UserId      string        `json:"userId" bson:"userId" binding:"required" example:"5f9e9b9b9b9b9b9b9b889b9b"`
	Coordinates Coordinates   `json:"coordinates" bson:"coordinates" binding:"required"`
	CreatedAt   time.Time     `json:"createdAt" bson:"createdAt" binding:"required" example:"2021-01-01T00:00:00Z"`
	Special     SearchSpecial `json:"special" bson:"special" binding:"required"`
	Atm         bool          `json:"atm" bson:"atm" binding:"required"`
	Online      bool          `json:"online" bson:"online" binding:"required"`
}

type SearchSpecial struct {
	VipZone   bool `json:"vipZone" example:"1"`
	VipOffice bool `json:"vipOffice" example:"0"`
	Ramp      bool `json:"ramp" example:"1"`
	Person    bool `json:"person" example:"1"`
	Juridical bool `json:"juridical" example:"1"`
	Prime     bool `json:"Prime" example:"0"`
}

type SearchRepository interface {
	InsertOne(c context.Context, searchData Search) (string, error)
	FindOne(c context.Context, searchId string) (*SearchResponse, error)
	FindMany(c context.Context, filter bson.M) ([]SearchResponse, error)
}

//     "банкомат": atm,
//     "онлайн", online
//     "привелегия",  # Привелегия vipzone
//     "прайм",  # Прайм vipoffice
//     "маломобильный",  # маломобильный ramp
//     "физ лицо",  # физ лицо person
//     "юр лицо",  # юр лицо juridical
