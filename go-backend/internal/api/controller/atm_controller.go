package controller

import (
	"errors"
	"more-tech/internal/model"
	"more-tech/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type atmController struct {
	ar        model.AtmRepository
	routeHost string
}

func NewAtmController(ar model.AtmRepository, routeHost string) atmController {
	return atmController{
		ar:        ar,
		routeHost: routeHost,
	}
}

// GetAtmById godoc
//	@Summary		Get atm by id
//	@Description	Get atm by id
//	@Tags			atm
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string				true	"atm id"
//	@Param			startLongitude	query		float64				false	"start longitude"
//	@Param			startLatitude	query		float64				false	"start latitude"
//	@Success		200				{object}	model.Atm			"Atm"
//	@Failure		404				{object}	model.ErrorResponse	"Atm not found"
//	@Router			/v1/atms/{id} [get]
func (ac *atmController) GetAtmById(c *gin.Context) {
	id := c.Param("id")

	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	atm, err := ac.ar.FindOne(c.Request.Context(), bson.M{"_id": hexId})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "atm not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	startLongitude, _ := strconv.ParseFloat(c.Query("startLongitude"), 64)
	startLatitude, _ := strconv.ParseFloat(c.Query("startLatitude"), 64)
	timeCar, timeWalk, err := service.GetEstimatedTime(startLongitude, startLatitude, atm.Location.Coordinates.Longitude, atm.Location.Coordinates.Latitude, ac.routeHost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	atm.EstimatedTimeCar = timeCar
	atm.EstimatedTimeWalk = timeWalk

	c.JSON(http.StatusOK, atm)
}

// GetAtmByRange godoc
//	@Summary		Get atm by range
//	@Description	Get atm by range
//	@Tags			atm
//	@Accept			json
//	@Produce		json
//	@Param			atmData	body		model.AtmRangeRequest	true	"Atm data"
//	@Success		200		{array}		model.AtmRangeResponse	"Atm"
//	@Failure		422		{object}	model.ErrorResponse		"Unprocessable entity"
//	@Router			/v1/atms/range [post]
func (ac *atmController) GetAtmByRange(c *gin.Context) {
	var atmData model.AtmRangeRequest
	err := c.BindJSON(&atmData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	atms, err := ac.ar.FindRange(c.Request.Context(), atmData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, atms)
}
