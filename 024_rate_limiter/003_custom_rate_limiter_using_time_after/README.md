# Go Sample Example - Custom Rate Limiter Using `time.After`

This repository demonstrates a custom rate limiter implementation in Go using `time.After`. It showcases a flexible approach to rate-limit each request individually.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers custom rate limiting techniques using `time.After` in Go.</li>
  <li>It includes an implementation where each request is delayed by a specified time interval before processing.</li>
</ul>

## 💻 Code Example

```go
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the 024_rate_limiter/003_custom_rate_limiter_using_time_after directory:

   ```bash
   cd go_sample_examples/024_rate_limiter/003_custom_rate_limiter_using_time_after
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see output similar to:

```
Request 1 processed at 2024-09-15 17:55:27.0855462 +0300 +03 m=+0.201428901
Request 2 processed at 2024-09-15 17:55:27.3143205 +0300 +03 m=+0.430203201
Request 3 processed at 2024-09-15 17:55:27.5147307 +0300 +03 m=+0.630613401
Request 4 processed at 2024-09-15 17:55:27.7153769 +0300 +03 m=+0.831259601
Request 5 processed at 2024-09-15 17:55:27.9158934 +0300 +03 m=+1.031776101
```