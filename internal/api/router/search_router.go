package router

import (
	"more-tech/internal/api/controller"
	"more-tech/internal/config"
	"more-tech/internal/repository"

	"github.com/gin-gonic/gin"
)

func (r *router) setupSearchtRoutes(group *gin.RouterGroup) {
	searchRepo := repository.NewSearchMongoRepository(r.mongoClient.Database(config.Cfg.MongoDb))
	sc := controller.NewSearchController(searchRepo)

	search := group.Group("/search")
	search.GET("/user", sc.GetSearchRecordsForUser)
	search.GET("/:searchId", sc.GetSearchRecordById)
	search.POST("", sc.CreateSearchRecord)
}
