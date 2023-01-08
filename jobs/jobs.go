package jobs

import "github.com/Reskill-2022/hoarder/config"

type Set struct {
	MoodleJobs *MoodleJobs
}

func NewSet(conf config.Config) *Set {
	return &Set{
		MoodleJobs: NewMoodleJobs(conf),
	}
}
