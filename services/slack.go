package services

import (
	"context"

	"github.com/Reskill-2022/hoarder/models"
	"github.com/Reskill-2022/hoarder/repositories"
)

type SlackService struct{}

type (
	ChannelMessageInput struct {
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

func (s *SlackService) ChannelMessage(ctx context.Context, input ChannelMessageInput, creator repositories.SlackMessageCreator) error {
	if input.Text == "" {
		return nil
	}

	slackMessage := models.SlackMessage{
		EventID:        input.EventID,
		EventType:      input.EventType,
		Text:           input.Text,
		UserID:         input.UserID,
		ChannelID:      input.ChannelID,
		ChannelType:    input.ChannelType,
		TeamID:         input.TeamID,
		Timestamp:      input.Timestamp,
		EventTimestamp: input.EventTimestamp,
	}
	return creator.CreateSlackMessage(ctx, slackMessage)
}

func NewSlackService() *SlackService {
	return &SlackService{}
}
