package main

import (
	"context"
	"fmt"
)

func printIntegers(done <-chan struct{}, intStream <-chan int) {
	for {
		select {
		case i := <-intStream:
			fmt.Println(i)
		case <-done:
			return
		}
	}
}

func printIntegersCtx(ctx context.Context, intStream <-chan int) {
	for {
		select {
		case i := <-intStream:
			fmt.Println(i)
		case <-ctx.Done():
			return
		}
	}
}
