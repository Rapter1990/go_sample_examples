package main

import (
	"fmt"
	"time"
)

func main() {

	// Custom Rate Limiting Using time.After
	// Show a more flexible approach where each request is rate-limited individually.

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		// Delay processing each request by 200 milliseconds
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Request", req, "processed at", time.Now())
	}

}
