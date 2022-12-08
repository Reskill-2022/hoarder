package controllers

import (
	"fmt"
	"net/http"

	"github.com/jcobhams/echoresponse"
	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/requests"
	"github.com/Reskill-2022/hoarder/services"
)

type CalendlyController struct {
	service services.CalendlyServiceInterface
}

func (c *CalendlyController) Events(creator repositories.CalendlyEventCreator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request requests.CalendlyEventRequest
		if err := c.Bind(&request); err != nil {
			return echoresponse.Format(c, "malformed request body", nil, http.StatusBadRequest)
		}

		fmt.Printf("request: %+v", request)
		return echoresponse.Format(c, "OK", nil, http.StatusOK)
	}
}

func NewCalendlyController(service services.CalendlyServiceInterface) *CalendlyController {
	return &CalendlyController{
		service: service,
	}
}
