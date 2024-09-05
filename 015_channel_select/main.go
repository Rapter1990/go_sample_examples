package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, ch chan<- string) {
	time.Sleep(time.Duration(id) * time.Second)
	ch <- fmt.Sprintf("Worker %d finished", id)
}

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic Channel Select
	// Demonstrates a basic use of select with multiple channels, where select waits for either of the channels to receive data

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine to send data to ch1 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message from ch1"
	}()

	// Goroutine to send data to ch2 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	// Using select to wait for either ch1 or ch2 to receive data
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Channel Select with Timeout
	// Use the select statement to implement a timeout by using Go's time.After function.
	// This example shows how to wait for a message with a timeout

	ch := make(chan string)

	// Goroutine to send data after a delay
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "result"
	}()

	// Using select to wait for either a message or a timeout
	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: No response received")
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Channel Select with Multiple Cases and Default
	// The select statement can include a default case, which is executed if none of the other cases are ready. This is useful for non-blocking channel operations

	ch1 = make(chan string)
	ch2 = make(chan string)

	// Goroutine to send data to ch1
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message from ch1"
	}()

	// Using select with a default case
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	default:
		fmt.Println("No message received, doing other work...")
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Channel Select with a Loop
	// The select statement can be used inside a loop to continuously wait for and handle multiple channel operations

	ch1 = make(chan string)
	ch2 = make(chan string)

	// Goroutine to send data to ch1
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message from ch1"
	}()

	// Goroutine to send data to ch2
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	// Looping with select to handle both channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Channel Select with Concurrent Workers
	// This example demonstrates using select to handle multiple workers concurrently, each sending results back to a common channel

	ch = make(chan string, 3)

	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Using select in a loop to receive messages from workers
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Channel Select with Randomized Work
	// The select statement handles different durations of work, simulating a more unpredictable flow of data

	ch = make(chan string, 3)

	// Function to simulate work that takes random time
	doWork := func(id int) {
		sleepTime := rand.Intn(3) + 1
		time.Sleep(time.Duration(sleepTime) * time.Second)
		ch <- fmt.Sprintf("Worker %d finished after %d seconds", id, sleepTime)
	}

	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go doWork(i)
	}

	// Using select in a loop to receive messages from workers
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Multiple Channels with Select in Different Goroutines
	// Use select in different goroutines to handle more complex scenarios with multiple channels

	ch1 = make(chan string)
	ch2 = make(chan string)

	// Goroutine to send data to ch1
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message from ch1"
	}()

	// Goroutine to send data to ch2
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	// Another goroutine that uses select to handle messages
	go func() {
		select {
		case msg1 := <-ch1:
			fmt.Println("Worker received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Worker received:", msg2)
		}
	}()

	// Allow time for all operations to complete
	time.Sleep(3 * time.Second)

	fmt.Println("-----------------------------------------------------------------------------------")

}
