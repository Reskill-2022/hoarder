package services

import (
	"context"
	"time"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/models"
	"github.com/Reskill-2022/hoarder/repositories"
)

type (
	MoodleService struct {
		conf config.Config
	}
)

func (m *MoodleService) ListLogs(ctx context.Context, since *time.Time, lister repositories.MoodleRepositoryInterface) ([]*models.MoodleLogLine, error) {
	if since == nil {
		return nil, errors.New("fetching all logs is not supported", 400)
	}
	return lister.ListLogs(ctx, since)
}

func NewMoodleService(conf config.Config) *MoodleService {
	return &MoodleService{
		conf: conf,
	}
}
