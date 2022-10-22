package services

import (
	"context"

	"github.com/Reskill-2022/hoarder/repositories"
)

type (
	SlackServiceInterface interface {
		EventOccurred(ctx context.Context, input EventInput, creator repositories.SlackMessageCreator) error
	}
)
