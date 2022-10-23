package services

import (
	"context"

	"github.com/Reskill-2022/hoarder/models"
	"github.com/Reskill-2022/hoarder/repositories"
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

func NewSlackService() *SlackService {
	return &SlackService{}
}
