# Go Sample Example - Channel Directions

This repository demonstrates the usage of directional channels in Go, including send-only, receive-only, bidirectional channels, and pipeline patterns. It showcases how channels can be used to control the flow of data and synchronize between goroutines.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic channel directions, enforcing send-only and receive-only channels in Go.</li>
  <li>It includes channel operations such as sending and receiving, using channels in goroutines, pipelines, and buffered channels.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

// sendOnly sends an integer to the channel (send-only)
func sendOnly(ch chan<- int, value int) {
	ch <- value
}

// receiveOnly receives an integer from the channel (receive-only)
func receiveOnly(ch <-chan int) int {
	return <-ch
}

func sendNumbers(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func printNumbers(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

// Stage 1: Generate numbers
func generateNumbers(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out)
}

// Stage 2: Square numbers
func squareNumbers(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

// Stage 3: Print numbers
func printNumbers1(in <-chan int) {
	for num := range in {
		fmt.Println("Square:", num)
	}
}

func pingPong(pings <-chan string, pongs chan<- string) {
	for {
		select {
		case msg := <-pings:
			fmt.Println("Received ping:", msg)
			pongs <- "pong"
		}
	}
}

func startGenerator() chan int { // Return chan int instead of <-chan int
	ch := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func bufferedSender(ch chan<- string) { // Channel for sending strings
	messages := []string{"hello", "world", "go", "channel"}
	for _, msg := range messages {
		ch <- msg
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func bufferedReceiver(ch <-chan string) { // Channel for receiving strings
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic Channel Directions
	ch := make(chan int)

	go sendOnly(ch, 42)

	result := receiveOnly(ch)
	fmt.Println("Received:", result)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Using Directional Channels with Goroutines
	ch = make(chan int)

	go sendNumbers(ch)
	printNumbers(ch)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Directional Channels in a Pipeline
	ch1 := make(chan int)
	ch2 := make(chan int)

	go generateNumbers(ch1)
	go squareNumbers(ch1, ch2)
	printNumbers(ch2)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Bidirectional Channels with Select Statement
	pings := make(chan string)
	pongs := make(chan string)

	go pingPong(pings, pongs)

	pings <- "ping"
	fmt.Println("Sent ping")

	fmt.Println("Received pong:", <-pongs)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Returning Channels from Functions
	ch = startGenerator()

	for num := range ch {
		fmt.Println("Received:", num)
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Buffered Channels with Directional Usage
	stringChannel := make(chan string, 3)

	go bufferedSender(stringChannel)
	bufferedReceiver(stringChannel)

	fmt.Println("-----------------------------------------------------------------------------------")

}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `014_channel_directions` directory:

```bash
cd go_sample_examples/014_channel_directions
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
--------------------------------------------------------------------------------
---
Received: 42
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Received: 1
Received: 9
Received: 16
Received: 25
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Received ping: ping
Received: 9
Received: 16
Received: 25
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Received ping: ping
Sent ping
Received pong: pong
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Received: 1
Received: 2
Received: 3
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Received: hello
Received: world
Received: go
Received: channel
-----------------------------------------------------------------------------------
```
