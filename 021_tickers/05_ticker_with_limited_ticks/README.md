# Go Sample Example - Ticker with Limited Ticks

This repository demonstrates how to use a ticker in Go with a set limit on the number of ticks. A ticker generates events at regular intervals, and this example shows how to stop the ticker after reaching a specified number of ticks.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to create and stop a ticker after a certain number of ticks.</li>
  <li>It demonstrates how to set a maximum number of ticks before stopping the ticker.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Ticker with Limited Ticks
	// Stop the ticker after a certain number of ticks

	// Create a ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)

	// Create a counter for the number of ticks
	count := 0
	maxTicks := 5

	for t := range ticker.C {
		fmt.Println("Tick at", t)
		count++
		if count >= maxTicks {
			ticker.Stop()
			break
		}
	}

	fmt.Println("Ticker stopped after", maxTicks, "ticks.")

	// The ticker is set to stop after 5 ticks
	// The loop counts the ticks and stops the ticker once the limit is reached
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `021_tickers/05_ticker_with_limited_ticks` directory:

```bash
cd go_sample_examples/021_tickers/05_ticker_with_limited_ticks
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Tick at 2024-09-14 11:48:18.9710563 +0300 +03 m=+1.000576301
Tick at 2024-09-14 11:48:19.9710563 +0300 +03 m=+2.000576301
Tick at 2024-09-14 11:48:20.9710563 +0300 +03 m=+3.000576301
Tick at 2024-09-14 11:48:21.9710563 +0300 +03 m=+4.000576301
Tick at 2024-09-14 11:48:22.9710563 +0300 +03 m=+5.000576301
Ticker stopped after 5 ticks.
```