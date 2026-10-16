package main

import (
	"fmt"
	"math/rand"
	"time"
)

func transcodeVideo(name string, ch chan string) {
	// имитация транскодирования: случайная задержка 500–1500 мс
	duration := time.Duration(500+rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)

	ch <- fmt.Sprintf("Файл %s перекодирован за %v", name, duration)
}

func main() {
	files := []string{"intro.mp4", "lesson1.mp4", "outro.mp4"}
	ch := make(chan string)
	for _, f := range files {
		go transcodeVideo(f, ch)
	}
	for range files {
		fmt.Println(<-ch)
	}
}
