package main

type Generator struct {
	// канал для буферизации новых значений
	ch chan string
}

// Конструктор создания Generator
func NewGenerator() *Generator {
	g := &Generator{
		// канал для складирования новых значений
		ch: make(chan string, 10),
	}

	// воркер генератора
	go func() {
		for {
			// кладём в канал новое значение
			g.ch <- "genrate_value"
		}

	}()

	return g
}

// Получения новго значения из канал
func (g *Generator) GetValue() string {
	return <-g.ch
}
