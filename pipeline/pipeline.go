package main

import "fmt"

// add — добавляет 2 к каждому значению из inputCh и возвращает канал с результатами
func addV2(doneCh chan struct{}, inputCh chan int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for value := range inputCh {
			result := value + 2

			select {
			case <-doneCh: // если нужно завершить горутину
				return
			case resultCh <- result: // отправляем результат
			}
		}
	}()

	return resultCh
}

// multiply — умножает каждое значение на 3 и возвращает канал с результатами
func multiplyV2(doneCh chan struct{}, inputCh chan int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for value := range inputCh {
			result := value * 3

			select {
			case <-doneCh:
				return
			case resultCh <- result:
			}
		}
	}()

	return resultCh
}

// generator — отправляет данные в канал
func generator(doneCh chan struct{}, numbers []int) chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)

		for _, num := range numbers {
			select {
			case <-doneCh:
				return
			case outputCh <- num:
			}
		}
	}()

	return outputCh
}

func main() {
	// данные, которые будем обрабатывать
	numbers := []int{1, 2, 3, 4, 5}

	// канал для остановки работы горутин
	doneCh := make(chan struct{})
	defer close(doneCh)

	// запускаем генератор, который отправляет числа
	inputCh := generator(doneCh, numbers)

	// этапы конвейера: сначала add, потом multiply
	addCh := addV2(doneCh, inputCh)
	resultCh := multiplyV2(doneCh, addCh)

	// выводим результаты
	for res := range resultCh {
		fmt.Println(res)
	}
}
