package controller

import (
	"errors"
	"more-tech/internal/logging"
	"more-tech/internal/model"
	"more-tech/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type departmentController struct {
	dr        model.DepartmentRepository
	rr        model.RatingRepository
	fr        model.FavouriteRepository
	tr        model.TicketRepository
	routeHost string
}

func NewDepartmentController(dr model.DepartmentRepository, rr model.RatingRepository, fr model.FavouriteRepository, tr model.TicketRepository, routeHost string) departmentController {
	return departmentController{
		dr:        dr,
		rr:        rr,
		fr:        fr,
		tr:        tr,
		routeHost: routeHost,
	}
}

// GetDepartmentById godoc
//
//	@Summary		Get department by id
//	@Description	Get department by id
//	@Tags			department
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Department ID"
//	@Param			startLatitude	query		string	true	"Start latitude"
//	@Param			startLongitude	query		string	true	"Start longitude"
//	@Success		200				{object}	model.Department
//	@Failure		404				{object}	model.ErrorResponse	"Department not found"
//	@Router			/v1/departments/{id} [get]
func (dc *departmentController) GetDepartmentById(c *gin.Context) {
	id := c.Param("id")
	var userId string
	var err error

	userId, err = c.Cookie("session")
	if err != nil {
		userId = c.GetString("session")
	}

	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	department, err := dc.dr.FindOne(c.Request.Context(), bson.M{"_id": hexId})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "department not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	startLongitude, _ := strconv.ParseFloat(c.Query("startLongitude"), 64)
	startLatitude, _ := strconv.ParseFloat(c.Query("startLatitude"), 64)
	timeCar, timeWalk, err := service.GetEstimatedTime(startLongitude, startLatitude, department.Coordinates.Longitude, department.Coordinates.Latitude, dc.routeHost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var minTime float64
	if timeCar < timeWalk {
		minTime = timeCar
	} else {
		minTime = timeWalk
	}

	timeSlot := service.GetClosestTimeSlot(minTime, department.Workload, "")
	logging.Log.Debugf("closest timeslot: %s", timeSlot)
	tickets, err := dc.tr.FindMany(c.Request.Context(), bson.M{"departmentId": department.MongoId, "timeSlot": timeSlot})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	logging.Log.Debugf("tickets on this timeslot: %+v", tickets)

	ticketsTime := 0.0
	for _, ticket := range tickets {
		ticketsTime += ticket.Duration
	}
	department.EstimatedTimeCar = timeCar + ticketsTime
	department.EstimatedTimeWalk = timeWalk + ticketsTime
	logging.Log.Debugf("tickets time: %f", ticketsTime)
	if ticketsTime >= 60*60 || timeSlot == "no timeslot found" {
		department.AvailableNow = false
	} else {
		department.AvailableNow = true
	}

	favourite, err := dc.fr.FindOne(c.Request.Context(), bson.M{"userId": userId})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusOK, department)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	for _, departmentId := range favourite.DepartmentIds {
		if departmentId == department.MongoId {
			department.Favourite = true
			break
		}
	}

	c.JSON(http.StatusOK, department)
}

// GetDepartmentByRange godoc
//
//	@Summary		Get department by range
//	@Description	Get department by range
//	@Tags			department
//	@Accept			json
//	@Produce		json
//	@Param			departmentData	body		model.DepartmentRangeRequest	true	"Department range request"
//	@Success		200				{object}	[]model.Department
//	@Failure		400				{object}	model.ErrorResponse	"Bad request"
//	@Failure		422				{object}	model.ErrorResponse	"Unprocessable entity"
//	@Router			/v1/departments [post]
func (dc *departmentController) GetDepartmentByRange(c *gin.Context) {
	var userId string
	var err error

	userId, err = c.Cookie("session")
	if err != nil {
		userId = c.GetString("session")
	}

	departmentData := model.DepartmentRangeRequest{}
	if err := c.BindJSON(&departmentData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	departments, err := dc.dr.FindRange(c.Request.Context(), departmentData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	for i := range departments {
		var ratings []model.DepartmentRating

		ratings, err = dc.rr.FindMany(c.Request.Context(), bson.M{"departmentId": departments[i].MongoId})
		if err != nil {
			break
		}

		avgRating := 0.0
		for _, rating := range ratings {
			avgRating += rating.Rating
		}

		if len(ratings) > 0 {
			avgRating /= float64(len(ratings))
		}

		departments[i].Rating = avgRating
	}

	favourite, err := dc.fr.FindOne(c.Request.Context(), bson.M{"userId": userId})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusOK, departments)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	for idx, department := range departments {
		for _, departmentId := range favourite.DepartmentIds {
			if department.MongoId == departmentId {
				departments[idx].Favourite = true
				break
			}
		}
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, departments)
}

// AddDepartmentRating godoc
//
//	@Summary		Add department rating
//	@Description	Add department rating
//	@Tags			department
//	@Accept			json
//	@Produce		json
//	@Param			ratingData	body		model.DepartmentRating	true	"Department rating"
//	@Success		200			{object}	string					"Rating added"
//	@Failure		400			{object}	model.ErrorResponse		"Bad request"
//	@Failure		422			{object}	model.ErrorResponse		"Unprocessable entity"
//	@Router			/v1/departments/rating [post]
func (dc *departmentController) AddDepartmentRating(c *gin.Context) {
	ratingData := model.DepartmentRating{}
	if err := c.BindJSON(&ratingData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	var userId string
	var err error

	userId, err = c.Cookie("session")
	if err != nil {

		userId = c.GetString("session")

	}
	ratingData.UserId = userId
	logging.Log.Debug(ratingData)

	err = dc.rr.InsertOne(c.Request.Context(), ratingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ratingData)
}

// AddDepartmentToFavourites godoc
//
//	@Summary		Add department to favourites
//	@Description	Add department to favourites
//	@Tags			department
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"Department ID"
//	@Success		200	{object}	string				"Added to favourites"
//	@Failure		400	{object}	model.ErrorResponse	"Bad request"
//	@Failure		422	{object}	model.ErrorResponse	"Unprocessable entity"
//	@Router			/v1/departments/favourite/{id} [post]
func (dc *departmentController) AddDepartmentToFavourites(c *gin.Context) {
	departmentId := c.Param("id")

	var userId string
	var err error

	userId, err = c.Cookie("session")
	if err != nil {

		userId = c.GetString("session")

	}

	doc, err := dc.fr.FindOne(c.Request.Context(), bson.M{"userId": userId})
	logging.Log.Debugf("error: %+v", err)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	logging.Log.Debug(doc)
	if doc != nil {
		err := dc.fr.UpdateOne(c.Request.Context(), bson.M{"userId": userId}, bson.M{"$addToSet": bson.M{"departmentIds": departmentId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, "Added to favourites")
	} else {
		err := dc.fr.InsertOne(c.Request.Context(), model.FavouriteDepartment{UserId: userId, DepartmentIds: []string{departmentId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, "Added to favourites")
	}
}

// GetFavouriteDepartments godoc
//
//	@Summary		Deletes department from favourites
//	@Description	Deletes department from favourites
//	@Tags			department
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Department ID"
//	@Success		200	{object}	string	"Deleted from favourites"
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		422	{object}	model.ErrorResponse
//	@Router			/v1/departments/favourite/{id} [delete]
func (dc *departmentController) DeleteDepartmentFromFavourites(c *gin.Context) {
	departmentID := c.Param("id")

	var userId string
	var err error

	userId, err = c.Cookie("session")
	if err != nil {

		userId = c.GetString("session")

	}

	filter := bson.M{"userId": userId}
	update := bson.M{"$pull": bson.M{"departmentIds": departmentID}}

	doc, err := dc.fr.FindOneAndUpdate(c.Request.Context(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	if len(doc.DepartmentIds) == 0 {
		err = dc.fr.DeleteOne(c.Request.Context(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, "Deleted from favourites")
		return
	}
	c.JSON(http.StatusOK, "Deleted from favourites")
}

// GetFavouriteDepartments godoc
//
//	@Summary		Get favourite departments
//	@Description	Get favourite departments
//	@Tags			department
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.Department
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		422	{object}	model.ErrorResponse
//	@Router			/v1/departments/favourite [get]
func (dc *departmentController) GetFavouriteDepartments(c *gin.Context) {
	var userId string
	var err error

	userId, err = c.Cookie("session")
	if err != nil {

		userId = c.GetString("session")

	}

	doc, err := dc.fr.FindOne(c.Request.Context(), bson.M{"userId": userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	hexDepartmentIds := make([]primitive.ObjectID, len(doc.DepartmentIds))
	for i, departmentId := range doc.DepartmentIds {
		hexDepartmentIds[i], err = primitive.ObjectIDFromHex(departmentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
			return
		}
	}

	departments, err := dc.dr.FindMany(c.Request.Context(), bson.M{"_id": bson.M{"$in": hexDepartmentIds}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, departments)
}
