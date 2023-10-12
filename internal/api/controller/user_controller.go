package controller

import (
	"errors"
	"more-tech/internal/model"
	"more-tech/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	repository model.UserRepository
}

func NewUserController(ur model.UserRepository) userController {
	return userController{
		repository: ur,
	}
}

func (uc *userController) CreateUser(c *gin.Context) {
	userData := model.UserCreateRequest{}
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userData.Password = string(hash)

	accessToken, err := service.CreateAccessToken(userData.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := uc.repository.AddOne(c.Request.Context(), userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.AuthResponse{
		AccessToken: accessToken,
		Type:        "bearer",
	})
}

func (uc *userController) GetUserById(c *gin.Context) {
	id := c.Param("id")

	hex_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.repository.FindOne(c.Request.Context(), bson.M{"_id": hex_id})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Id = id
	c.JSON(http.StatusOK, user)
}

func (uc *userController) Login(c *gin.Context) {
	userData := model.UserLoginRequest{}
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	
	user, err := uc.repository.FindOne(c.Request.Context(), bson.M{"email": userData.Email})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := service.AuthenticateUser(userData, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accessToken)
}