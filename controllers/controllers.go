package controllers

type Set struct {
	SlackController *SlackController
}

func NewSet() *Set {
	return &Set{
		SlackController: NewSlackController(),
	}
}
