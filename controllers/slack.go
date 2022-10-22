package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jcobhams/echoresponse"
	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/log"
	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/requests"
	"github.com/Reskill-2022/hoarder/services"
)

const (
	EventTypeURLVerification = "url_verification"
	EventTypeEventCallback   = "event_callback"

	EventCallbackTypeMessage = "message"
)

type SlackController struct {
	service services.SlackServiceInterface
}

func (s *SlackController) Message(creator repositories.SlackMessageCreator) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var requestBody requests.SlackEventCallback
		if err := c.Bind(&requestBody); err != nil {
			return ErrorHandler(c, err)
		}

		if requestBody.Event.Type != EventCallbackTypeMessage {
			// can't handle this, not a failure
			return echoresponse.Format(c, "OK", nil, http.StatusOK)
		}

		messageInput := services.ChannelMessageInput{
			EventType:      requestBody.Event.Type,
			Text:           requestBody.Event.Text,
			Timestamp:      requestBody.Event.Timestamp,
			ChannelID:      requestBody.Event.Channel,
			EventID:        requestBody.EventID,
			TeamID:         requestBody.TeamID,
			UserID:         requestBody.Event.User,
			ChannelType:    requestBody.Event.ChannelType,
			EventTimestamp: requestBody.EventTime,
		}
		if err := s.service.ChannelMessage(ctx, messageInput, creator); err != nil {
			return ErrorHandler(c, err)
		}

		return echoresponse.Format(c, "OK", nil, http.StatusOK)
	}
}

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
func (s *SlackController) Events(slackMessageCreator repositories.SlackMessageCreator) echo.HandlerFunc {
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
		next := s.getEventHandler(event.EventType, slackMessageCreator)
		if next == nil {
			return ErrorHandler(c, errors.New("no handler for event type", 500))
		}
		return next(c)
	}
}

// getEventHandler returns the appropriate handler for the given event type.
func (s *SlackController) getEventHandler(eventType string,
	slackMessageCreator repositories.SlackMessageCreator,
) echo.HandlerFunc {

	switch strings.ToLower(eventType) {
	case EventTypeURLVerification:
		return s.AuthorizationChallenge()

	case EventTypeEventCallback:
		return s.Message(slackMessageCreator)

	default:
		return nil
	}
}

func NewSlackController(service services.SlackServiceInterface) *SlackController {
	return &SlackController{
		service: service,
	}
}
