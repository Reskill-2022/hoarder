package jobs

import (
	"context"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/cron"
	"github.com/Reskill-2022/hoarder/errors"
)

type MoodleJobs struct {
	conf config.Config
}

func (m *MoodleJobs) ExtractTransformLoadLogs() cron.Job {
	return func(ctx context.Context) error {
		return errors.New("not implemented", 500)
	}
}

func NewMoodleJobs(conf config.Config) *MoodleJobs {
	return &MoodleJobs{
		conf: conf,
	}
}
