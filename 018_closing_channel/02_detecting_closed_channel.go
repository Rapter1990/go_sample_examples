package main

import "fmt"

func main() {

	// Detecting Closed Channel

	// value, ok := <-ch is used to check if the channel is closed.
	// The ok will be false when the channel is closed, and the loop break

	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel after sending all values
		done <- true
	}()

	// Continuously receive until done is true
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed.")
			break
		}
		fmt.Println("Received:", value)
	}
	<-done // Ensure the goroutine has finished
}
