package main

import "fmt"

func main() {
	time_in_seconds := 8520
	hours := time_in_seconds / 3600
	minutes := (time_in_seconds % 3600) / 60
	seconds := time_in_seconds % 60
	fmt.Printf(" %d часов, %d минут, %d секунд\n", hours, minutes, seconds)
}
