package repositories

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/env"
	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/models"
)

var _ MoodleRepositoryInterface = (*MoodleDB)(nil) // compile-time check to ensure that MoodleDB fully implements the MoodleRepositoryInterface.

type MoodleDB struct {
	conf   config.Config
	client *gorm.DB
}

func NewMoodleDB(ctx context.Context, conf config.Config) (*MoodleDB, error) {
	gormCfg := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	}
	db, err := gorm.Open(mysql.Open(buildDSN(conf)), gormCfg)
	if err != nil {
		return nil, errors.From(err, "failed to connect to database", 500)
	}

	return &MoodleDB{
		conf:   conf,
		client: db,
	}, nil
}

func (mdb *MoodleDB) ListLogs(ctx context.Context, since *time.Time) ([]*models.MoodleLogLine, error) {
	if since == nil {
		return nil, errors.New("fetching all logs is not supported", 400)
	}

	var logLines []*models.MoodleLogLine

	tx := mdb.client.WithContext(ctx).Model(&models.MoodleLogLine{})
	tx.Where("timecreated > ?", since.Unix()).Find(&logLines)
	if tx.Error != nil {
		return nil, errors.From(tx.Error, "failed to fetch logs", 500)
	}

	return logLines, nil
}

func buildDSN(conf config.Config) string {
	dsnFmt := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4"
	return fmt.Sprintf(dsnFmt,
		conf.GetString(env.MoodleDBUser),
		conf.GetString(env.MoodleDBPassword),
		conf.GetString(env.MoodleDBHost),
		conf.GetString(env.MoodleDBPort),
		conf.GetString(env.MoodleDBName),
	)
}
