package main

import (
	"fmt"
	"time"
)

func main() {

	// Rate Limiting with a Burst Capacity
	// Demonstrate how to allow a burst of requests before enforcing the rate limit

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Create a buffered channel to allow a burst of up to 3 requests
	burstLimiter := make(chan time.Time, 3)

	// Initially fill the burstLimiter channel to allow immediate processing of up to 3 requests
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	// Fill the burstLimiter channel every 200 milliseconds to maintain the rate limit
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	for req := range requests {
		<-burstLimiter // Wait for a token from burstLimiter
		fmt.Println("Request", req, "processed at", time.Now())
	}

}
