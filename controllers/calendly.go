package controllers

import (
	"net/http"

	"github.com/jcobhams/echoresponse"
	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/requests"
	"github.com/Reskill-2022/hoarder/services"
)

type CalendlyController struct {
	service services.CalendlyServiceInterface
}

func (c *CalendlyController) Events(service services.CalendlyServiceInterface, creator repositories.CalendlyEventCreator) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var requestBody requests.CalendlyEventRequest
		if err := c.Bind(&requestBody); err != nil {
			return ErrorHandler(c, errors.New("malformed request body", http.StatusBadRequest))
		}

		scheduledEvent, err := service.ResolveScheduledEvent(ctx, c.Param("memberId"), requestBody.Payload.Event)
		if err != nil {
			return ErrorHandler(c, err)
		}

		input := services.CalendlyEventInput{
			EventKind:    requestBody.EventKind,
			InviteeEmail: requestBody.Payload.Email,
			InviteeName:  requestBody.Payload.Name,
			CreatedBy:    requestBody.CreatedBy,
			EventURI:     scheduledEvent.Resource.URI,
			EventName:    scheduledEvent.Resource.Name,
			CreatedAt:    scheduledEvent.Resource.CreatedAt,
			UpdatedAt:    scheduledEvent.Resource.UpdatedAt,
			StartTime:    scheduledEvent.Resource.StartTime,
			EndTime:      scheduledEvent.Resource.EndTime,
		}
		err = service.EventOccurred(ctx, input, creator)
		if err != nil {
			return ErrorHandler(c, err)
		}

		return echoresponse.Format(c, "OK", nil, http.StatusOK)
	}
}

func NewCalendlyController(service services.CalendlyServiceInterface) *CalendlyController {
	return &CalendlyController{
		service: service,
	}
}
