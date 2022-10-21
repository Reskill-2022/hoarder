package controllers

import "github.com/labstack/echo/v4"

type SlackController struct{}

// Events is the multiplexer for all Slack events.
// It routes each event to the appropriate handler based on the event type.
func (s *SlackController) Events() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func NewSlackController() *SlackController {
	return &SlackController{}
}
