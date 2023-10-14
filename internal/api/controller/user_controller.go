package controller

import (
	"errors"
	"more-tech/internal/model"
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
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	userData.Password = string(hash)

	// accessToken, err := service.CreateAccessToken(userData.Email)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
	// 	return
	// }

	_, err = uc.repository.InsertOne(c.Request.Context(), userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	// c.SetCookie("session", , 60*60*24*400, "/", "localhost", false, true)

	c.Status(http.StatusCreated)

	// c.JSON(http.StatusCreated, model.AuthResponse{
	// 	AccessToken: accessToken,
	// 	Type:        "bearer",
	// })
}

func (uc *userController) GetUserById(c *gin.Context) {
	id := c.Param("id")

	hex_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := uc.repository.FindOne(c.Request.Context(), bson.M{"_id": hex_id})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	user.Id = id
	c.JSON(http.StatusOK, user)
}

func (uc *userController) Login(c *gin.Context) {
	userData := model.UserLoginRequest{}
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	_, err := uc.repository.FindOne(c.Request.Context(), bson.M{"email": userData.Email})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusOK)

	// accessToken, err := service.AuthenticateUser(userData, user.Password)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, accessToken)
}
