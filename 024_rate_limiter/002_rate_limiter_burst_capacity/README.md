# Go Sample Example - Rate Limiter with Burst Capacity

This repository demonstrates a rate limiter implementation in Go that includes burst capacity. It shows how to allow a burst of requests before enforcing a rate limit.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers rate limiting techniques with burst capacity in Go.</li>
  <li>It includes an implementation that allows a burst of requests before enforcing a rate limit using a buffered channel.</li>
</ul>

## 💻 Code Example

```go
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the 024_rate_limiter/002_rate_limiter_burst_capacity directory:

   ```bash
   cd go_sample_examples/024_rate_limiter/002_rate_limiter_burst_capacity
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see output similar to:

```
Request 1 processed at 2024-09-15 17:53:36.3855689 +0300 +03 m=+0.000517001
Request 2 processed at 2024-09-15 17:53:36.4178194 +0300 +03 m=+0.032767501
Request 3 processed at 2024-09-15 17:53:36.4178194 +0300 +03 m=+0.032767501
Request 4 processed at 2024-09-15 17:53:36.5866023 +0300 +03 m=+0.201550401
Request 5 processed at 2024-09-15 17:53:36.7855888 +0300 +03 m=+0.400536901
```