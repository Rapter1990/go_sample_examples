# Go Sample Example - Ticker with Select

This repository demonstrates how to use a ticker with a `select` statement in Go. Tickers trigger actions at regular intervals, and combining them with `select` allows you to handle multiple concurrent events efficiently.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to use a ticker inside a `select` statement to manage multiple concurrent events.</li>
  <li>It demonstrates starting a ticker, listening for tick events, and stopping the ticker after receiving a stop signal through a channel.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Using Ticker with Select
	// Use a ticker inside a select statement to manage multiple concurrent events

	ticker := time.NewTicker(1 * time.Second)
	stop := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		stop <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		case <-stop:
			fmt.Println("Stop signal received.")
			ticker.Stop()
			return
		}
	}

	// The select statement listens for ticks and a stop signal
	// When the stop channel receives a signal, the ticker is stopped, and the program exits the loop

}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `021_tickers/03_ticker_with_select` directory:

```bash
cd go_sample_examples/021_tickers/03_ticker_with_select
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Tick at 2024-09-14 11:41:28.3952744 +0300 +03 m=+1.000525401
Tick at 2024-09-14 11:41:29.3952744 +0300 +03 m=+2.000525401
Tick at 2024-09-14 11:41:30.3952744 +0300 +03 m=+3.000525401
Stop signal received.
```