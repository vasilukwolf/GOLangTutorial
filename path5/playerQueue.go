package main

import (
	"fmt"
	"time"
)

func main() {
	playerChannel := make(chan string, 10)

	playerChannel <- "Play"
	playerChannel <- "Pause"
	playerChannel <- "Stop"
	playerChannel <- "Seek"

	fmt.Println("Все 4 события успешно отправлены в буфер без блокировки!")

	close(playerChannel)

	// 4. Читаем события из канала
	for event := range playerChannel {
		fmt.Printf("Обработано -> %s\n", event)
		time.Sleep(500 * time.Millisecond) // Имитация обработки
	}
	fmt.Println("Все события обработаны, канал закрыт.")
}
