# Go Sample Example - Closing Channels

This repository demonstrates the concept of closing channels in Go. It shows how to close a channel after sending all values and how to use a `range` loop to receive values until the channel is closed.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic channel closing operations in Go.</li>
  <li>It explains how to close a channel after sending all values and how to handle channel closure during value reception using a range loop.</li>
</ul>

## 💻 Code Example

```go
package main

import "fmt"

func main() {

	// Basic Channel Closing
	// In this example, the channel is closed using close(ch) after all values are sent.
	// The range loop is used to receive values until the channel is closed.
	// Once the channel is closed, the loop automatically exits.

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel after sending all values
	}()

	// Receiving values until the channel is closed
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

3. Navigate to the `018_closing_channel/01_basic_channel_closing` directory:

```bash
cd go_sample_examples/018_closing_channel/01_basic_channel_closing
```

4. Run the Go program:

```bash
go run 01_basic_channel_closing.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Received: 0
Received: 1                  
Received: 2                  
Received: 3                  
Received: 4                  
Channel closed, no more data.
```
