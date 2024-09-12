package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Atomic Flag Example Using sync/atomic
	// Demonstrate using an atomic boolean flag to control access to a critical section

	var flag int32 // Shared atomic flag
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			if atomic.CompareAndSwapInt32(&flag, 0, 1) { // Try to set flag to 1
				fmt.Printf("Goroutine %d entered critical section\n", id)
				atomic.StoreInt32(&flag, 0) // Reset flag to 0
			} else {
				fmt.Printf("Goroutine %d did not enter critical section\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines done.")
}
