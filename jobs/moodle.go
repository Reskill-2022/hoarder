package jobs

import (
	"context"
	"time"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/cron"
	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/services"
)

type MoodleJobs struct {
	conf config.Config
}

func (m *MoodleJobs) ExtractTransformLoadLogs(service services.MoodleServiceInterface, repo repositories.MoodleRepositoryInterface, logLineCreator repositories.MoodleLogLineCreator) cron.Job {
	return func(ctx context.Context) error {

		t := time.Now().UTC().Add(-10 * time.Second)
		//todo: read last timestamp

		logLines, err := service.ListLogs(ctx, &t, repo)
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
