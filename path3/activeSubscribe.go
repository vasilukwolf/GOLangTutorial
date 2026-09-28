package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrNotFound = errors.New("user not found")

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

func main() {

}
