package requests

type (
	SlackEvent struct {
		EventType string `json:"type"`
	}

	SlackChallengeRequest struct {
		Token     string `json:"token"`
		Challenge string `json:"challenge"`
	}

	SlackEventCallback struct {
		Event struct {
			Type           string `json:"type"`
			Channel        string `json:"channel"`
			User           string `json:"user"`
			Text           string `json:"text"`
			Timestamp      string `json:"ts"`
			EventTimestamp string `json:"event_ts"`
			ChannelType    string `json:"channel_type"`
		} `json:"event"`
		EventID   string `json:"event_id"`
		EventTime int64  `json:"event_time"`
		TeamID    string `json:"team_id"`
	}
)
