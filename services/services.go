package services

type Set struct {
	SlackService *SlackService
}

func NewSet() *Set {
	return &Set{
		SlackService: NewSlackService(),
	}
}
