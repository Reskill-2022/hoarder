package repositories

import (
	"context"
)

var _ MoodleRepositoryInterface = (*MoodleDB)(nil) // compile-time check to ensure that MoodleDB fully implements the MoodleRepositoryInterface.

type MoodleDB struct{}

func NewMoodleDB(ctx context.Context) (*MoodleDB, error) {
	return &MoodleDB{}, nil
}
