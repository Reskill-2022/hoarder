package services

import (
	"context"

	"github.com/Reskill-2022/hoarder/repositories"
)

type (
	SlackServiceInterface interface {
		ChannelMessage(ctx context.Context, input ChannelMessageInput, creator repositories.SlackMessageCreator) error
	}
)
