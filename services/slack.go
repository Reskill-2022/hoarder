package services

import (
	"context"
	"net/http"

	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/models"
	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/labstack/echo/v4"
)

const (
	EventTypeMessage = "message"
)

type SlackService struct{}

type (
	EventInput struct {
		EventType      string
		Text           string
		Timestamp      string
		ChannelID      string
		EventID        string
		TeamID         string
		UserID         string
		ChannelType    string
		EventTimestamp int64
	}

	SendMessageInput struct {
		ChannelID string
		Text      string
	}
)

func (s *SlackService) EventOccurred(ctx context.Context, input EventInput, creator repositories.SlackMessageCreator) error {
	if input.Text == "" {
		return nil
	}

	if input.EventType != EventTypeMessage {
		// can only handle message events for now
		return nil
	}

	slackMessage := models.SlackMessage{
		EventID:     input.EventID,
		EventType:   input.EventType,
		Text:        input.Text,
		UserID:      input.UserID,
		ChannelID:   input.ChannelID,
		ChannelType: input.ChannelType,
		TeamID:      input.TeamID,
		Timestamp:   input.Timestamp,
		EventTime:   input.EventTimestamp,
	}
	return creator.CreateSlackMessage(ctx, slackMessage)
}

func (s *SlackService) SendMessage(ctx context.Context, input SendMessageInput) error {
	if input.ChannelID == "" {
		return errors.New("channel id is required", 400)
	}
	if input.Text == "" {
		return errors.New("text is required", 400)
	}

	endpoint := "https://slack.com/api/chat.postMessage"
	payload := map[string]interface{}{
		"channel": input.ChannelID,
		"text":    input.Text,
	}
	resp, err := http.Post(endpoint, echo.MIMEApplicationJSON, JSONPayloadReader(payload))
	if err != nil {
		return errors.From(err, "failed to send message to slack", 500)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("got non-200 response from slack", resp.StatusCode)
	}

	return nil
}

func NewSlackService() *SlackService {
	return &SlackService{}
}
