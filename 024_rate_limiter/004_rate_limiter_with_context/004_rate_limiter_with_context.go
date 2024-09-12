package main

import (
	"context"
	"fmt"
	"time"
)

func processRequest(ctx context.Context, req int) {
	select {
	case <-ctx.Done():
		fmt.Println("Request", req, "cancelled at", time.Now())
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Request", req, "processed at", time.Now())
	}
}

func main() {

	// Rate Limiting with context.Context
	// Use context.Context to implement rate limiting, which allows for more control over request cancellation

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		processRequest(ctx, req)
		cancel()
	}

}
