package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic Timeout with time.After
	// Set a timeout for receiving a message from a channel

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	/*

				case res := <-ch:
			         If a value is available in the channel ch, it is received, assigned to res, and "Received: res" is printed

				case <-time.After(1 * time.Second):
		             This sets a timeout using time.After(1 * time.Second). If no value is received from ch within 1 second, the timeout case will trigger, printing "Timeout: No data received"

	*/

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second): // Timeout after 1 second
		fmt.Println("Timeout: No response received")
	}

	// Approach 1: Shorten the time.Sleep Duration

	ch = make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond) // Sleep for 0.5 seconds
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second): // Timeout after 1 second
		fmt.Println("Timeout: No response received")
	}

	// Approach 2: Increase the time.After Duration

	ch = make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Sleep for 2 seconds
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(3 * time.Second): // Timeout after 3 seconds
		fmt.Println("Timeout: No response received")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Timeout with select and Multiple Cases
	// The select statement is used to handle both channel communication and timeouts simultaneously

	ch = make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "data"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No data received within 1 second")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Loop with Timeout for Multiple Attempts
	// Demonstrates how to repeatedly attempt to receive data from a channel, with a timeout for each attempt

	ch = make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "response"
	}()

	for i := 0; i < 5; i++ {
		select {
		case res := <-ch:
			fmt.Println("Received:", res)
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Attempt", i+1, "timed out")
		}
	}
	fmt.Println("All attempts timed out")

	fmt.Println("-----------------------------------------------------------------------------------")

	// Timeout in HTTP Requests
	// A common use case for timeouts is in network requests.
	// The following example demonstrates how to set a timeout on an HTTP request using Go's http package

	client := http.Client{
		Timeout: 2 * time.Second, // Set the timeout for the HTTP client
	}

	resp, err := client.Get("https://httpstat.us/200?sleep=5000")
	if err != nil {
		fmt.Println("Request timed out:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Received response with status:", resp.Status)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Timeout with Context
	// The context package provides a more structured way to handle timeouts, especially in more complex programs

	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch = make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		select {
		case ch <- "done":
		case <-ctx.Done(): // Stop sending if the context is done
		}
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Timeout for Reading from Multiple Channels
	// In scenarios where you need to read from multiple channels and
	// want to ensure that none of the reads take too long, you can use a timeout

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "data from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "data from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No data received within 1 second")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

}
