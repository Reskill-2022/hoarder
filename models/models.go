package models

type (
	SlackMessage struct {
		EventID        string `json:"event_id" bigquery:"event_id"`
		EventType      string `json:"event_type" bigquery:"event_type"`
		Text           string `json:"text" bigquery:"text"`
		UserID         string `json:"user_id" bigquery:"user_id"`
		ChannelID      string `json:"channel_id" bigquery:"channel_id"`
		ChannelType    string `json:"channel_type" bigquery:"channel_type"`
		TeamID         string `json:"team_id" bigquery:"team_id"`
		Timestamp      string `json:"timestamp" bigquery:"timestamp"`
		EventTimestamp int64  `json:"event_timestamp" bigquery:"event_timestamp"`
	}
)
