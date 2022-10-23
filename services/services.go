package services

import (
	"bytes"
	"encoding/json"
	"io"
)

type Set struct {
	SlackService   *SlackService
	ZendeskService *ZendeskService
}

func NewSet() *Set {
	return &Set{
		SlackService:   NewSlackService(),
		ZendeskService: NewZendeskService(),
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
