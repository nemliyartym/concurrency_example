package main

import (
	"context"
	"log"
	"time"
)

type Item struct{}

type Subscription interface {
	Updates() <-chan Item
}

type Fetcher interface {
	Fetch() (Item, error)
}

func NewSubscription(ctx context.Context, fetcher Fetcher, freq int) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
	}
	// Запуск задачи, предназначенной для получения наших данных
	go s.serve(ctx, freq)
	return s
}

type sub struct {
	fetcher Fetcher
	updates chan Item
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) serve(ctx context.Context, checkFrequency int) {
	clock := time.NewTicker(time.Duration(checkFrequency) * time.Second)
	type fetchResult struct {
		fetched Item
		err     error
	}
	fetchDone := make(chan fetchResult, 1)

	for {
		select {
		// Таймер, который запускает фетчер
		case <-clock.C:
			go func() {
				fetched, err := s.fetcher.Fetch()
				fetchDone <- fetchResult{fetched, err}
			}()
		// Случай, когда результат фетчера готов к использованию
		case result := <-fetchDone:
			fetched := result.fetched
			if result.err != nil {
				log.Println("Fetch error: %w \n Waiting the next iteration", result.err.Error())
				break
			}
			s.updates <- fetched
		// Случай, когда нам нужно закрыть сервер
		case <-ctx.Done():
			return
		}
	}
}

func (s *sub) servev2(ctx context.Context, checkFrequency int) {
	clock := time.NewTicker(time.Duration(checkFrequency) * time.Second)
	type fetchResult struct {
		fetched Item
		err     error
	}
	fetchDone := make(chan fetchResult, 1)

	var fetched Item
	var fetchResponseStream chan Item
	var pending bool

	for {
		if pending {
			fetchResponseStream = s.updates
		} else {
			fetchResponseStream = nil
		}

		select {
		// Таймер, который запускает фетчер
		case <-clock.C:
			// Ждем следующей итерации, если pending истина
			if pending {
				break
			}
			go func() {
				fetched, err := s.fetcher.Fetch()
				fetchDone <- fetchResult{fetched, err}
			}()

		// Случай, когда результат фетчера готов быть считанным
		case result := <-fetchDone:
			fetched = result.fetched
			if result.err != nil {
				log.Println("Fetch error: %w \n Waiting the next iteration", result.err.Error())
				break
			}
			pending = true

		// Данные можно отправлять по каналу
		case fetchResponseStream <- fetched:
			pending = false

		// Случай, когда нам нужно закрыть сервер
		case <-ctx.Done():
			return
		}
	}
}
