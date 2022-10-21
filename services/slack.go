package services

import (
	"context"

	"github.com/Reskill-2022/hoarder/log"
	"go.uber.org/zap"
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

func (s *SlackService) ChannelMessage(ctx context.Context, input ChannelMessageInput) error {
	//todo: write to BigQuery

	// log the input
	log.FromContext(ctx).Named("SlackService.ChannelMessage").Info("input",
		zap.String("EventType", input.EventType),
		zap.String("Text", input.Text),
		zap.String("Timestamp", input.Timestamp),
		zap.String("ChannelID", input.ChannelID),
		zap.String("EventID", input.EventID),
		zap.String("TeamID", input.TeamID),
		zap.String("UserID", input.UserID),
		zap.String("ChannelType", input.ChannelType),
		zap.Int64("EventTimestamp", input.EventTimestamp),
	)

	return nil
}

func NewSlackService() *SlackService {
	return &SlackService{}
}
