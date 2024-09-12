package main

import (
	"fmt"
	"time"
)

func main() {

	// Basic Example: Simple Timer
	// Create a basic timer that fires after 2 seconds

	// Create a timer that will fire after 2 seconds
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for the timer to fire...")

	// Block until the timer's channel sends a value
	<-timer.C

	fmt.Println("Timer fired!")
}
