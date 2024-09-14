# Go Sample Example - Stop Ticker

This repository demonstrates how to stop a ticker in Go. Tickers are used to trigger actions at regular intervals, but stopping them when no longer needed is crucial to avoid unnecessary resource consumption.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to stop a ticker in Go.</li>
  <li>It demonstrates creating a ticker, running it for a limited time, and stopping it gracefully to avoid resource waste.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Stopping a Ticker
	// Stopping a ticker is crucial when you no longer need it to avoid unnecessary resource consumption

	// Create a ticker that ticks every 500 milliseconds
	ticker := time.NewTicker(500 * time.Millisecond)

	// Create a channel to signal the stop of the ticker
	done := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second) // Let the ticker run for 2 seconds
		ticker.Stop()
		done <- true
	}()

	// Print tick events until the ticker is stopped
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	<-done
	fmt.Println("Ticker stopped.")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `021_tickers/02_stop_ticker` directory:

```bash
cd go_sample_examples/021_tickers/02_stop_ticker
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see the following output, where the ticker prints a message every 500 milliseconds and stops after 2 seconds:

```bash
Tick at 2024-09-14 11:38:40.5725702 +0300 +03 m=+0.500607401
Tick at 2024-09-14 11:38:41.0725702 +0300 +03 m=+1.000607401
Tick at 2024-09-14 11:38:41.5725702 +0300 +03 m=+1.500607401
Ticker stopped.
```