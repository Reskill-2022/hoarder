package repositories

import (
	"context"

	"github.com/Reskill-2022/hoarder/models"
)

type (
	SlackMessageCreator interface {
		CreateSlackMessage(ctx context.Context, message models.SlackMessage) error
	}
)
