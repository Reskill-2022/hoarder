package services

import (
	"context"
	"time"

	"github.com/Reskill-2022/hoarder/models"
	"github.com/Reskill-2022/hoarder/repositories"
)

type (
	SlackMessageSender interface {
		SendMessage(ctx context.Context, input SendMessageInput) error
		SendTicketMessage(ctx context.Context, input TicketMessageInput) error
	}

	SlackServiceInterface interface {
		EventOccurred(ctx context.Context, input SlackEventInput, creator repositories.SlackMessageCreator) error
		SlackMessageSender
	}

	ZendeskServiceInterface interface {
		CreateTicket(ctx context.Context, input CreateTicketInput, creator repositories.ZendeskTicketCreator) (*models.ZendeskTicket, error)
	}

	CalendlyServiceInterface interface {
		ResolveScheduledEvent(ctx context.Context, memberId, eventURI string) (*CalendlyScheduledEvent, error)
		EventOccurred(ctx context.Context, input CalendlyEventInput, creator repositories.CalendlyEventCreator) error
	}

	MoodleServiceInterface interface {
		ListLogs(ctx context.Context, since *time.Time, lister repositories.MoodleRepositoryInterface) ([]*models.MoodleLogLine, error)
		CreateLogLine(ctx context.Context, line *models.MoodleLogLine, creator repositories.MoodleLogLineCreator) error
		GetLatestLog(ctx context.Context, getter repositories.MoodleLogLineGetter) (*models.MoodleLogLine, error)
	}
)
