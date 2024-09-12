package main

import (
	"fmt"
	"sync"
)

func worker1(id int, wg *sync.WaitGroup, errors chan error) {
	defer wg.Done()

	// Simulate an error for worker 2
	if id == 2 {
		errors <- fmt.Errorf("worker %d encountered an error", id)
		return
	}

	fmt.Printf("Worker %d completed successfully\n", id)
}

func main() {

	// WaitGroup with Error Handling
	// we demonstrate using a WaitGroup with error handling in the workers

	var wg sync.WaitGroup
	errors := make(chan error, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker1(i, &wg, errors)
	}

	// Wait for all workers to finish
	wg.Wait()
	close(errors)

	// Check for errors
	for err := range errors {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("All workers completed.")

}
