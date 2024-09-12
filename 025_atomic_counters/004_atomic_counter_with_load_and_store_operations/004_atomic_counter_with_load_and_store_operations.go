package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Atomic Counter with Load and Store Operations
	// Illustrate using atomic load and store operations to ensure safe reading and updating of a counter

	var counter int64 // Shared atomic counter
	var wg sync.WaitGroup

	// Increment counter in goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}

	// Load and store operations
	wg.Add(1)
	go func() {
		defer wg.Done()
		current := atomic.LoadInt64(&counter) // Atomically load the counter value
		fmt.Println("Current Counter:", current)
		atomic.StoreInt64(&counter, current+1000) // Atomically set counter to current + 1000
		fmt.Println("Updated Counter:", atomic.LoadInt64(&counter))
	}()

	wg.Wait()
	fmt.Println("Final Counter:", counter)

}
