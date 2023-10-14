package controller

import (
	"errors"
	"more-tech/internal/logging"
	"more-tech/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type departmentController struct {
	dr model.DepartmentRepository
	rr model.RatingRepository
	fr model.FavouriteRepository
}

func NewDepartmentController(dr model.DepartmentRepository, rr model.RatingRepository, fr model.FavouriteRepository) departmentController {
	return departmentController{
		dr: dr,
		rr: rr,
		fr: fr,
	}
}

// GetDepartmentById godoc
//
//	@Summary		Get department by id
//	@Description	Get department by id
//	@Tags			department
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Department ID"
//	@Success		200	{object}	model.Department
//	@Failure		404	{object}	model.ErrorResponse	"Department not found"
//	@Router			/v1/departments/{id} [get]
func (dc *departmentController) GetDepartmentById(c *gin.Context) {
	// estimating time leaving department
	id := c.Param("id")
	userId, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
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
	favourite, err := dc.fr.FindOne(c.Request.Context(), bson.M{"userId": userId})
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	if favourite != nil {

		favouriteSet := make(map[string]struct{})
		for _, departmentId := range favourite.DepartmentIds {
			favouriteSet[departmentId] = struct{}{}
		}
		if _, ok := favouriteSet[id]; ok {
			department.Favourite = true
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
	userId, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	departmentData := model.DepartmentRangeRequest{}
	if err := c.BindJSON(&departmentData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}
	favourite, err := dc.fr.FindOne(c.Request.Context(), bson.M{"userId": userId})
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	favouriteSet := make(map[string]struct{})
	if favourite != nil {

		for _, departmentId := range favourite.DepartmentIds {
			favouriteSet[departmentId] = struct{}{}
		}
	}

	departments, err := dc.dr.FindRange(c.Request.Context(), departmentData)
	for i := range departments {
		var ratings []model.DepartmentRating
		ratings, err = dc.rr.FindMany(c.Request.Context(), bson.M{"department_id": departments[i].MongoId})
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
		if favourite != nil {
			if _, ok := favouriteSet[departments[i].MongoId]; ok {
				departments[i].Favourite = true
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

	userId, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
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

	userId, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
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

	userID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	filter := bson.M{"userId": userID}
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
func (dc *departmentController) GetFavouriteDepartments(c *gin.Context) {
	userID, err := c.Cookie("session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	doc, err := dc.fr.FindOne(c.Request.Context(), bson.M{"userId": userID})
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
