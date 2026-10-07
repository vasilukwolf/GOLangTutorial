package main

import "fmt"

type Profile struct {
	UserName string
	Email    string
	Age      int
	Country  string
	Premium  bool
}

func main() {
	var Profile2 Profile
	fmt.Println("\nEmpty Profile2:", Profile2)

	Profile := Profile{
		UserName: "Alice",
		Email:    "alice@example.com",
		Age:      30,
		Country:  "USA",
		Premium:  true,
	}
	fmt.Printf("User Profile:\nUsername: %s\nEmail: %s\nAge: %d\nCountry: %s\nPremium: %t\n",
		Profile.UserName, Profile.Email, Profile.Age, Profile.Country, Profile.Premium)

}
