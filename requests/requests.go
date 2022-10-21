package requests

type (
	SlackEvent struct {
		EventType string `json:"type"`
	}

	SlackChallengeRequest struct {
		Token     string `json:"token"`
		Challenge string `json:"challenge"`
	}
)
