package router

import (
	"more-tech/internal/api/controller"
	"more-tech/internal/config"
	"more-tech/internal/repository"

	"github.com/gin-gonic/gin"
)

func (r *router) setupUserRoutes(group *gin.RouterGroup) {
	userRepo := repository.NewUserMongoRepository(r.mongoClient.Database(config.Cfg.MongoDb))
	uc := controller.NewUserController(userRepo)
	
	users := group.Group("/users")
	users.GET("/:id", uc.GetUserById)
	users.POST("/create", uc.CreateUser)
	users.POST("/login", uc.Login)
}
