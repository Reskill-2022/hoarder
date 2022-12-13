package controllers

import (
	"net/http"

	"github.com/jcobhams/echoresponse"
	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/log"
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

		var request requests.CalendlyEventRequest
		if err := c.Bind(&request); err != nil {
			log.FromContext(ctx).Warn("malformed request body", errors.ErrorLogFields(err)...)
			return echoresponse.Format(c, "malformed request body", nil, http.StatusBadRequest)
		}

		event, err := service.ResolveScheduledEvent(ctx, c.Param("memberId"), request.Event)
		if err != nil {
			log.FromContext(ctx).Warn("failed to resolved scheduled event", errors.ErrorLogFields(err)...)
			return echoresponse.Format(c, "failed to resolved scheduled event", nil, http.StatusInternalServerError)
		}

		return echoresponse.Format(c, "OK", nil, http.StatusOK)
	}
}

func NewCalendlyController(service services.CalendlyServiceInterface) *CalendlyController {
	return &CalendlyController{
		service: service,
	}
}
