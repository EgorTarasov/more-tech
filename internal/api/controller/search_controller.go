package controller

import (
	"more-tech/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type searchController struct {
	sr model.SearchRepository
}

func NewSearchController(sr model.SearchRepository) *searchController {
	return &searchController{
		sr: sr,
	}
}

func (sc *searchController) CreateSearchRecord(c *gin.Context) {
	searchData := model.SearchCreateRequest{}
	if err := c.BindJSON(&searchData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	searchId, err := sc.sr.InsertOne(c.Request.Context(), searchData)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, searchId)
}

func (sc *searchController) GetSearchRecordById(c *gin.Context) {
	searchId := c.Param("searchId")

	search, err := sc.sr.FindOne(c.Request.Context(), searchId)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, search)
}
