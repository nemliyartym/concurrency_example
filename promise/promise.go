package main

import (
	"errors"
	"fmt"
	"time"
)

type Result struct {
	value int
	err   error
}

func Promise(task func() (int, error)) chan Result {
	resultCh := make(chan Result, 1) // создаем канал для результата

	go func() {
		value, err := task()                       // выполняем задачу
		resultCh <- Result{value: value, err: err} // отправляем результат и ошибку в канал
		close(resultCh)                            // закрываем канал
	}()

	return resultCh
}

func main() {
	// Задача, которая возвращает ошибку
	taskWithError := func() (int, error) {
		time.Sleep(2 * time.Second)
		return 0, errors.New("что-то пошло не так")
	}

	// Запускаем задачу через Promise
	future := Promise(taskWithError)

	fmt.Println("Задача запущена, можно делать что-то еще...")

	// Ожидаем результат
	result := <-future
	if result.err != nil {
		fmt.Println("Ошибка:", result.err)
	} else {
		fmt.Println("Результат:", result.value)
	}
}
