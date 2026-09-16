package main

import "fmt"

func tiketPrice(basePrice float64, evening bool, weekend bool) float64 {
	if evening {
		basePrice += 50.0
	}
	if weekend {
		basePrice += 100.0
	}
	return basePrice

}
func main() {
	ticketPrice := tiketPrice(200.0, true, false)
	fmt.Println("Цена билета:", ticketPrice)
}
