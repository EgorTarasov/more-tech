package router

import (
	"more-tech/internal/api/controller"
	"more-tech/internal/config"
	"more-tech/internal/repository"

	"github.com/gin-gonic/gin"
)

func (r *router) setupTicketRoutes(group *gin.RouterGroup) {
	ticketRepo := repository.NewTicketMongoRepository(r.mongoClient.Database(config.Cfg.MongoDb))
	departmentRepo := repository.NewDepartmentMongoRepository(r.mongoClient.Database(config.Cfg.MongoDb))

	tc := controller.NewTicketController(ticketRepo, departmentRepo)

	tickets := group.Group("/tickets")
	tickets.POST("", tc.CreateTicket)
}
