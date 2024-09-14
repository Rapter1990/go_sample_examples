# Go Sample Example - Resetting a Ticker

This repository demonstrates how to reset a ticker in Go, allowing dynamic changes to the ticker's interval. Tickers trigger actions at regular intervals, and resetting a ticker adjusts the interval while keeping the ticker running.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to reset a ticker to dynamically change the interval during program execution.</li>
  <li>It demonstrates starting a ticker, resetting the ticker's interval, and stopping the ticker after a specified time.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Resetting a Ticker
	// Resetting a ticker allows you to change the ticker's interval dynamically

	// Create a ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(3 * time.Second)

	// Reset the ticker to tick every 2 seconds instead of 1 second
	ticker.Reset(2 * time.Second)

	time.Sleep(6 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped.")

	// The ticker initially ticks every 1 second.
	// After 3 seconds, ticker.Reset(2 * time.Second) changes the interval to 2 seconds
	// The ticker continues ticking at the new interval until it is stopped
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `021_tickers/04_reset_ticker` directory:

```bash
cd go_sample_examples/021_tickers/04_reset_ticker
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Tick at 2024-09-14 11:45:58.073901 +0300 +03 m=+1.000614901
Tick at 2024-09-14 11:45:59.073901 +0300 +03 m=+2.000614901
Tick at 2024-09-14 11:46:00.073901 +0300 +03 m=+3.000614901
Tick at 2024-09-14 11:46:02.0745411 +0300 +03 m=+5.001255001
Tick at 2024-09-14 11:46:04.0745411 +0300 +03 m=+7.001255001
Tick at 2024-09-14 11:46:06.0745411 +0300 +03 m=+9.001255001
Ticker stopped.
```