package repositories

import (
	"context"
	"time"

	"github.com/Reskill-2022/hoarder/models"
)

type (
	SlackMessageCreator interface {
		CreateSlackMessage(ctx context.Context, message models.SlackMessage) error
	}

	ZendeskTicketCreator interface {
		CreateZendeskTicket(ctx context.Context, ticket models.ZendeskTicket) error
	}

	CalendlyEventCreator interface {
		CreateCalendlyEvent(ctx context.Context, event models.CalendlyEvent) error
	}

	MoodleLogLineCreator interface {
		CreateMoodleLogLine(ctx context.Context, logLine models.MoodleLogLine) error
	}

	MoodleRepositoryInterface interface {
		ListLogs(ctx context.Context, since *time.Time) ([]*models.MoodleLogLine, error)
	}
)
