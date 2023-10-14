package controller

import (
	"more-tech/internal/config"
	"more-tech/internal/logging"
	"more-tech/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ticketController struct {
	tr model.TicketRepository
}

func NewTicketController(tr model.TicketRepository) *ticketController {
	return &ticketController{
		tr: tr,
	}
}

// CreateTicket godoc
//
//	@Summary		Create a new ticket
//	@Description	Create a new ticket
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			ticket	body		model.TicketCreate	true	"Ticket"
//	@Success		201		{object}	string				"Ticket id"
//	@Failure		400		{object}	model.ErrorResponse	"Bad Request"
//	@Failure		422		{object}	model.ErrorResponse	"Unprocessable Entity"
//	@Router			/v1/tickets [post]
func (tc *ticketController) CreateTicket(c *gin.Context) {
	ticket := model.TicketCreate{}
	if err := c.BindJSON(&ticket); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{Message: err.Error()})
		return
	}

	ticketsCount, err := tc.tr.Count(c.Request.Context(), bson.M{"departmentId": ticket.DepartmentId, "timeSlot": ticket.TimeSlot})
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	if ticketsCount >= config.Cfg.DepartmentCapacityPerTimeSlot {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "department is full"})
		return
	}

	var userId string

	userId, err = c.Cookie("session")
	if err != nil {
		userId = c.GetString("session")
	}

	ticket.UserId = userId
	ticketId, err := tc.tr.InsertOne(c.Request.Context(), ticket)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ticketId)
}

// GetTickets godoc
//
//	@Summary		Get ticket by id
//	@Description	Get ticket by id
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			ticketId	path		string			true	"Ticket id"
//	@Success		200			{object}	model.Ticket	"Ticket"
//	@Failure		404			{object}	string			"Not Found"
//	@Router			/v1/tickets/{ticketId} [get]
func (tc *ticketController) GetTicketById(c *gin.Context) {
	ticketId := c.Param("ticketId")

	ticket, err := tc.tr.FindOne(c.Request.Context(), ticketId)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// GetTicketsForDepartment godoc
//
//	@Summary		Get all tickets for department
//	@Description	Get all tickets for department
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			departmentId	path		string			true	"Department id"
//	@Success		200				{array}		model.Ticket	"Tickets"
//	@Failure		404				{object}	string			"Not Found"
//	@Router			/v1/tickets/department/{departmentId} [get]
func (tc *ticketController) GetTicketsForDepartment(c *gin.Context) {
	departmentId := c.Param("departmentId")

	tickets, err := tc.tr.FindMany(c.Request.Context(), bson.M{"departmentId": departmentId})
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// GetTicketsForUser godoc
//
//	@Summary		Get all tickets for user
//	@Description	Get all tickets for user
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Ticket	"Tickets"
//	@Failure		404	{object}	string			"Not Found"
//	@Router			/v1/tickets/user [get]
func (tc *ticketController) GetTicketsForUser(c *gin.Context) {
	var userId string
	var err error

	userId, err = c.Cookie("session")
	if err != nil {
		userId = c.GetString("session")
	}

	tickets, err := tc.tr.FindMany(c.Request.Context(), bson.M{"userId": userId})
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// CancelTicket godoc
//
//	@Summary		Cancel ticket
//	@Description	Cancel ticket
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			ticketId	path		string	true	"Ticket id"
//	@Success		204			{object}	string	"No Content"
//	@Failure		404			{object}	string	"Not Found"
//	@Router			/v1/tickets/{ticketId} [delete]
func (tc *ticketController) CancelTicket(c *gin.Context) {
	ticketId := c.Param("ticketId")
	logging.Log.Debugf("ticketId: %s", ticketId)

	err := tc.tr.DeleteOne(c.Request.Context(), ticketId)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
