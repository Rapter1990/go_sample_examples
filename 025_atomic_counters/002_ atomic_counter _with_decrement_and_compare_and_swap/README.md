# Go Sample Example - Atomic Counter with Decrement and Compare-And-Swap (CAS)

This repository demonstrates how to use atomic operations in Go, including incrementing, decrementing, and using the Compare-And-Swap (CAS) function to manipulate shared variables safely in concurrent programs.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the use of atomic operations to perform safe updates on shared variables across multiple goroutines.</li>
  <li>It showcases how to increment and decrement a counter using `atomic.AddInt32`, and how to use the `Compare-And-Swap` (CAS) function for conditional updates.</li>
  <li>A `WaitGroup` is used to ensure that all goroutines complete before printing the final value of the counter.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Atomic Counter with Decrement and Compare-And-Swap (CAS)
	// Use atomic operations to increment, decrement, and use the compare-and-swap operation

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
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the **025_atomic_counters/002_atomic_counter_with_decrement_and_compare_and_swap** directory:

```bash
cd go_sample_examples/025_atomic_counters/002_atomic_counter_with_decrement_and_compare_and_swap
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see an output similar to:

```bash
CAS successful, new value: 200
Final Counter: 200
```