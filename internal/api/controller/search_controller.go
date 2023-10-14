package controller

import (
	"bytes"
	"encoding/json"
	"more-tech/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type searchController struct {
	sr model.SearchRepository
}

func NewSearchController(sr model.SearchRepository) *searchController {
	return &searchController{
		sr: sr,
	}
}

// CreateSearchRecord godoc
//
//	@Summary		Create a new search record
//	@Description	Create a new search record
//	@Tags			search
//	@Accept			json
//	@Produce		json
//	@Param			search	body		model.SearchCreate	true	"Search"
//	@Success		201		{object}	string				"Search id"
//	@Failure		400		{object}	model.ErrorResponse	"Bad Request"
//	@Failure		422		{object}	model.ErrorResponse	"Unprocessable Entity"
//	@Router			/v1/search [post]
func (sc *searchController) CreateSearchRecord(c *gin.Context) {
	searchData := model.SearchCreate{}
	if err := c.BindJSON(&searchData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	userId, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	encoded, err := json.Marshal(searchData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	response, err := http.Post("http://ml:8000/service", "application/json", bytes.NewReader(encoded))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	} else if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: "ml service error"})
		return
	}

	search := model.Search{}
	err = json.NewDecoder(response.Body).Decode(&search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	search.UserId = userId

	searchId, err := sc.sr.InsertOne(c.Request.Context(), search)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, searchId)
}

// GetSearchRecordById godoc
//
//	@Summary		Get search record by id
//	@Description	Get search record by id
//	@Tags			search
//	@Accept			json
//	@Produce		json
//	@Param			searchId	path		string					true	"Search id"
//	@Success		200			{object}	model.SearchResponse	"Search"
//	@Failure		404			{object}	model.ErrorResponse		"Search not found"
//	@Router			/v1/search/{searchId} [get]
func (sc *searchController) GetSearchRecordById(c *gin.Context) {
	searchId := c.Param("searchId")

	search, err := sc.sr.FindOne(c.Request.Context(), searchId)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, search)
}

// GetSearchRecordsForUser godoc
//
//	@Summary		Get search records for user
//	@Description	Get search records for user
//	@Tags			search
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.SearchResponse	"Searches"
//	@Failure		404	{object}	model.ErrorResponse		"Searches not found"
//	@Router			/v1/search/user [get]
func (sc *searchController) GetSearchRecordsForUser(c *gin.Context) {
	userId, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	searches, err := sc.sr.FindMany(c.Request.Context(), bson.M{"userId": userId})
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, searches)
}
