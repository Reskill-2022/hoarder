package main

import (
	"context"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/controllers"
	"github.com/Reskill-2022/hoarder/env"
	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/jobs"
	"github.com/Reskill-2022/hoarder/log"
	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/server"
	"github.com/Reskill-2022/hoarder/services"
)

func main() {
	ctx := context.Background()

	conf := config.New()
	conf.AddFromProvider(environment(ctx))

	ctx = log.WithContext(ctx, log.New(conf.GetString(env.ServiceLogLevel)))

	svs := services.NewSet(conf)
	cts := controllers.NewSet(svs)
	jbs := jobs.NewSet(conf)

	rcs, err := repositories.NewSet(ctx, conf)
	if err != nil {
		log.FromContext(ctx).Named("main").Fatal("failed to create repositories set", errors.ErrorLogFields(err)...)
	}

	if err := server.ScheduleJobs(ctx, jbs, svs, rcs); err != nil { // async
		log.FromContext(ctx).Named("main").Fatal("failed to schedule cron jobs", errors.ErrorLogFields(err)...)
	}

	if err := server.Start(ctx, cts, svs, rcs, conf.GetString(env.ServerPort)); err != nil { // blocking
		log.FromContext(ctx).Named("main").Fatal("failed to start HTTP server", errors.ErrorLogFields(err)...)
	}
}

func environment(ctx context.Context) config.Provider {
	return config.NewStaticProvider(map[string]interface{}{
		env.ServiceLogLevel:           config.GetEnv(env.ServiceLogLevel, "INFO"),
		env.ServerPort:                config.GetEnv(env.ServerPort, "8001"),
		env.BigQueryServiceAccount:    config.GetBase64EncodedEnv(env.BigQueryServiceAccount, ""),
		env.BigQuerySlackDatasetID:    config.MustGetEnv(ctx, env.BigQuerySlackDatasetID),
		env.BigQuerySlackTableID:      config.MustGetEnv(ctx, env.BigQuerySlackTableID),
		env.BigQueryZendeskDatasetID:  config.MustGetEnv(ctx, env.BigQueryZendeskDatasetID),
		env.BigQueryZendeskTableID:    config.MustGetEnv(ctx, env.BigQueryZendeskTableID),
		env.BigQueryCalendlyDatasetID: config.MustGetEnv(ctx, env.BigQueryCalendlyDatasetID),
		env.BigQueryCalendlyTableID:   config.MustGetEnv(ctx, env.BigQueryCalendlyTableID),
		env.BigQueryMoodleDatasetID:   config.MustGetEnv(ctx, env.BigQueryMoodleDatasetID),
		env.BigQueryMoodleLogsTableID: config.MustGetEnv(ctx, env.BigQueryMoodleLogsTableID),
		env.BigQueryProjectID:         config.MustGetEnv(ctx, env.BigQueryProjectID),
		env.SlackToken:                config.MustGetEnv(ctx, env.SlackToken),
		env.CalendlyMember1Token:      config.MustGetEnv(ctx, env.CalendlyMember1Token),
		env.CalendlyMember2Token:      config.MustGetEnv(ctx, env.CalendlyMember2Token),
		env.MoodleDBUser:              config.MustGetEnv(ctx, env.MoodleDBUser),
		env.MoodleDBPassword:          config.MustGetEnv(ctx, env.MoodleDBPassword),
		env.MoodleDBHost:              config.MustGetEnv(ctx, env.MoodleDBHost),
		env.MoodleDBPort:              config.MustGetEnv(ctx, env.MoodleDBPort),
		env.MoodleDBName:              config.MustGetEnv(ctx, env.MoodleDBName),
	})
}
