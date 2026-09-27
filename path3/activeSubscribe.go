package main

import (
	"errors"
	"fmt"
)

var userData = map[string]string{
	"user1": "24-08-1989",
	"user2": "12-05-1995",
	"user3": "30-11-2000",
}

func checkUserAndData(userID string, userDate string, userData map[string]string) (bool, error) {
	// Check if the user exists in the database
	userExists := false
	for id := range userData {
		if id == userID {
			userExists = true
		}
	}

	dataExists := false
	for _, date := range userData {
		if date == userDate {
			dataExists = true
			break
		}
	}

	if !userExists {
		return false, errors.New(fmt.Sprintf("Данного пользователя не найдено %s", userID))
	}
	if !dataExists {
		return false, errors.New(fmt.Sprintf("Данной даты не найдено %s", userDate))
	}
	return true, nil
}

func main() {
	fmt.Print(checkUserAndData("user1", "24-08-1989", userData))
	fmt.Print(checkUserAndData("user1", "24-08-1976", userData))
}
