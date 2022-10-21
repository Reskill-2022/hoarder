package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/log"
	"github.com/Reskill-2022/hoarder/requests"
)

const (
	EventTypeURLVerification = "url_verification"
)

type SlackController struct{}

// AuthorizationChallenge is the handler for the Slack authorization challenge.
// It responds '200 OK' with the challenge string found in the request body.
func (s *SlackController) AuthorizationChallenge() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var requestBody requests.SlackChallengeRequest
		if err := c.Bind(&requestBody); err != nil {
			return ErrorHandler(c, err)
		}

		log.FromContext(ctx).Named("SlackController.AuthorizationChallenge").Info("received challenge: " + requestBody.Challenge)

		return c.String(http.StatusOK, requestBody.Challenge)
	}
}

// Events is the multiplexer for all Slack events.
// It routes each event to the appropriate handler based on the event type.
// Events is not a middleware in the traditional sense, but it is a handler that routes.
func (s *SlackController) Events() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		// because binding exhausts the request body, we need to copy it
		// so that we can use it again later
		bodyCpy, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return ErrorHandler(c, err)
		}

		var event requests.SlackEvent
		if err := json.Unmarshal(bodyCpy, &event); err != nil {
			return ErrorHandler(c, errors.From(err, "malformed request body.", 400))
		}

		c.Request().Body = ioutil.NopCloser(bytes.NewReader(bodyCpy))

		// log event type
		log.FromContext(ctx).Named("SlackController.Events").Info("received event type: " + event.EventType)

		// multiplex to appropriate handler
		next := s.getEventHandler(event.EventType)
		if next == nil {
			return ErrorHandler(c, errors.New("no handler for event type", 500))
		}
		return next(c)
	}
}

// getEventHandler returns the appropriate handler for the given event type.
func (s *SlackController) getEventHandler(eventType string) echo.HandlerFunc {
	switch strings.ToLower(eventType) {
	case EventTypeURLVerification:
		return s.AuthorizationChallenge()
	default:
		return nil
	}
}

func NewSlackController() *SlackController {
	return &SlackController{}
}
