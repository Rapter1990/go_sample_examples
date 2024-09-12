# Go Sample Example - Sending a Signal with a Closed Channel

This repository demonstrates how to use a closed channel as a signal in Go. It shows how a goroutine can signal the completion of a task by closing a channel, and how the main function can wait for this signal to proceed.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example demonstrates how to use a channel to signal the completion of a task by closing the channel.</li>
  <li>The main function waits for the task goroutine to complete by reading from the closed channel, which signals that the task is done.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func task(ch chan bool) {
	time.Sleep(2 * time.Second)
	close(ch) // Close the channel to send a signal
}

func main() {

	// Sending a Signal with a Closed Channel

	// A channel is used as a signal. The task goroutine performs some work and then closes the channel to signal that the task is complete.
	// The main function waits for the signal by reading from the channel

	signal := make(chan bool)

	go task(signal)

	fmt.Println("Waiting for task to complete...")

	// Wait for the signal
	<-signal

	fmt.Println("Task completed.")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `018_closing_channel/04_sending_a_signal_with_a_closed_channel` directory:

```bash
cd go_sample_examples/018_closing_channel/04_sending_a_signal_with_a_closed_channel
```

4. Run the Go program:

```bash
go run 04_sending_a_signal_with_a_closed_channel.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Waiting for task to complete...
Task completed.
```