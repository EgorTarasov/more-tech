package controller

import (
	"errors"
	"more-tech/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type departmentController struct {
	repository model.DepartmentRepository
}

func NewDepartmentController(dr model.DepartmentRepository) departmentController {
	return departmentController{
		repository: dr,
	}
}

func (dc *departmentController) GetDepartmentById(c *gin.Context) {
	id := c.Param("id")

	hex_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	department, err := dc.repository.FindOne(c.Request.Context(), bson.M{"_id": hex_id})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusNotFound, gin.H{"error": "department not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, department)
}

func (dc *departmentController) GetDepartmentByRange(c *gin.Context) {
	departmentData := model.DepartmentRangeRequest{}
	if err := c.BindJSON(&departmentData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	departments, err := dc.repository.FindMany(c.Request.Context(), departmentData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, departments)
}
