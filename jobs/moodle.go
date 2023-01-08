package jobs

import (
	"context"
	"time"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/cron"
	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/services"
)

type MoodleJobs struct {
	conf config.Config
}

func (m *MoodleJobs) ExtractTransformLoadLogs(service services.MoodleServiceInterface,
	repo repositories.MoodleRepositoryInterface,
	logLineCreator repositories.MoodleLogLineCreator,
	logLineGetter repositories.MoodleLogLineGetter,
) cron.Job {
	return func(ctx context.Context) error {

		last, err := service.GetLatestLog(ctx, logLineGetter)
		if err != nil {
			return errors.From(err, "failed to get latest log line", 500)
		}

		lastFetchTime := time.Unix(last.TimeCreated, 0)

		logLines, err := service.ListLogs(ctx, &lastFetchTime, repo)
		if err != nil {
			return err
		}

		for _, line := range logLines {
			err := service.CreateLogLine(ctx, line, logLineCreator)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func NewMoodleJobs(conf config.Config) *MoodleJobs {
	return &MoodleJobs{
		conf: conf,
	}
}
