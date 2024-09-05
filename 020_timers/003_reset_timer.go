package main

import (
	"fmt"
	"time"
)

func main() {

	// Resetting a Timer
	// The Reset method is used to reset a timer to a new duration.
	// If the timer had already fired, it will restart with the new duration

	// Create a timer that will fire after 2 seconds
	timer := time.NewTimer(2 * time.Second)

	// Start a goroutine to reset the timer after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		timer.Reset(4 * time.Second)
		fmt.Println("Timer reset to 4 seconds.")
	}()

	// Wait for the timer to fire
	<-timer.C
	fmt.Println("Timer fired!")
}
