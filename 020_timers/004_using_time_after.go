package main

import (
	"fmt"
	"time"
)

func main() {

	// Using time.After
	// time.After is a simpler alternative to time.NewTimer for creating a timer that only needs to fire once.
	// It returns a channel that will receive the current time after the specified duration.

	fmt.Println("Waiting for 2 seconds...")

	// Wait for 2 seconds using time.After
	<-time.After(2 * time.Second)

	fmt.Println("2 seconds passed.")
}
