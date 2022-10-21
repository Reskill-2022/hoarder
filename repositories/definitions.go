package repositories

type (
	SlackMessageCreator interface {
		CreateMessage() error
	}
)
