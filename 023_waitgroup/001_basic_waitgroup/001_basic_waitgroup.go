package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// Simulate some work with sleep
	// time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// we use a WaitGroup to wait for a set of goroutines to complete their execution before the program exits

	// Basic WaitGroup Example

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers completed.")

}
