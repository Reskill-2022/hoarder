package cron

import (
	"context"
	"sync"

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
func (s *Scheduler) Schedule(spec string, do Job) *Scheduler {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.specJobs = append(s.specJobs, specJob{spec: spec, do: do})
	return s
}

// Start starts the scheduler. It is a non-blocking call.
// It returns any error that occurred while scheduling the jobs.
func (s *Scheduler) Start(ctx context.Context) error {
	for _, specJob := range s.specJobs {
		cmd := func() {
			_ = specJob.do(ctx)
		}

		if _, err := s.cron.AddFunc(specJob.spec, cmd); err != nil {
			return err
		}
	}

	s.cron.Start()
	return nil
}
