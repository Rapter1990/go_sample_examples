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
