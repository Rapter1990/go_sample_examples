Here is the `README.md` for the provided code in the same format as the previous example:

# Go Sample Example - Detecting Closed Channels

This repository demonstrates how to detect when a channel is closed in Go. It uses the `value, ok := <-ch` pattern to check if the channel is closed and gracefully exits the loop when no more data is available.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers detecting the closure of a channel in Go.</li>
  <li>It uses the boolean `ok` value from channel operations to identify when a channel is closed, allowing the program to handle it appropriately.</li>
</ul>

## 💻 Code Example

```go
package main

import "fmt"

func main() {

	// Detecting Closed Channel
	// value, ok := <-ch is used to check if the channel is closed.
	// The ok will be false when the channel is closed, and the loop breaks.

	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel after sending all values
		done <- true
	}()

	// Continuously receive until done is true
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed.")
			break
		}
		fmt.Println("Received:", value)
	}
	<-done // Ensure the goroutine has finished
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `018_closing_channel/02_detecting_closed_channel` directory:

```bash
cd go_sample_examples/018_closing_channel/02_detecting_closed_channel
```

4. Run the Go program:

```bash
go run 02_detecting_closed_channel.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Received: 0
Received: 1       
Received: 2       
Received: 3       
Received: 4       
Channel is closed.
```