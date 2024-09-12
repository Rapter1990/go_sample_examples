package main

import (
	"fmt"
	"time"
)

func main() {

	// Basic Rate Limiting Using time.Tick
	// Show how to limit the processing of requests to one every 200 milliseconds.

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Create a rate limiter that ticks every 200 milliseconds
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter // Wait for the next tick
		fmt.Println("Request", req, "processed at", time.Now())
	}

}
