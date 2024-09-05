package main

import (
	"fmt"
)

func main() {

	// Basic Example: Range Over a Channel
	// a single goroutine sends a series of integers to a channel,
	// and the main goroutine iterates over these values using range

	ch := make(chan int)

	// Start a goroutine that sends values to the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel after sending all values
	}()

	// Use range to receive values from the channel until it's closed
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("Channel closed, no more data.")

}
