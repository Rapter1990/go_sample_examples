# Go Sample Example - Atomic Counter with Load and Store Operations

This repository demonstrates the use of atomic load and store operations in Go to safely read and update a counter in concurrent goroutines. It showcases how to use the `sync/atomic` package to ensure thread-safe manipulation of shared variables.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the use of atomic load and store operations to safely read and update a shared counter variable in concurrent goroutines.</li>
  <li>It demonstrates the use of `atomic.LoadInt64` for reading the current value of a counter and `atomic.StoreInt64` for updating it.</li>
  <li>Multiple goroutines increment the counter, while one goroutine uses load and store operations to read and update the counter concurrently.</li>
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the **025_atomic_counters/004_atomic_counter_with_load_and_store_operations** directory:

```bash
cd go_sample_examples/025_atomic_counters/004_atomic_counter_with_load_and_store_operations
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see an output similar to:

```bash
Current Counter: 0
Updated Counter: 1000
Final Counter: 1000
```
