package main

import (
	"fmt"
)

var action = []string{"play", "pause", "stop", "next", "prev", "volume_up", "volume_down", "info", "like", "quit", "share"}

func identifyType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("Тип: int, Значение:\n", val)
	case string:
		fmt.Println("Тип: string, Значение:\n", val)
	case bool:
		fmt.Println("Тип: bool, Значение:\n", val)
	case float64:
		fmt.Println("Тип: float64, Значение:\n", val)
	default:
		fmt.Println("Неизвестный тип")
	}
}

func player(action string) {
	switch action {
	case "play":
		fmt.Println("Проигрывание...")
	case "pause":
		fmt.Println("Пауза...")
	case "stop":
		fmt.Println("Остановка...")
	case "next":
		fmt.Println("Следующее видео...")
	case "prev":
		fmt.Println("Предыдущее видео...")
	case "volume_up":
		fmt.Println("Увеличение громкости...")
	case "volume_down":
		fmt.Println("Уменьшение громкости...")
	case "like":
		fmt.Println("Лайк текущего видео...")
	case "quit":
		fmt.Println("Выход...")
	case "share":
		fmt.Println("Поделиться текущим видео...")
	default:
		fmt.Println("Неверное действие.")
	}
	switch {
	case "info" == action:
		fmt.Println("Информация о текущем видео: Название, Автор, Длительность, Рейтинг")
	}
}

func main() {
	for _, action := range action {
		player(action)
	}
	identifyType("play")
}
