# Go Sample Example - Timeouts

This repository demonstrates various techniques for handling timeouts in Go, including basic timeouts using `time.After`, handling timeouts in HTTP requests, and managing timeouts using Go's `context` package. It showcases how to manage blocking operations efficiently by setting timeouts.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers different ways of setting timeouts in Go, including basic timeouts for channel operations, HTTP requests, and more.</li>
  <li>It demonstrates how to use timeouts with channels, select statements, HTTP requests, and context cancellation.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic Timeout with time.After
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second): // Timeout after 1 second
		fmt.Println("Timeout: No response received")
	}

	// Shorter time.Sleep duration
	ch = make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond) // Sleep for 0.5 seconds
		ch <- "result"
	}()
	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second): // Timeout after 1 second
		fmt.Println("Timeout: No response received")
	}

	// Longer time.After duration
	ch = make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()
	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(3 * time.Second): // Timeout after 3 seconds
		fmt.Println("Timeout: No response received")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Timeout with select and Multiple Cases
	ch = make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "data"
	}()
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No data received within 1 second")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Loop with Timeout for Multiple Attempts
	ch = make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "response"
	}()

	for i := 0; i < 5; i++ {
		select {
		case res := <-ch:
			fmt.Println("Received:", res)
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Attempt", i+1, "timed out")
		}
	}
	fmt.Println("All attempts timed out")

	fmt.Println("-----------------------------------------------------------------------------------")

	// Timeout in HTTP Requests
	client := http.Client{
		Timeout: 2 * time.Second, // Set the timeout for the HTTP client
	}
	resp, err := client.Get("https://httpstat.us/200?sleep=5000")
	if err != nil {
		fmt.Println("Request timed out:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Received response with status:", resp.Status)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Timeout with Context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch = make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		select {
		case ch <- "done":
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Timeout for Reading from Multiple Channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "data from ch1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "data from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No data received within 1 second")
	}

	fmt.Println("-----------------------------------------------------------------------------------")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `016_timeouts` directory:

   ```bash
   cd go_sample_examples/016_timeouts
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
-----------------------------------------------------------------------------------
Received: message from ch2
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Timeout: No response received
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
No message received, doing other work...
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Received: message from ch2
Received: message from ch1
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Worker 1 finished
Worker 2 finished
Worker 3 finished
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Worker 1 finished after 3 seconds
Worker 3 finished after 3 seconds
Worker 2 finished after 3 seconds
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Worker received: message from ch2
-----------------------------------------------------------------------------------
```