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
}

func NewDepartmentController(dr model.DepartmentRepository, rr model.RatingRepository) departmentController {
	return departmentController{
		dr: dr,
		rr: rr,
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
//	@Router			/v1/department/{id} [get]
func (dc *departmentController) GetDepartmentById(c *gin.Context) {
	id := c.Param("id")

	hex_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	department, err := dc.dr.FindOne(c.Request.Context(), bson.M{"_id": hex_id})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "department not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
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
//	@Router			/v1/department/range [post]
func (dc *departmentController) GetDepartmentByRange(c *gin.Context) {
	departmentData := model.DepartmentRangeRequest{}
	if err := c.BindJSON(&departmentData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	departments, err := dc.dr.FindMany(c.Request.Context(), departmentData)
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
//	@Router			/v1/department/rating [post]
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
