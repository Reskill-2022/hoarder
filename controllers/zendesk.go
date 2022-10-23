package controllers

import (
	"net/http"
	"net/http/httputil"

	"github.com/jcobhams/echoresponse"
	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/log"
)

type ZendeskController struct{}

func (z *ZendeskController) CreateTicket(x interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		b, _ := httputil.DumpRequest(c.Request(), true)
		log.FromContext(ctx).Named("zendesk.createTicket").Info(string(b))

		return echoresponse.Format(c, "OK", nil, http.StatusOK)
	}
}

func NewZendeskController() *ZendeskController {
	return &ZendeskController{}
}
