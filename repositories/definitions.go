package repositories

import (
	"context"

	"github.com/Reskill-2022/hoarder/models"
)

type (
	SlackMessageCreator interface {
		CreateSlackMessage(ctx context.Context, message models.SlackMessage) error
	}

	ZendeskTicketCreator interface {
		CreateZendeskTicket(ctx context.Context, ticket models.ZendeskTicket) error
	}

	CalendlyEventCreator interface{}
)
