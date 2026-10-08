package main

import "fmt"

type Subscription struct {
	Plan      string
	Price     float64
	AutoRenew bool
}

func NewSubscription(plan string, price float64, autoRenew bool) *Subscription {
	return &Subscription{
		Plan:      plan,
		Price:     price,
		AutoRenew: true,
	}
}

func cancelAutoRenew(subscription *Subscription) {
	subscription.AutoRenew = false
}

func PrintDetails(subscription *Subscription) {
	fmt.Printf("Plan: %s\n", subscription.Plan)
	fmt.Printf("Price: $%.2f\n", subscription.Price)
	fmt.Printf("Auto-Renew: %t\n", subscription.AutoRenew)
}

func main() {
	subscription := NewSubscription("Premium", 29.99, true)
	PrintDetails(subscription)
	cancelAutoRenew(subscription)
	PrintDetails(subscription)
}
