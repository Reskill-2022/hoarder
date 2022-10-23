package services

import (
	"context"

	"github.com/Reskill-2022/hoarder/repositories"
)

type (
	SlackMessageSender interface {
		SendMessage(ctx context.Context, input SendMessageInput) error
	}

	SlackServiceInterface interface {
		EventOccurred(ctx context.Context, input EventInput, creator repositories.SlackMessageCreator) error
		SlackMessageSender
	}

	ZendeskServiceInterface interface {
		CreateTicket(ctx context.Context, input CreateTicketInput, creator repositories.ZendeskTicketCreator) error
	}
)
