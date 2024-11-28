package main

import (
	"context"
)

func takeFirstN(ctx context.Context, dataSource <-chan interface{}, n int) <-chan interface{} {
	// 1
	takeChannel := make(chan interface{})

	// 2
	go func() {
		defer close(takeChannel)

		// 3
		for i := 0; i < n; i++ {
			select {
			case val, ok := <-dataSource:
				if !ok {
					return
				}
				takeChannel <- val
			case <-ctx.Done():
				return
			}
		}
	}()
	return takeChannel
}
