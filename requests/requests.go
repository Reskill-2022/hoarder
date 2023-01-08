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

	ZendeskTicketCreateRequest struct {
		ID               string `json:"id"`
		TicketType       string `json:"type"`
		Title            string `json:"title"`
		DescriptionPlain string `json:"description_plain"`
		Link             string `json:"link"`
		Via              string `json:"via"`
		Status           string `json:"status"`
		Priority         string `json:"priority"`
		LatestComment    string `json:"latest_comment"`
		Requester        string `json:"requester"`
		RequesterEmail   string `json:"requester_email"`
		Satisfaction     string `json:"satisfaction"`
		Assignee         string `json:"assignee"`
		Requested        string `json:"requested"`
	}

	CalendlyEventRequest struct {
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		EventKind string `json:"event"`
		Payload   struct {
			CancelURL           string `json:"cancel_url"`
			CreatedAt           string `json:"created_at"`
			Email               string `json:"email"`
			Event               string `json:"event"`
			Name                string `json:"name"`
			NewInvitee          string `json:"new_invitee"`
			OldInvitee          string `json:"old_invitee"`
			QuestionsAndAnswers []struct {
				Question string `json:"question"`
				Answer   string `json:"answer"`
			} `json:"questions_and_answers"`
			RescheduleURL string `json:"reschedule_url"`
			Rescheduled   bool   `json:"rescheduled"`
			Status        string `json:"status"`
			TextReminder  string `json:"text_reminder_number"`
			Timezone      string `json:"timezone"`
			UpdatedAt     string `json:"updated_at"`
			URI           string `json:"uri"`
			Canceled      bool   `json:"canceled"`
		} `json:"payload"`
	}
)
