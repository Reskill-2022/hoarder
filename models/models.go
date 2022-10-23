package models

import (
	"fmt"
	"time"
)

type (
	SlackMessage struct {
		EventID     string `json:"event_id" bigquery:"event_id"`
		EventType   string `json:"event_type" bigquery:"event_type"`
		Text        string `json:"text" bigquery:"text"`
		UserID      string `json:"user_id" bigquery:"user_id"`
		ChannelID   string `json:"channel_id" bigquery:"channel_id"`
		ChannelType string `json:"channel_type" bigquery:"channel_type"`
		TeamID      string `json:"team_id" bigquery:"team_id"`
		Timestamp   string `json:"timestamp" bigquery:"timestamp"`
		EventTime   int64  `json:"event_time" bigquery:"event_time"`
	}

	ZendeskTicket struct {
		ID            int       `json:"id" bigquery:"id"`
		Status        string    `json:"status" bigquery:"status"`
		Satisfaction  string    `json:"satisfaction" bigquery:"satisfaction"`
		Subject       string    `json:"subject" bigquery:"subject"`
		Requester     string    `json:"requester" bigquery:"requester"`
		RequestedAt   time.Time `json:"requested_at" bigquery:"requested"`
		Assignee      string    `json:"assignee" bigquery:"assignee"`
		TicketType    string    `json:"type" bigquery:"type"`
		Description   string    `json:"description" bigquery:"description"`
		Link          string    `json:"link" bigquery:"link"`
		Via           string    `json:"via" bigquery:"via"`
		Priority      string    `json:"priority" bigquery:"priority"`
		LatestComment string    `json:"latest_comment" bigquery:"latest_comment"`
	}
)

func (z ZendeskTicket) String() string {
	return fmt.Sprintf(`
New Zendesk Ticket #%d

%s

%s

Requested By: %s

%s
	`, z.ID, z.Subject, z.Description, z.Requester, z.Link)
}
