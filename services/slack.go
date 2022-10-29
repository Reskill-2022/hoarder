package services

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/env"
	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/models"
	"github.com/Reskill-2022/hoarder/repositories"
)

var (
	ChannelBlacklist = []string{
		"C047AC49K0F",
	}
)

const (
	EventTypeMessage = "message"
)

type SlackService struct {
	conf config.Config
}

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

	TicketMessageInput struct {
		ChannelID    string
		MarkdownText string
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

	for _, channelID := range ChannelBlacklist {
		if input.ChannelID == channelID {
			// ignore messages from blacklisted channels
			return nil
		}
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

func (s *SlackService) SendTicketMessage(ctx context.Context, input TicketMessageInput) error {
	if input.ChannelID == "" {
		return errors.New("channel id is required", 400)
	}
	if input.MarkdownText == "" {
		return errors.New("markdown text is required", 400)
	}

	payload := map[string]interface{}{
		"channel": input.ChannelID,
		"blocks": []map[string]interface{}{
			{
				"type": "section",
				"text": map[string]interface{}{
					"type": "mrkdwn",
					"text": input.MarkdownText,
				},
			},
		},
	}

	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.postMessage", JSONPayloadReader(payload))
	req.Header.Set(echo.HeaderContentType, "application/json")
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+s.conf.GetString(env.SlackToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.From(err, "failed to send message to slack", 500)
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("got non-200 response from slack", resp.StatusCode)
	}

	return nil
}

func (s *SlackService) SendMessage(ctx context.Context, input SendMessageInput) error {
	if input.ChannelID == "" {
		return errors.New("channel id is required", 400)
	}
	if input.Text == "" {
		return errors.New("text is required", 400)
	}

	payload := map[string]interface{}{
		"channel": input.ChannelID,
		"text":    input.Text,
	}
	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.postMessage", JSONPayloadReader(payload))
	req.Header.Set(echo.HeaderContentType, "application/json")
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+s.conf.GetString(env.SlackToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.From(err, "failed to send message to slack", 500)
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("got non-200 response from slack", resp.StatusCode)
	}

	return nil
}

func NewSlackService(conf config.Config) *SlackService {
	return &SlackService{
		conf: conf,
	}
}
