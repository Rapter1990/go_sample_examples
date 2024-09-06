package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Basic Atomic Counter Using sync/atomic
	// use an atomic counter to increment and retrieve values safely across multiple goroutines

	var counter int32 // Shared counter variable
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1) // Atomically increment the counter
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)

}
