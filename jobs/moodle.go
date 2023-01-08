package jobs

import "github.com/Reskill-2022/hoarder/config"

type MoodleJobs struct {
	conf config.Config
}

func NewMoodleJobs(conf config.Config) *MoodleJobs {
	return &MoodleJobs{
		conf: conf,
	}
}
