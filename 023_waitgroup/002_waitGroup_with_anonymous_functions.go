package main

import (
	"fmt"
	"sync"
)

func main() {

	// WaitGroup with Anonymous Functions
	// Demonstrates using a WaitGroup with an anonymous function

	var wg sync.WaitGroup
	n := 5

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d is processing\n", id)
		}(i)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers completed.")
}
