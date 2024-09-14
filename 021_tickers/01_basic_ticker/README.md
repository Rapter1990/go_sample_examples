# Go Sample Example - Basic Ticker

This repository demonstrates the usage of tickers in Go, which are used to trigger actions at regular intervals. The example showcases how to create a simple ticker that ticks every second and stops after a specified duration.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic ticker usage in Go.</li>
  <li>It demonstrates how to create a ticker, run it for a set amount of time, and stop it after a certain number of ticks.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Tickers are used to perform an action at regular intervals for a specified duration.
	// A ticker can be created using the NewTicker function from the time package

	// Simple Ticker
	// create a ticker that ticks every 1 second

	// Create a ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)

	// Stop the ticker after 5 ticks
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	// Print the tick events
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}

	fmt.Println("Ticker stopped.")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `021_tickers/01_basic_ticker` directory:

```bash
cd go_sample_examples/021_tickers/01_basic_ticker
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see the following output, where the ticker prints a message every second and stops after 5 seconds:

```bash
Tick at 2024-09-14 11:36:20.9222214 +0300 +03 m=+1.000559901
Tick at 2024-09-14 11:36:21.9222214 +0300 +03 m=+2.000559901
Tick at 2024-09-14 11:36:22.9222214 +0300 +03 m=+3.000559901
Tick at 2024-09-14 11:36:23.9222214 +0300 +03 m=+4.000559901
Tick at 2024-09-14 11:36:24.9222214 +0300 +03 m=+5.000559901
fatal error: all goroutines are asleep - deadlock! 
```