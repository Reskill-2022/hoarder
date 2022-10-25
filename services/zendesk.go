package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/log"
	"github.com/Reskill-2022/hoarder/models"
	"github.com/Reskill-2022/hoarder/repositories"
)

var blacklistRequesters = []string{
	"info@twitter.com",
	"noreply@mandrill.com",
	"feedback@slack.com",
	"account-insights@mailchimp.com",
	"survey@mailchimp.com",
	"jobs-listings@linkedin.com",
	"confirm@mailchimp.com",
	"security@mail.instagram.com",
	"notify@twitter.com",
	"notifications-noreply@linkedin.com",
}

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

func (z *ZendeskService) CreateTicket(ctx context.Context, input CreateTicketInput, creator repositories.ZendeskTicketCreator) (*models.ZendeskTicket, error) {
	for _, requester := range blacklistRequesters {
		if caselessEqual(requester, input.Requester) {
			log.FromContext(ctx).Named("zendesk.createTicket").Debug(fmt.Sprintf("requester '%s' is blacklisted", input.Requester))
			return nil, errors.New("requester is blacklisted", 400)
		}
	}

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
	if err := creator.CreateZendeskTicket(ctx, ticket); err != nil {
		return nil, err
	}

	return &ticket, nil
}

func NewZendeskService() *ZendeskService {
	return &ZendeskService{}
}
