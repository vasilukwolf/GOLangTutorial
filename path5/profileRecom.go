package main

import (
	"fmt"
	"time"
)

func loadUserTaste(ch chan<- string) {
	time.Sleep(500 * time.Millisecond) // имитация загрузки
	ch <- "комедии, фантастика"
	time.Sleep(300000 * time.Millisecond) // имитация загрузки
	ch <- "Интерстеллар, Мальчишник в Вегасе"
}

func main() {
	tasteCh := make(chan string, 2) // буфер: горутина не зависнет, если после таймаута результат никто не читает

	go loadUserTaste(tasteCh)

	timeout := time.After(1 * time.Second) // общий лимит ожидания

	// ждём два значения от loadUserTaste, но не дольше таймаута
	for i := 0; i < 2; i++ {
		select {
		case taste := <-tasteCh:
			fmt.Println("Вкусы пользователя:", taste)
		case <-timeout:
			fmt.Println("Превышено время ожидания")
			return
		}
	}
}
