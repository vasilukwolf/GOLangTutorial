package main

import (
	"encoding/json"
	"fmt"
)

type Episode struct {
	Title    string `json:"title"`
	Season   int    `json:"season"`
	Duration int    `json:"duration"`
}

type WatchHistory struct {
	Profile         Profile `json:"profile"`
	Episode         Episode `json:"episode"`
	WathchedPercent int     `json:"watched_percent"`
	LastPositionMin int     `json:"last_position_min"`
}

type Profile struct {
	UserName string
	Email    string
	Age      int
	Country  string
	Premium  bool
}

func NewProfile(userName, email string, age int, country string) *Profile {
	return &Profile{
		UserName: userName,
		Email:    email,
		Age:      age,
		Country:  country,
		Premium:  false,
	}
}

func main() {
	p := NewProfile("Alice", "alice@example.com", 30, "USA")

	history := WatchHistory{
		Profile:         *p,
		Episode:         Episode{Title: "Pilot", Season: 1, Duration: 45},
		WathchedPercent: 80,
		LastPositionMin: 36,
	}
	fmt.Printf("Watch History:\nProfile: %+v\nEpisode: %+v\nWatched Percent: %d%%\nLast Position: %d min\n",
		history.Profile, history.Episode, history.WathchedPercent, history.LastPositionMin)

	data, err := json.Marshal(history.Episode)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
