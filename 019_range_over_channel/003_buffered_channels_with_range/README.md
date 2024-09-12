# Go Sample Example - Buffered Channels with Range

This repository demonstrates how to use buffered channels with the `range` keyword in Go. It showcases how to control the flow of data using buffered channels, which allow Goroutines to send values into a channel without requiring an immediate receiver. The `range` construct is used to receive and process data from the channel until it is closed.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the basics of buffered channels and their usage in Go.</li>
  <li>It demonstrates how to send data to a buffered channel and how to use the `range` keyword to receive data until the channel is closed.</li>
  <li>Buffered channels are useful when you need to handle multiple values in a non-blocking way, allowing for better flow control in concurrent programs.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
)

func main() {

	// Buffered Channels with Range
	// Combine buffered channels with range to control the flow of data

	ch := make(chan int, 3) // Create a buffered channel

	// Send values to the buffered channel
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // Close the channel

	// Receive values using range
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

3. Navigate to the `019_range_over_channel/003_buffered_channels_with_range` directory:

```bash
   cd go_sample_examples/019_range_over_channel/003_buffered_channels_with_range
```

4. Run the Go program:

```bash
   go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```
Received: 1
Received: 2
Received: 3
Channel closed, no more data.
```