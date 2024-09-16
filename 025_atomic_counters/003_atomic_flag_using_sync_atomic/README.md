# Go Sample Example - Atomic Flag Using sync/atomic

This repository demonstrates the use of an atomic flag in Go to control access to a critical section. It showcases how to use the `sync/atomic` package to implement an atomic boolean flag for concurrent programming.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the use of an atomic flag to manage access to a shared resource in concurrent goroutines.</li>
  <li>It demonstrates how to use the `atomic.CompareAndSwapInt32` function to set and reset the flag, allowing only one goroutine to enter the critical section at a time.</li>
  <li>Multiple goroutines attempt to access a critical section, but only one succeeds at a time by modifying the atomic flag.</li>
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
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the **025_atomic_counters/003_atomic_flag_using_sync_atomic** directory:

```bash
cd go_sample_examples/025_atomic_counters/003_atomic_flag_using_sync_atomic
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see an output similar to:

```bash
Goroutine 3 did not enter critical section
Goroutine 0 entered critical section      
Goroutine 4 did not enter critical section
Goroutine 2 did not enter critical section
Goroutine 1 did not enter critical section
All goroutines done.
```