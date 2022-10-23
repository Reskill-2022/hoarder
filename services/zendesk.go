package services

import (
	"context"
	"time"

	"github.com/Reskill-2022/hoarder/models"
	"github.com/Reskill-2022/hoarder/repositories"
)

type ZendeskService struct{}

type (
	CreateTicketInput struct {
		ID            int
		TicketType    string
		Title         string
		Description   string
		Link          string
		Via           string
		Status        string
		Priority      string
		LatestComment string
		Requester     string
		Satisfaction  string
		Assignee      string
		RequestedAt   time.Time
	}
)

func (z *ZendeskService) CreateTicket(ctx context.Context, input CreateTicketInput, creator repositories.ZendeskTicketCreator) error {

	ticket := models.ZendeskTicket{
		ID:            input.ID,
		TicketType:    input.TicketType,
		Subject:       input.Title,
		Description:   input.Description,
		Link:          input.Link,
		Via:           input.Via,
		Status:        input.Status,
		Priority:      input.Priority,
		LatestComment: input.LatestComment,
		Requester:     input.Requester,
		Satisfaction:  input.Satisfaction,
		Assignee:      input.Assignee,
		RequestedAt:   input.RequestedAt,
	}
	return creator.CreateZendeskTicket(ctx, ticket)
}

func NewZendeskService() *ZendeskService {
	return &ZendeskService{}
}
