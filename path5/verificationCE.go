package main

import (
	"fmt"
	"sync"
	"time"
)

func checkEpisode(title string, delay int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(delay) * time.Second)
	wg.Done()
	fmt.Printf("Episode %s is verified after %d seconds\n", title, delay)
}

type Episode struct {
	Title string
	Delay int
}

func main() {
	episodes := []Episode{
		{"S01E01", 3},
		{"S01E02", 1},
		{"S01E03", 2},
	}

	var wg sync.WaitGroup

	for _, ep := range episodes {
		wg.Add(1) // до запуска горутины
		go checkEpisode(ep.Title, ep.Delay, &wg)
	}

	wg.Wait() // ждём, пока все горутины вызовут Done
	fmt.Println("All episodes verified")
}
