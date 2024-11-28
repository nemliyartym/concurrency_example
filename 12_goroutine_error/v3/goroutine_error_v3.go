package main

import (
	"context"
	"errors"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, _ := errgroup.WithContext(context.Background())
	input := []int{1, 2, 3, 4}

	inputCh := generator(input)

	for data := range inputCh {
		data := data // создание новой переменной для каждого запуска горутины
		g.Go(func() error {
			err := callDatabase(data)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Println("Ошибка:", err)
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

// callDatabase возвращает ошибку, если data равно 3
func callDatabase(data int) error {
	if data == 3 {
		return errors.New("ошибка запроса к базе данных")
	}
	return nil
}
