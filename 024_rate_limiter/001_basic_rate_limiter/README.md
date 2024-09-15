# Go Sample Example - Basic Rate Limiter

This repository demonstrates a basic rate limiter implementation in Go. It shows how to control the rate of request processing using Go's `time.Tick` function.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic rate limiting techniques in Go.</li>
  <li>It includes a simple implementation to limit the processing of requests at a fixed interval.</li>
</ul>

## 💻 Code Example

```go
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the 024_rate_limiter/001_basic_rate_limiter directory:

   ```bash
   cd go_sample_examples/024_rate_limiter/001_basic_rate_limiter
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
Request 1 processed at 2024-09-15 17:49:27.3630742 +0300 +03 m=+0.200725901
Request 2 processed at 2024-09-15 17:49:27.5638677 +0300 +03 m=+0.401519401
Request 3 processed at 2024-09-15 17:49:27.7631892 +0300 +03 m=+0.600840901
Request 4 processed at 2024-09-15 17:49:27.9639358 +0300 +03 m=+0.801587501
Request 5 processed at 2024-09-15 17:49:28.1636925 +0300 +03 m=+1.001344201
```