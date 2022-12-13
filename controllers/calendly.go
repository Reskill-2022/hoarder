package controllers

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/jcobhams/echoresponse"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Reskill-2022/hoarder/log"
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

		b, err := httputil.DumpRequest(c.Request(), true)
		if err != nil {
			log.FromContext(c.Request().Context()).Error("error dumping request")
		}
		log.FromContext(c.Request().Context()).Warn("request", zap.ByteString("request", b))

		// log.FromContext(c.Request().Context()).Warn("failed to dump",
		// 	zap.Field{Key: "request", Type: zapcore.StringType, String: fmt.Sprintf("%+v", request)})

		if err := c.Bind(&request); err != nil {
			return echoresponse.Format(c, "malformed request body", nil, http.StatusBadRequest)
		}

		log.FromContext(c.Request().Context()).Warn("parsed request",
			zap.Field{Key: "request", Type: zapcore.StringType, String: fmt.Sprintf("%+v", request)})

		return echoresponse.Format(c, "OK", nil, http.StatusOK)
	}
}

func NewCalendlyController(service services.CalendlyServiceInterface) *CalendlyController {
	return &CalendlyController{
		service: service,
	}
}
