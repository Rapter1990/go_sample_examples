package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ch <- id*10 + i
	}
}

func main() {

	// Multiple Goroutines Sending to a Channel
	// we'll have multiple goroutines sending values to the same channel.
	// The range will receive all these values until the channel is closed

	ch := make(chan int)
	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	// Close the channel once all workers are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive and print values from the channel until it's closed
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("All workers done, channel closed.")
}
