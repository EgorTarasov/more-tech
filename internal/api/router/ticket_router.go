package router

import (
	"more-tech/internal/api/controller"
	"more-tech/internal/config"
	"more-tech/internal/repository"

	"github.com/gin-gonic/gin"
)

func (r *router) setupTicketRoutes(group *gin.RouterGroup) {
	ticketRepo := repository.NewTicketMongoRepository(r.mongoClient.Database(config.Cfg.MongoDb))
	tc := controller.NewTicketController(ticketRepo)

	tickets := group.Group("/tickets")
	tickets.GET("/user", tc.GetTicketsForUser)
	tickets.GET("/department/:departmentId", tc.GetTicketsForDepartment)
	tickets.GET("/:ticketId", tc.GetTicketById)
	tickets.DELETE("/:ticketId", tc.CancelTicket)
	tickets.POST("", tc.CreateTicket)
}
