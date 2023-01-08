package repositories

import (
	"context"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/env"
)

type Set struct {
	BigQuery *BigQuery
	MoodleDB *MoodleDB
}

func NewSet(ctx context.Context, conf config.Config) (*Set, error) {
	var set Set

	bq, err := NewBigQuery(ctx, conf.GetString(env.BigQueryServiceAccount), conf)
	if err != nil {
		return nil, err
	}

	mdb, err := NewMoodleDB(ctx, conf)
	if err != nil {
		return nil, err
	}

	set.BigQuery = bq
	set.MoodleDB = mdb
	return &set, nil
}
