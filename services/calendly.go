package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/constants"
	"github.com/Reskill-2022/hoarder/env"
	"github.com/Reskill-2022/hoarder/errors"
	"github.com/Reskill-2022/hoarder/models"
	"github.com/google/uuid"
)

const (
	CalendlyScheduledEventsURLFmt = "https://api.calendly.com/scheduled_events/%s"
)

type CalendlyService struct {
	conf config.Config
}

// ResolveScheduledEvent resolves a Calendly event by its UUID
func (c *CalendlyService) ResolveScheduledEvent(ctx context.Context, memberId, eventURI string) (*models.CalendlyEvent, error) {
	if eventURI == "" {
		return nil, errors.New("event URI is required", 400)
	}
	eventUUID := c.UUIDFromURI(eventURI)

	if memberId == "" {
		return nil, errors.New("member ID is required", 400)
	}

	endpoint := fmt.Sprintf(CalendlyScheduledEventsURLFmt, eventUUID)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, errors.New("error creating HTTP request", 500)
	}

	// obtain member token based on ID
	var authorizationToken string

	switch memberId {
	case constants.CalendlyMember1UUID:
		authorizationToken = c.conf.GetString(env.CalendlyMember1Token)
	case constants.CalendlyMember2UUID:
		authorizationToken = c.conf.GetString(env.CalendlyMember2Token)
	default:
		return nil, errors.New("unknown member ID", 400)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authorizationToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("error sending HTTP request", 500)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("error resolving event, got status: '%d'", resp.StatusCode), 500)
	}

	var event models.CalendlyEvent
	if err := json.NewDecoder(resp.Body).Decode(&event); err != nil {
		return nil, errors.New("error decoding response body", 500)
	}

	return &event, nil
}

// uuidFromURL returns the uuid embedded in a calendly uri
// For a URL https://api.calendly.com/users/2d30380d-afb9-4e6d-a2f2-5b210f960945,
// it returns 2d30380d-afb9-4e6d-a2f2-5b210f960945
func (c *CalendlyService) UUIDFromURI(uri string) string {
	uuidPart := uri[len(uri)-36:]
	_, err := uuid.Parse(uuidPart)
	if err != nil {
		return ""
	}
	return uuidPart
}

func NewCalendlyService(conf config.Config) *CalendlyService {
	return &CalendlyService{
		conf: conf,
	}
}
