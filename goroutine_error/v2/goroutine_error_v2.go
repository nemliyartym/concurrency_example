package main

import (
	"errors"
	"log"
)

// Result — структура для хранения результата и ошибки
type Result struct {
	data int
	err  error
}

func main() {
	input := []int{1, 2, 3, 4}

	resultCh := make(chan Result)

	// запускаем потребителя, который будет отправлять результаты и ошибки
	go consumer(generator(input), resultCh)

	// читаем результаты
	for res := range resultCh {
		if res.err != nil {
			log.Println("Ошибка:", res.err)
		} else {
			log.Println("Результат:", res.data)
		}
	}
}

// generator отправляет данные в канал
func generator(input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)
		for _, data := range input {
			inputCh <- data
		}
	}()

	return inputCh
}

// consumer вызывает функцию, которая может возвращать ошибку
func consumer(inputCh chan int, resultCh chan Result) {
	defer close(resultCh)

	for data := range inputCh {
		resp, err := callDatabase(data)
		resultCh <- Result{data: resp, err: err}
	}
}

// callDatabase возвращает ошибку
func callDatabase(data int) (int, error) {
	return data, errors.New("ошибка запроса к базе данных")
}
