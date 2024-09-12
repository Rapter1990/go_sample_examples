package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	// Atomic Pointer
	// demonstrate how to atomically load and store a pointer value

	var data atomic.Value

	// Store an initial value
	data.Store("Initial Value")

	// Load and print the current value
	fmt.Println("Current Value:", data.Load())

	// Start a goroutine to change the value
	go func() {
		data.Store("New Value")
	}()

	// Load and print the new value
	fmt.Println("New Value:", data.Load())
}
