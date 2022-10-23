package services

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
