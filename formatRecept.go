package main

import "fmt"

func formatRecept(filmName string, numberOfTickets int, priceOfTicket float64) string {
	totalPrice := float64(numberOfTickets) * priceOfTicket
	return fmt.Sprintf("Abkmv %s, Количество билетов: %d, Цена билета: %.2f\nИтого: %.2f", filmName, numberOfTickets, priceOfTicket, totalPrice)
}

func main() {
	filmName := "Inception"
	numberOfTickets := 3
	priceOfTicket := 200.0
	recept := formatRecept(filmName, numberOfTickets, priceOfTicket)
	fmt.Println(recept)
}
