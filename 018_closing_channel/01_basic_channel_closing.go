package main

import "fmt"

func main() {

	// Basic Channel Closing
	// In this example, the channel is closed using close(ch) after all values are sent.
	// The range loop is used to receive values until the channel is closed.
	// Once the channel is closed, the loop automatically exits.

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel after sending all values
	}()

	// Receiving values until the channel is closed
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("Channel closed, no more data.")
}
