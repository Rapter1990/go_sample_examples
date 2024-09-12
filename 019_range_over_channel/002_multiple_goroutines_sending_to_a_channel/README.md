# Go Sample Example - Multiple Goroutines Sending to a Channel

This repository demonstrates how to use multiple Goroutines to send values to a single channel in Go. It showcases a scenario where multiple worker Goroutines send values to the same channel, and the main Goroutine receives and processes those values until the channel is closed.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the basics of concurrent programming in Go using Goroutines and channels.</li>
  <li>It demonstrates how to send data from multiple Goroutines into the same channel and how to use the `range` construct to receive data until the channel is closed.</li>
  <li>The code includes synchronization with a `WaitGroup` to ensure that all Goroutines complete their work before closing the channel.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ch <- id*10 + i
	}
}

func main() {

	// Multiple Goroutines Sending to a Channel
	// We'll have multiple Goroutines sending values to the same channel.
	// The range will receive all these values until the channel is closed.

	ch := make(chan int)
	var wg sync.WaitGroup

	// Start 3 worker Goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	// Close the channel once all workers are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive and print values from the channel until it's closed
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

3. Navigate to the `019_range_over_channel/002_multiple_goroutines_sending_to_a_channel` directory:

```bash
   cd go_sample_examples/019_range_over_channel/002_multiple_goroutines_sending_to_a_channel
```

4. Run the Go program:

```bash
   go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```
Received: 10
Received: 20                     
Received: 30                     
Received: 11                     
Received: 21                     
Received: 31                     
Received: 12                     
Received: 22                     
Received: 32                     
All workers done, channel closed.
```