package main

import "fmt"

func main() {
	const price = 9.99
	const title = "Базовый"

	var month = 12
	var discount = 15
	total := price * float64(month) * (1 - float64(discount)/100)
	fmt.Printf("План: %s\nЦена: %.2f\nСрок: %d месяцев\nСкидка: %d%%\nИтоговая стоимость: %.2f", title, price, month, discount, total)
}
