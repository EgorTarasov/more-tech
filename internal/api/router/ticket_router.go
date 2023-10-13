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
	tickets.GET("/department/:departmentId", tc.GetTicketsForDepartment)
	tickets.GET("/user/:userId", tc.GetTicketsForUser)
	tickets.GET("/:ticketId", tc.GetTicketById)
	tickets.DELETE("/:ticketId", tc.CancelTicket)
	tickets.POST("", tc.CreateTicket)
}
