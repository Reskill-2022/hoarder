package services

import (
	"context"
	"time"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/repositories"
)

type (
	MoodleService struct {
		conf config.Config
	}

	LogsETLInput struct {
		StartTime *time.Time //todo: find a better name
	}
)

func (m *MoodleService) ExtractTransformLoadLogs(ctx context.Context, input LogsETLInput, repo repositories.MoodleRepositoryInterface) error {
	if input.StartTime == nil {
		return errors.New("start time is required", 400)
	}

	logs, err := repo.ListLogs(ctx, input.StartTime)
	if err != nil {
		return err
	}

	// print log IDs
	for _, log := range logs {
		println(log.ID)
	}

	return nil
}

func NewMoodleService(conf config.Config) *MoodleService {
	return &MoodleService{
		conf: conf,
	}
}
