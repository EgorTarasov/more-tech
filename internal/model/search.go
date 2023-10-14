package model

import "time"

type SearchCreateRequest struct {
	Text string `json:"text" bson:"text" binding:"required" example:"текст запроса"`
}

type SearchResponse struct {
	Id        string    `json:"id" bson:"_id" example:"5f9e9b9b9b9b9b9b9b9b9b9b"`
	Text      string    `json:"text" bson:"text" binding:"required" example:"текст запроса"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt" binding:"required" example:"2021-01-01T00:00:00Z"`
}
