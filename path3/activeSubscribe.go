package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrUserNotFound = errors.New("user not found")

type ValidtionError struct {
	Field string
	Value string
}

func (e *ValidtionError) Error() string {
	return fmt.Sprintf("validation error: field=%s value=%s", e.Field, e.Value)
}

func ParseExpiry(input string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", input)
	if err != nil {
		return time.Time{}, &ValidtionError{Field: "expiresAt", Value: input}
	}

	return t, nil
}

func ActivateSubscription(userID int, expiresAt string) error {
	if userID > 1 {
		return fmt.Errorf("activate subsctiption: %w", ErrUserNotFound)
	}

	if _, err := ParseExpiry(expiresAt); err != nil {
		return fmt.Errorf("activate subsctiption: %w", err)
	}
	return nil
}

func main() {
	err := ActivateSubscription(0, "2026-12-31")
	fmt.Println(errors.Is(err, ErrUserNotFound))

	err = ActivateSubscription(2, "2006-01-02")
	fmt.Println(errors.Is(err, ErrUserNotFound))

	err = ActivateSubscription(1, "not-a-date")
	var ve *ValidtionError
	if errors.As(err, &ve) {
		fmt.Printf("%s=%s\n", ve.Field, ve.Value)

	}

}
