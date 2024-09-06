package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Atomic Counter with Decrement and Compare-And-Swap (CAS)
	// use atomic operations to increment, decrement, and use the compare-and-swap operation

	var counter int32 = 100 // Initialize counter
	var wg sync.WaitGroup

	// Increment goroutine
	wg.Add(1)
	go func() {
		atomic.AddInt32(&counter, 1) // Increment counter
		wg.Done()
	}()

	// Decrement goroutine
	wg.Add(1)
	go func() {
		atomic.AddInt32(&counter, -1) // Decrement counter
		wg.Done()
	}()

	// Compare-And-Swap goroutine
	wg.Add(1)
	go func() {
		// CAS: If counter equals 100, set it to 200
		if atomic.CompareAndSwapInt32(&counter, 100, 200) {
			fmt.Println("CAS successful, new value:", counter)
		} else {
			fmt.Println("CAS failed, current value:", counter)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
