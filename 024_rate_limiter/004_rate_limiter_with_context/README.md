# Go Sample Example - Rate Limiter with `context.Context`

This repository demonstrates how to implement rate limiting in Go using `context.Context`. It shows how to use context for request cancellation and control over processing times.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers implementing rate limiting using `context.Context` in Go.</li>
  <li>It includes techniques for request cancellation and managing processing delays.</li>
</ul>

## 💻 Code Example

```go
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the 024_rate_limiter/004_rate_limiter_with_context directory:

   ```bash
   cd go_sample_examples/024_rate_limiter/004_rate_limiter_with_context
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see output similar to:

```
Request 1 processed at 2024-09-15 18:03:29.1501573 +0300 +03 m=+0.200704201
Request 2 processed at 2024-09-15 18:03:29.3787454 +0300 +03 m=+0.429292301
Request 3 processed at 2024-09-15 18:03:29.5790582 +0300 +03 m=+0.629605101
Request 4 processed at 2024-09-15 18:03:29.7799347 +0300 +03 m=+0.830481601
Request 5 processed at 2024-09-15 18:03:29.9803233 +0300 +03 m=+1.030870201
```