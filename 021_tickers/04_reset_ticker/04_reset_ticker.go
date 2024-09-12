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
