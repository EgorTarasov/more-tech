package controller

import (
	"errors"
	"more-tech/internal/logging"
	"more-tech/internal/model"
	"more-tech/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ticketController struct {
	tr        model.TicketRepository
	dr        model.DepartmentRepository
	routeHost string
}

func NewTicketController(tr model.TicketRepository, dr model.DepartmentRepository, routeHost string) *ticketController {
	return &ticketController{
		tr:        tr,
		dr:        dr,
		routeHost: routeHost,
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

	if ticket.Duration == 0 {
		ticket.Duration = 15 * 60
	}
	if ticket.Description == "" {
		ticket.Description = "Открытие вклада"
	}

	hexId, err := primitive.ObjectIDFromHex(ticket.DepartmentId)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	department, err := tc.dr.FindOne(c.Request.Context(), bson.M{"_id": hexId})
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "department not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	timeCar, timeWalk, err := service.GetEstimatedTime(ticket.StartLongitude, ticket.StartLatitude, department.Coordinates.Longitude, department.Coordinates.Latitude, tc.routeHost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var minTime float64
	if timeCar < timeWalk {
		minTime = timeCar
	} else {
		minTime = timeWalk
	}

	if service.GetClosestTimeSlot(minTime, nil, ticket.TimeSlot) != ticket.TimeSlot {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "choose later timeslot"})
		return
	}

	tickets, err := tc.tr.FindMany(c.Request.Context(), bson.M{"departmentId": ticket.DepartmentId, "timeSlot": ticket.TimeSlot})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	logging.Log.Debugf("tickets on this timeslot: %+v", tickets)

	ticketsTime := 0.0
	for _, ticket := range tickets {
		ticketsTime += ticket.Duration
	}
	timeCar += ticketsTime
	timeWalk += ticketsTime

	if ticketsTime >= 60*60 {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "choose later timeslot"})
		return
	}

	_, err = tc.tr.Count(c.Request.Context(), bson.M{"departmentId": ticket.DepartmentId, "timeSlot": ticket.TimeSlot})
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
		return
	}

	// if ticketsCount >= config.Cfg.DepartmentCapacityPerTimeSlot {
	// 	c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "department is full"})
	// 	return
	// }

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

	c.JSON(http.StatusCreated, gin.H{
		"ticketId":          ticketId,
		"estimatedTimeCar":  timeCar + ticket.Duration,
		"estimatedTimeWalk": timeWalk + ticket.Duration,
	})
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
