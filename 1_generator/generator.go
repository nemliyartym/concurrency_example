package main

import "fmt"

func main() {
	// Получаем канал с данными из генератора
	dataChannel := generator()

	// Потребитель обрабатывает данные из канала
	process(dataChannel)
}

// generator создает канал и запускает горутину для отправки данных
func generator() chan int {
	// Данные, которые будут отправляться в канал
	items := []int{10, 20, 30, 40, 50}
	ch := make(chan int)

	go func() {
		// Закрываем канал после завершения отправки данных
		defer close(ch)

		// Перебираем элементы и отправляем их в канал
		for _, item := range items {
			ch <- item
		}
	}()

	return ch
}

// process получает данные из канала и выводит их
func process(ch chan int) {
	for item := range ch {
		fmt.Println(item)
	}
}
