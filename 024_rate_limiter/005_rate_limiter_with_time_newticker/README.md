# Go Sample Example - Rate Limiter with `time.NewTicker`

This repository demonstrates how to implement rate limiting in Go using `time.NewTicker`. It shows how to use a ticker to control the rate at which requests are processed.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers implementing rate limiting using `time.NewTicker` in Go.</li>
  <li>It includes techniques for controlling the processing rate of requests with a ticker.</li>
</ul>

## 💻 Code Example

```go
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the 024_rate_limiter/005_rate_limiter_with_time_newticker directory:

   ```bash
   cd go_sample_examples/024_rate_limiter/005_rate_limiter_with_time_newticker
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see output similar to:

```
Request 1 processed at 2024-09-15 18:03:59.5443568 +0300 +03 m=+0.200915601
Request 2 processed at 2024-09-15 18:03:59.7447692 +0300 +03 m=+0.401328001
Request 3 processed at 2024-09-15 18:03:59.9442464 +0300 +03 m=+0.600805201
Request 4 processed at 2024-09-15 18:04:00.1447645 +0300 +03 m=+0.801323301
Request 5 processed at 2024-09-15 18:04:00.3447322 +0300 +03 m=+1.001291001
```