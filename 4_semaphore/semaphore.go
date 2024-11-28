package main

import (
	"log"
	"sync"
	"time"
)

// Semaphore — структура для управления количеством параллельных горутин
type Semaphore struct {
	semaCh chan struct{}
}

// NewSemaphore — создает новый семафор с заданной максимальной емкостью
func NewSemaphore(maxReq int) *Semaphore {
	return &Semaphore{
		semaCh: make(chan struct{}, maxReq),
	}
}

// Acquire — резервирует место в семафоре
func (s *Semaphore) Acquire() {
	s.semaCh <- struct{}{}
}

// Release — освобождает место в семафоре
func (s *Semaphore) Release() {
	<-s.semaCh
}

func main() {
	var wg sync.WaitGroup

	// создаем семафор, который позволит работать только двум горутинам одновременно
	semaphore := NewSemaphore(2)

	// запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(taskID int) {
			// резервируем место в семафоре перед началом работы
			semaphore.Acquire()

			// когда горутина завершает работу, освобождаем место и уменьшаем счетчик WaitGroup
			defer wg.Done()
			defer semaphore.Release()

			// симулируем работу горутины
			log.Printf("Запущен рабочий %d", taskID)
			time.Sleep(1 * time.Second)
		}(i)
	}

	// ждем завершения всех горутин
	wg.Wait()
}
