package models

import (
	"fmt"
	"strings"
	"time"
)

const TicketShortDescriptionLength = 1000

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
		CreatedAt     time.Time `json:"created_at" bigquery:"created_at"`
	}

	CalendlyEvent struct {
		Name         string    `json:"name" bigquery:"name"`
		Status       string    `json:"status" bigquery:"status"`
		EventURI     string    `json:"event_uri" bigquery:"event_uri"`
		EventKind    string    `json:"event_kind" bigquery:"event_kind"`
		InviteeEmail string    `json:"invitee_email" bigquery:"invitee_email"`
		InviteeName  string    `json:"invitee_name" bigquery:"invitee_name"`
		CreatedBy    string    `json:"created_by" bigquery:"created_by"`
		CreatedAt    time.Time `json:"created_at" bigquery:"created_at"`
		UpdatedAt    time.Time `json:"updated_at" bigquery:"updated_at"`
		StartTime    time.Time `json:"start_time" bigquery:"start_time"`
		EndTime      time.Time `json:"end_time" bigquery:"end_time"`
	}

	MoodleLogLine struct {
		ID              int    `json:"id" bigquery:"id"`
		EventName       string `json:"event_name" bigquery:"event_name" gorm:"column:eventname"`
		Component       string `json:"component" bigquery:"component" gorm:"column:component"` // core, mod_page, ...
		Action          string `json:"action" bigquery:"action" gorm:"column:action"`          // viewed, updated, ...
		Target          string `json:"target" bigquery:"target" gorm:"column:target"`          // course, course_module_completion, dashboard, ...
		ObjectTableName string `json:"object_table" bigquery:"object_table" gorm:"column:objecttable"`
		ObjectID        int    `json:"object_id" bigquery:"object_id" gorm:"column:objectid"`
		UserID          int    `json:"user_id" bigquery:"user_id" gorm:"column:userid"`
		CourseID        int    `json:"course_id" bigquery:"course_id" gorm:"column:courseid"`
		TimeCreated     int64  `json:"time_created" bigquery:"time_created" gorm:"column:timecreated"`
		IPAddress       string `json:"ip_address" bigquery:"ip_address" gorm:"column:ip"`
	}
)

// ShortDescription returns the first TicketShortDescriptionLength characters of the description
func (z ZendeskTicket) ShortDescription() string {
	if len(z.Description) < TicketShortDescriptionLength {
		return z.Description
	}
	return z.Description[:TicketShortDescriptionLength] + "..."
}

func (z ZendeskTicket) MarkdownString() string {
	return fmt.Sprintf("> *Ticket ⌗%d*\n*Subject*: %s\n*Status: %s* \n\n```%s```\n\n*Requested By*: %s\n*<%s|View ticket on Zendesk>*",
		z.ID, z.Subject, strings.ToUpper(z.Status), z.ShortDescription(), z.Requester, z.Link)
}

func (z ZendeskTicket) String() string {
	return fmt.Sprintf(`
New Zendesk Ticket #%d

%s

%s

Requested By: %s

%s
	`, z.ID, z.Subject, z.Description, z.Requester, z.Link)
}

func (MoodleLogLine) TableName() string {
	return "resk_logstore_standard_log"
}
