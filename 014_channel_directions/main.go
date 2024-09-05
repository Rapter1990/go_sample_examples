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
	// function that can only send data to a channel

	ch := make(chan int)

	go sendOnly(ch, 42)

	result := receiveOnly(ch)
	fmt.Println("Received:", result)

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Using Directional Channels with Goroutines
	// Directional channels are often used with goroutines to enforce that certain channels are only used for sending or receiving

	ch = make(chan int)

	go sendNumbers(ch)
	printNumbers(ch)

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Directional Channels in a Pipeline
	// Pipelines are a common pattern in Go where data flows through multiple stages, each stage represented by a function. Directional channels help enforce the flow of data

	ch1 := make(chan int)
	ch2 := make(chan int)

	go generateNumbers(ch1)
	go squareNumbers(ch1, ch2)
	printNumbers(ch2)

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Bidirectional Channels with Select Statement
	// Handle both sending and receiving in a function, particularly when using the select statement

	pings := make(chan string)
	pongs := make(chan string)

	go pingPong(pings, pongs)

	pings <- "ping"
	fmt.Println("Sent ping")

	fmt.Println("Received pong:", <-pongs)

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Returning Channels from Functions
	// Channels can be returned from functions, and you can enforce whether the returned channel is directional

	ch = startGenerator()

	for num := range ch {
		fmt.Println("Received:", num)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Buffered Channels with Directional Usage
	stringChannel := make(chan string, 3) // Creating a buffered channel for strings

	go bufferedSender(stringChannel) // Launch the sender goroutine
	bufferedReceiver(stringChannel)  // Receive and print messages

	fmt.Println("-----------------------------------------------------------------------------------")

}
