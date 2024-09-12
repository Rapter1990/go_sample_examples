# Go Sample Example - Range Over Channel

This repository demonstrates how to use the `range` keyword to iterate over values received from a channel in Go. It showcases a basic example where one Goroutine sends values to a channel, and the main Goroutine receives and processes those values using the `range` construct.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the basics of using channels and the `range` keyword to iterate over the values sent through a channel.</li>
  <li>It includes the concept of channel creation, sending data into a channel using a Goroutine, and receiving data until the channel is closed.</li>
  <li>The code illustrates the proper use of the `close()` function to signal that no more values will be sent to the channel.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
)

func main() {

	// Basic Example: Range Over a Channel
	// A single Goroutine sends a series of integers to a channel,
	// and the main Goroutine iterates over these values using range.

	ch := make(chan int)

	// Start a Goroutine that sends values to the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel after sending all values
	}()

	// Use range to receive values from the channel until it's closed
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("Channel closed, no more data.")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `019_range_over_channel/001_basic_example_range_over_channel` directory:

```bash
   cd go_sample_examples/019_range_over_channel/001_basic_example_range_over_channel
```

4. Run the Go program:

```bash
   go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```
Received: 0
Received: 1                  
Received: 2                  
Received: 3                  
Received: 4                  
Channel closed, no more data.
```

This example demonstrates how Goroutines can be used to send data through channels and how the `range` keyword can be used to receive and process data until the channel is closed.
```
