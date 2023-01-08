package server

import (
	"context"

	"github.com/Reskill-2022/hoarder/cron"
	"github.com/Reskill-2022/hoarder/jobs"
	"github.com/Reskill-2022/hoarder/repositories"
	"github.com/Reskill-2022/hoarder/services"
)

// ScheduleJobs schedules all cron jobs. It is a non-blocking call.
func ScheduleJobs(ctx context.Context, jbs *jobs.Set, svs *services.Set, rcs *repositories.Set) error {
	scheduler := cron.New()

	scheduler.Schedule("@hourly", jbs.MoodleJobs.ExtractTransformLoadLogs(svs.MoodleService, rcs.MoodleDB))

	return scheduler.Start(ctx)
}
