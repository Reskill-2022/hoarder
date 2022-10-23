package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/requests"
	"github.com/Reskill-2022/hoarder/services"
	"github.com/jcobhams/echoresponse"
	"github.com/labstack/echo/v4"
)

type ZendeskController struct {
	service services.ZendeskServiceInterface
}

func (z *ZendeskController) CreateTicket(ticketCreator repositories.ZendeskTicketCreator) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var requestBody requests.ZendeskTicketCreateRequest
		if err := c.Bind(&requestBody); err != nil {
			return echoresponse.Format(c, "malformed request body", nil, http.StatusBadRequest)
		}

		ticketID, err := strconv.Atoi(requestBody.ID)
		if err != nil {
			return echoresponse.Format(c, "failed to convert ticket id to int", nil, http.StatusBadRequest)
		}

		requestedAt, err := time.Parse(time.RFC3339, requestBody.Requested)
		if err != nil {
			return echoresponse.Format(c, "failed to parse requested time", nil, http.StatusBadRequest)
		}

		createInput := services.CreateTicketInput{
			ID:            ticketID,
			TicketType:    requestBody.TicketType,
			Title:         requestBody.Title,
			Description:   requestBody.DescriptionPlain,
			Link:          requestBody.Link,
			Via:           requestBody.Via,
			Status:        requestBody.Status,
			Priority:      requestBody.Priority,
			LatestComment: requestBody.LatestComment,
			Requester:     requestBody.RequesterEmail,
			Satisfaction:  requestBody.Satisfaction,
			Assignee:      requestBody.Assignee,
			RequestedAt:   requestedAt,
		}
		if err := z.service.CreateTicket(ctx, createInput, ticketCreator); err != nil {
			return ErrorHandler(c, err)
		}

		return echoresponse.Format(c, "OK", nil, http.StatusOK)
	}
}

func NewZendeskController(service services.ZendeskServiceInterface) *ZendeskController {
	return &ZendeskController{
		service: service,
	}
}
