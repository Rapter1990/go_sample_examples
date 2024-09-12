# Go Sample Example - Channel Select

This repository demonstrates various use cases of the `select` statement with channels in Go. It covers examples of basic channel operations, timeouts, non-blocking selects, loops with select, concurrent workers, and handling multiple channels.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers various channel operations using the select statement in Go.</li>
  <li>It includes examples with timeouts, non-blocking selects, and managing multiple goroutines.</li>
  <li>It demonstrates how to use channels to coordinate concurrent tasks and handle different scenarios using the select statement.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, ch chan<- string) {
	time.Sleep(time.Duration(id) * time.Second)
	ch <- fmt.Sprintf("Worker %d finished", id)
}

func main() {
	// Basic Channel Select
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}

	// Channel Select with Timeout
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: No response received")
	}

	// Channel Select with Multiple Cases and Default
	ch1 = make(chan string)
	ch2 = make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message from ch1"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	default:
		fmt.Println("No message received, doing other work...")
	}

	// Channel Select with a Loop
	ch1 = make(chan string)
	ch2 = make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}

	// Channel Select with Concurrent Workers
	ch = make(chan string, 3)
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}

	// Channel Select with Randomized Work
	ch = make(chan string, 3)

	doWork := func(id int) {
		sleepTime := rand.Intn(3) + 1
		time.Sleep(time.Duration(sleepTime) * time.Second)
		ch <- fmt.Sprintf("Worker %d finished after %d seconds", id, sleepTime)
	}

	for i := 1; i <= 3; i++ {
		go doWork(i)
	}

	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}

	// Multiple Channels with Select in Different Goroutines
	ch1 = make(chan string)
	ch2 = make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	go func() {
		select {
		case msg1 := <-ch1:
			fmt.Println("Worker received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Worker received:", msg2)
		}
	}()

	time.Sleep(3 * time.Second)
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `015_channel_select` directory:

   ```bash
   cd go_sample_examples/015_channel_select
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see output similar to the following:

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