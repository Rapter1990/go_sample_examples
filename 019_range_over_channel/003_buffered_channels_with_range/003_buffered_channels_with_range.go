package main

import (
	"fmt"
)

func main() {

	// Buffered Channels with Range
	// Combine buffered channels with range to control the flow of data

	ch := make(chan int, 3) // Create a buffered channel

	// Send values to the buffered channel
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // Close the channel

	// Receive values using range
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("Channel closed, no more data.")
}
