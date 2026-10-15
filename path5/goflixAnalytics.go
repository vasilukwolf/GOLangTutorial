package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Event struct {
	ID   string
	Type string
}

type Result struct {
	WorkerID int
	EventID  string
	Message  string
}

// Stats — счётчики событий, общие для всех воркеров
type Stats struct {
	mu     sync.Mutex
	total  int
	byType map[string]int
}

func NewStats() *Stats {
	return &Stats{byType: make(map[string]int)}
}

// Add безопасно увеличивает счётчики из любой горутины
func (s *Stats) Add(eventType string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.total++
	s.byType[eventType]++
}

// worker читает задачи из jobs, пока канал не закроют или не истечёт ctx, и пишет ответы в results
func worker(ctx context.Context, id int, jobs <-chan Event, results chan<- Result, stats *Stats, wg *sync.WaitGroup) {
	defer wg.Done() // сообщаем WaitGroup, что воркер закончил

	for e := range jobs {
		// имитация обработки, которую можно прервать по таймауту
		select {
		case <-time.After(100 * time.Millisecond):
		case <-ctx.Done():
			fmt.Printf("воркер %d: остановлен на событии %s — %v\n", id, e.ID, ctx.Err())
			return
		}

		stats.Add(e.Type)
		results <- Result{
			WorkerID: id,
			EventID:  e.ID,
			Message:  "обработано: " + e.Type,
		}
	}
}

func main() {
	// 1. Готовим события (вместо этого могут быть данные из JSON или БД)
	events := []Event{
		{ID: "1", Type: "Просмотр начат"},
		{ID: "2", Type: "Просмотр остановлен"},
		{ID: "3", Type: "Просмотр завершён"},
		{ID: "4", Type: "лайк поставлен"},
		{ID: "5", Type: "контент добавлен в избранное"},
		{ID: "6", Type: "Просмотр начат"},
		{ID: "7", Type: "лайк поставлен"},
		{ID: "8", Type: "Просмотр начат"},
	}

	// 2. Создаём каналы с буфером на все события
	jobs := make(chan Event, len(events))     // задачи для воркеров
	results := make(chan Result, len(events)) // ответы воркеров

	// 3. Кладём события в jobs и закрываем канал
	for _, e := range events {
		jobs <- e
	}
	close(jobs) // так воркеры поймут, что новых задач не будет

	// 4. Общий лимит времени на обработку всех событий
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel() // освобождаем ресурсы контекста, даже если уложились раньше

	// 5. Запускаем воркеров
	const workersCount = 3
	var wg sync.WaitGroup
	stats := NewStats() // один экземпляр на всех воркеров

	for i := 1; i <= workersCount; i++ {
		wg.Add(1) // до запуска горутины, иначе Wait может вернуться раньше времени
		go worker(ctx, i, jobs, results, stats, &wg)
	}

	// 6. Закрываем results, когда все воркеры закончат писать
	go func() {
		wg.Wait()
		close(results)
	}()

	// 7. Читаем результаты, пока канал не закроется
	for r := range results {
		fmt.Printf("воркер %d: событие %s — %s\n", r.WorkerID, r.EventID, r.Message)
	}

	if stats.total == len(events) {
		fmt.Println("\nОбработка завершена успешно: все события обработаны")
	} else if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("\nТаймаут: обработаны не все события")
	}

	// 8. Все воркеры завершились (results закрыт), можно читать статистику
	fmt.Printf("\nОбработано событий: %d из %d\n", stats.total, len(events))
	for t, n := range stats.byType {
		fmt.Printf("  %s: %d\n", t, n)
	}
}
