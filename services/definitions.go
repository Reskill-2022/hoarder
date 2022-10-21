package services

import "context"

type (
	SlackServiceInterface interface {
		ChannelMessage(ctx context.Context, input ChannelMessageInput) error
	}
)
