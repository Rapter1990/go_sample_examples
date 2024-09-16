# Go Sample Example - Basic Atomic Counter

This repository demonstrates how to use an atomic counter in Go to safely increment and retrieve values across multiple goroutines using the `sync/atomic` package. It showcases concurrency control with atomic operations, ensuring thread-safe updates to shared variables.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the use of the `sync/atomic` package to create a thread-safe counter.</li>
  <li>It shows how to increment a shared variable concurrently from multiple goroutines without data races.</li>
  <li>The example uses a `WaitGroup` to ensure all goroutines finish before printing the final counter value.</li>
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

	// Basic Atomic Counter Using sync/atomic
	// Use an atomic counter to increment and retrieve values safely across multiple goroutines

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
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the **025_atomic_counters/001_basic_atomic_counter_using_sync_atomic** directory:

```bash
cd go_sample_examples/025_atomic_counters/001_basic_atomic_counter_using_sync_atomic
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see an output similar to:

```bash
Final Counter: 5000
```