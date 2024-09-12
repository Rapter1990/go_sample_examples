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
