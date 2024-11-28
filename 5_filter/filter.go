package main

func Filter(done <-chan struct{}, inputStream <-chan int, operator func(int) bool) <-chan int {
	filteredStream := make(chan int)
	go func() {
		defer close(filteredStream)

		for {
			select {
			case <-done:
				return
			case i, ok := <-inputStream:
				if !ok {
					return
				}

				if !operator(i) {
					break
				}
				select {
				case <-done:
					return
				case filteredStream <- i:
				}
			}
		}
	}()
	return filteredStream
}
