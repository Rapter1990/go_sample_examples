# Go Sample Example - Closing Channels in Multiple Goroutines

This repository demonstrates how to handle channel closing in Go across multiple goroutines. It uses the `sync.WaitGroup` to ensure that the channel is closed only after all worker goroutines have completed their tasks, and showcases how to receive values from a channel until it's closed.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example demonstrates how to synchronize the closing of channels across multiple goroutines using the `sync.WaitGroup`.</li>
  <li>It includes a simple worker pattern where multiple goroutines send values to a channel, and the channel is closed only after all workers are done.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
)

// Each worker sends the integers 0, 1, 2 to the channel ch
func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // wg.Done() is called to indicate that the worker has finished its task
	for i := 0; i < 3; i++ {
		ch <- i
	}
}

func main() {

	// Closing Channels in Multiple Goroutines

	// Three worker goroutines send values to the channel.
	// The channel is closed after all the workers finish their tasks.
	// The sync.WaitGroup is used to synchronize the workers and close the channel only after all of them have finished

	// ch: A channel of type int used for communication between goroutines.
	ch := make(chan int)

	// wg: A sync.WaitGroup to synchronize the completion of multiple goroutines.
	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1) // Increment the wait group counter for each worker goroutine
		go worker(ch, &wg)
	}

	go func() {
		wg.Wait() // Wait for all worker goroutines to complete
		close(ch) // Close the channel when all workers are done
	}()

	// Receiving values until the channel is closed
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("All workers done, channel closed.")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `018_closing_channel/03_closing_channels_in_multiple_goroutines` directory:

```bash
cd go_sample_examples/018_closing_channel/03_closing_channels_in_multiple_goroutines
```

4. Run the Go program:

```bash
go run 03_closing_channels_in_multiple_goroutines.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Received: 0
Received: 0                      
Received: 0                      
Received: 1                      
Received: 1                      
Received: 1                      
Received: 2                      
Received: 2                      
Received: 2                      
All workers done, channel closed.
```