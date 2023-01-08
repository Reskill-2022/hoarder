package services

import "github.com/Reskill-2022/hoarder/config"

type (
	MoodleService struct {
		conf config.Config
	}
)

func NewMoodleService(conf config.Config) *MoodleService {
	return &MoodleService{
		conf: conf,
	}
}
