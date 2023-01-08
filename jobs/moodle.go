package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/cron"
	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/services"
)

type MoodleJobs struct {
	conf config.Config
}

func (m *MoodleJobs) ExtractTransformLoadLogs(service services.MoodleServiceInterface, repo repositories.MoodleRepositoryInterface) cron.Job {
	return func(ctx context.Context) error {

		t := time.Now().UTC().Add(-10 * time.Second)
		//todo: read last timestamp

		logLines, err := service.ListLogs(ctx, &t, repo)
		if err != nil {
			return err
		}

		//todo: write to bigquery
		for _, line := range logLines {
			fmt.Println(line.ID)
		}

		return nil
	}
}

func NewMoodleJobs(conf config.Config) *MoodleJobs {
	return &MoodleJobs{
		conf: conf,
	}
}
