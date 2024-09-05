package main

import (
	"fmt"
	"time"
)

func main() {

	// Stopping a Timer
	// You can stop a timer before it fires using the Stop method.
	//If the timer is stopped before it fires, the program will not receive any value from the timer's channel

	// Create a timer that will fire after 3 seconds
	timer := time.NewTimer(3 * time.Second)

	// Start a goroutine to stop the timer after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		stop := timer.Stop()
		if stop {
			fmt.Println("Timer stopped before it fired.")
		}
	}()

	// Wait for the timer to fire
	<-timer.C
	fmt.Println("This will not be printed if the timer is stopped.")
}
