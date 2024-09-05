package main

import (
	"fmt"
	"time"
)

func main() {

	// In Go, channels are blocking by nature, meaning that a goroutine will be blocked when sending to or receiving from a channel until the operation completes.
	// However, by using the `select` statement, non-blocking operations on channels can be achieved.

	fmt.Println("-----------------------------------------------------------------------------------")

	// Non-Blocking Receive
	// Demonstrate how to attempt to receive a value from a channel without blocking if no value is available

	ch := make(chan string)

	// Non-blocking receive from channel
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Non-Blocking Send
	// Show how to attempt to send a value to a channel without blocking if the channel is full or unbuffered

	ch = make(chan string, 1)

	// Send a message to the channel
	ch <- "hello"

	// Non-blocking send to channel
	select {
	case ch <- "world":
		fmt.Println("Sent message to channel")
	default:
		fmt.Println("Channel is full, message not sent")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Non-Blocking Multiple Operations
	// Demonstrate how to handle multiple non-blocking operations (both send and receive) in a select statement

	ch = make(chan string)
	done := make(chan bool)

	// Start a goroutine to send a message to ch
	go func() {
		ch <- "hello"
		done <- true
	}()

	time.Sleep(1 * time.Second)

	// Non-blocking receive and done signal handling
	select {
	case msg := <-ch: // checks if there's a message to receive from the ch channel
		fmt.Println("Received from ch:", msg)
	case <-done: // checks if the done channel has received a value.
		fmt.Println("Done signal received")
	default:
		fmt.Println("No activity")
	}

	ch = make(chan string)
	done = make(chan bool)

	// Start a goroutine to send a message to ch
	go func() {
		done <- true                       // Send to done first
		time.Sleep(100 * time.Millisecond) // Delay before sending to ch
		ch <- "hello"
	}()

	time.Sleep(1 * time.Second)

	// Non-blocking receive and done signal handling
	select {
	case msg := <-ch:
		fmt.Println("Received from ch:", msg)
	case <-done:
		fmt.Println("Done signal received")
	default:
		fmt.Println("No activity")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Non-Blocking Channel Operation with Timeout
	// Introduces a timeout using time.After in combination with non-blocking operations

	ch = make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout, no message received")
	}

	// Attempt to send without blocking
	select {
	case ch <- "test":
		fmt.Println("Message sent to ch")
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout, message not sent")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

}
