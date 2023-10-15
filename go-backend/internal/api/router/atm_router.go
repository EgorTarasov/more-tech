package router

import (
	"more-tech/internal/api/controller"
	"more-tech/internal/config"
	"more-tech/internal/repository"

	"github.com/gin-gonic/gin"
)

func (r *router) setupAtmRoutes(group *gin.RouterGroup, routeHost string) {
	atmRepo := repository.NewAtmMongoRepository(r.mongoClient.Database(config.Cfg.MongoDb))

	ac := controller.NewAtmController(atmRepo, routeHost)

	atms := group.Group("/atms")
	atms.GET("/:id", ac.GetAtmById)
	atms.POST("", ac.GetAtmByRange)
}
