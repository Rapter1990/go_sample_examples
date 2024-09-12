package main

import (
	"fmt"
	"time"
)

func main() {

	// Rate Limiting Using time.NewTicker
	// Use time.NewTicker for more control over the ticker's lifecycle

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Create a ticker that ticks every 200 milliseconds
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for req := range requests {
		<-ticker.C // Wait for the next tick
		fmt.Println("Request", req, "processed at", time.Now())
	}

}
