package main

import (
	"fmt"
	"time"
)

// worker — функция, представляющая нашего рабочего процесса
// Принимает id рабочего, канал задач и канал для отправки результатов
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Рабочий %d начал выполнение задачи %d\n", id, job)
		time.Sleep(time.Second) // симулируем выполнение задачи
		fmt.Printf("Рабочий %d завершил выполнение задачи %d\n", id, job)
		results <- job * 2 // отправляем результат
	}
}

func main() {
	const numJobs = 5 // количество задач для выполнения
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// создаем пул из 3 рабочих
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// отправляем задачи в канал jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// закрываем канал задач, чтобы рабочие поняли, что больше задач не будет
	close(jobs)

	// получаем результаты от воркеров
	for r := 1; r <= numJobs; r++ {
		res := <-results
		fmt.Printf("Результат: %d\n", res)
	}
}
