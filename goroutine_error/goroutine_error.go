package main

import (
	"errors"
	"log"
	"time"
)

func main() {
	input := []int{1, 2, 3, 4}

	// генератор возвращает канал с данными
	inputCh := generator(input)

	// потребитель, обрабатывающий данные
	go consumer(inputCh)

	// добавим время сна, чтобы ошибки успели вывести на экран
	time.Sleep(time.Second)
}

// generator отправляет данные в канал и закрывает его
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

// consumer принимает данные и вызывает функцию, которая возвращает ошибку
func consumer(ch chan int) {
	for data := range ch {
		err := callDatabase(data)
		if err != nil {
			log.Println(err) // простой вывод ошибки в лог
		}
	}
}

// callDatabase симулирует вызов к базе данных и всегда возвращает ошибку
func callDatabase(_ int) error {
	return errors.New("ошибка запроса к базе данных")
}
