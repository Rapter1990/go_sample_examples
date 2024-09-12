package main

import (
	"fmt"
	"time"
)

func task(ch chan bool) {
	time.Sleep(2 * time.Second)
	close(ch) // Close the channel to send a signal
}

func main() {

	// Sending a Signal with a Closed Channel

	// a channel is used as a signal. The task goroutine performs some work and then closes the channel to signal that the task is complete.
	// The main function waits for the signal by reading from the channel

	signal := make(chan bool)

	go task(signal)

	fmt.Println("Waiting for task to complete...")

	// Wait for the signal
	<-signal

	fmt.Println("Task completed.")
}
