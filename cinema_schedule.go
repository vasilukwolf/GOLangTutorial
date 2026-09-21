package main

import (
	"fmt"
	"math/rand"
)

var sessions = []string{"10:00", "13:00", "16:00", "19:00"}
var film_names = []string{"Дюна", "Матрица", "Интерстеллар", "Начало"}
var halls = []string{"Зал 1", "Зал 2", "Зал 3"}

var halls_counter = 4

func schedule() {
	for _, hall := range halls {
		fmt.Print(hall, ":\n")
		for i := 0; i < len(sessions); i++ {
			session := sessions[i]
			film := film_names[rand.Intn(halls_counter)]
			if film == "Начало" {
				fmt.Print("Фильм \"Начало\":")
				fmt.Print("  ", session, "\n")
				break
			}
		}
	}
}

func main() {
	schedule()
}
