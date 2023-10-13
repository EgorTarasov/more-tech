package controller

import (
	"more-tech/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ticketController struct {
	tr model.TicketRepository
	dr model.DepartmentRepository
}

func NewTicketController(tr model.TicketRepository, dr model.DepartmentRepository) *ticketController {
	return &ticketController{
		tr: tr,
		dr: dr,
	}
}

func (tc *ticketController) CreateTicket(c *gin.Context) {
	ticket := model.Ticket{}
	if err := c.BindJSON(&ticket); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	err := tc.tr.InsertOne(c.Request.Context(), ticket)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ticket)
}
