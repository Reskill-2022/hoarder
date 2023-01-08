package services

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/Reskill-2022/hoarder/config"
)

type Set struct {
	SlackService    *SlackService
	ZendeskService  *ZendeskService
	CalendlyService *CalendlyService
	MoodleService   *MoodleService
}

func NewSet(conf config.Config) *Set {
	return &Set{
		SlackService:    NewSlackService(conf),
		ZendeskService:  NewZendeskService(),
		CalendlyService: NewCalendlyService(conf),
		MoodleService:   NewMoodleService(conf),
	}
}

// JSONPayloadReader returns a JSON reader for the payload
// It returns nil if marshalling fails
func JSONPayloadReader(payload map[string]interface{}) io.Reader {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil
	}
	return bytes.NewReader(b)
}

func caselessEqual(a, b string) bool {
	return strings.EqualFold(a, b)
}

func cleanToLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func clean(s string) string {
	return strings.TrimSpace(s)
}
