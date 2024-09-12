package main

import (
	"fmt"
	"sync"
)

// Each worker sends the integers 0, 1, 2 to the channel ch
func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // wg.Done() is called to indicate that the worker has finished its task
	for i := 0; i < 3; i++ {
		ch <- i
	}
}

func main() {

	// Closing Channels in Multiple Goroutines

	// Three worker goroutines send values to the channel.
	// The channel is closed after all the workers finish their tasks.
	// The sync.WaitGroup is used to synchronize the workers and close the channel only after all of them have finished

	// ch: A channel of type int used for communication between goroutines.
	ch := make(chan int)

	// wg: A sync.WaitGroup to synchronize the completion of multiple goroutines.
	var wg sync.WaitGroup

	// Start 3 worker goroutines
	// // Three worker goroutines are started. Each worker will send values (0, 1, 2) to the channel ch
	for i := 0; i < 3; i++ {
		wg.Add(1) // wg.Add(1) increments the wait group counter by 1 for each worker goroutine
		go worker(ch, &wg)
	}

	go func() {
		wg.Wait() // A separate goroutine waits for all worker goroutines to complete using wg.Wait()
		close(ch) // Close the channel when all workers are done
	}()

	// Receiving values until the channel is closed
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("All workers done, channel closed.")

	/*
		Goroutine Execution Order
			Each worker goroutine sends three integers (0, 1, 2) to the channel ch.
			Since the goroutines execute concurrently, the exact order of sending is not guaranteed
	*/

}
