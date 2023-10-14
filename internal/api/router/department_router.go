package router

import (
	"more-tech/internal/api/controller"
	"more-tech/internal/config"
	"more-tech/internal/repository"

	"github.com/gin-gonic/gin"
)

func (r *router) setupDepartmentRoutes(group *gin.RouterGroup) {
	departmentRepo := repository.NewDepartmentMongoRepository(r.mongoClient.Database(config.Cfg.MongoDb))
	ratingRepo := repository.NewRatingMongoRepository(r.mongoClient.Database(config.Cfg.MongoDb))
	favouriteRepo := repository.NewFavouriteMongoRepository(r.mongoClient.Database(config.Cfg.MongoDb))
	dc := controller.NewDepartmentController(departmentRepo, ratingRepo, favouriteRepo)

	departments := group.Group("/departments")
	departments.GET("/favourite", dc.GetFavouriteDepartments)
	departments.GET("/:id", dc.GetDepartmentById)
	departments.POST("", dc.GetDepartmentByRange)
	departments.POST("/rating", dc.AddDepartmentRating)
	departments.POST("/favourite/:id", dc.AddDepartmentToFavourites)
	departments.DELETE("/favourite/:id", dc.DeleteDepartmentFromFavourites)
}
