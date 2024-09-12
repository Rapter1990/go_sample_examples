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
