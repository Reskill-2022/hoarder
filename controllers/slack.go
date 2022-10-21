package controllers

import (
	"net/http"

	"github.com/jcobhams/echoresponse"
	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/log"
	"github.com/Reskill-2022/hoarder/requests"
)

type SlackController struct{}

// Events is the multiplexer for all Slack events.
// It routes each event to the appropriate handler based on the event type.
func (s *SlackController) Events() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var event requests.SlackEvent
		if err := c.Bind(&event); err != nil {
			return ErrorHandler(c, err)
		}

		// log event type
		log.FromContext(ctx).Named("SlackController.Events").Debug("received event type: " + event.EventType)

		// todo: multiplex to appropriate handler

		return echoresponse.Format(c, "OK", nil, http.StatusOK)
	}
}

func NewSlackController() *SlackController {
	return &SlackController{}
}
