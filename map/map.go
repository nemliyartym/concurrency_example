package main

func Map(done <-chan struct{}, inputStream <-chan int, operator func(int) int) <-chan int {
	// 1
	mappedStream := make(chan int)
	go func() {
		defer close(mappedStream)
		// 2
		for {
			select {
			case <-done:
				return
			// 3
			case i, ok := <-inputStream:
				if !ok {
					return
				}

				//4
				select {
				case <-done:
					return
				case mappedStream <- operator(i):
				}
			}
		}
	}()
	return mappedStream
}
