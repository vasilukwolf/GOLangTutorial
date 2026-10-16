package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func analyticsWorker(
	ctx context.Context,
	id int,
	jobs <-chan string,
	results chan<- string,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	processed *int,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Воркер", id, "остановлен по таймауту")
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}
			time.Sleep(350 * time.Millisecond)
			mu.Lock()
			*processed++
			mu.Unlock()
			results <- fmt.Sprintf("Воркер %d обработал событие %s", id, job)
		}
	}
}
func main() {
	events := []string{
		"просмотр начат",
		"просмотр остановлен",
		"просмотр завершен",
		"лайк поставлен",
		"контент добавлен в избранное",
	}
	jobs := make(chan string)
	results := make(chan string)
	ctx, chanel := context.WithTimeout(context.Background(), 1400*time.Millisecond)
	defer chanel()
	var wg sync.WaitGroup
	var mu sync.Mutex
	processed := 0
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go analyticsWorker(ctx, i, jobs, results, &wg, &mu, &processed)
	}

	go func() {
		for _, event := range events {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			case jobs <- event:
			}
		}
		close(jobs)
	}()
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Println(result)
	}
	if ctx.Err() != nil {
		fmt.Println("Обработка остановлена", ctx.Err())
	} else {
		fmt.Println("Все события обработаны успешно")
	}
	fmt.Println("Количество обработаных событий", processed)
}
