package main

import "fmt"

type Notifier interface {
	Send() string
	Channel() string
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Send() string {
	return "Письмо отправлено на " + e.Email
}

func (e EmailNotifier) Channel() string {
	return "email"
}

type PushNotifier struct {
	PhoneNumber string
}

func (p PushNotifier) Send() string {
	return "Push отправлен на " + p.PhoneNumber
}

func (p PushNotifier) Channel() string {
	return "push"
}

var _ Notifier = EmailNotifier{}
var _ Notifier = PushNotifier{}

func main() {
	notifiers := []Notifier{
		EmailNotifier{Email: "alice@example.com"},
		PushNotifier{PhoneNumber: "+79991234567"},
	}

	for _, n := range notifiers {
		fmt.Printf("[%s] %s\n", n.Channel(), n.Send())
	}
}
