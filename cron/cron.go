package cron

import (
	"context"
	"fmt"
	"sync"

	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/log"
	"github.com/robfig/cron/v3"
)

type (
	Job func(ctx context.Context) error

	Scheduler struct {
		cron *cron.Cron

		mu       sync.Mutex
		specJobs []specJob
	}

	specJob struct {
		name string
		spec string
		do   Job
	}
)

// New creates a new Job Scheduler.
func New() *Scheduler {
	s := &Scheduler{
		cron: cron.New(),
	}
	return s
}

// Schedule registers a job to run as described by the cron spec.
// Specs can be @hourly, @daily, etc.
// See https://pkg.go.dev/github.com/robfig/cron#hdr-Predefined_schedules for full Spec list.
// It is safe to call this method concurrently.
func (s *Scheduler) Schedule(name, spec string, do Job) *Scheduler {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.specJobs = append(s.specJobs, specJob{name: name, spec: spec, do: do})
	return s
}

// Start starts the scheduler. It is a non-blocking call.
// It returns any error that occurred while scheduling the jobs.
func (s *Scheduler) Start(ctx context.Context) error {
	for _, specJob := range s.specJobs {
		cmd := func() {
			log.FromContext(ctx).Info(fmt.Sprintf("starting scheduled Job: '%s'", specJob.name))
			if err := specJob.do(ctx); err != nil {
				log.FromContext(ctx).Error("job failed with error", errors.ErrorLogFields(err)...)
			}
			log.FromContext(ctx).Info(fmt.Sprintf("finished scheduled Job: '%s'", specJob.name))
		}

		if _, err := s.cron.AddFunc(specJob.spec, cmd); err != nil {
			return err
		}
	}

	s.cron.Start()
	return nil
}
