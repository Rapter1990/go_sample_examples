# Go Sample Example - WaitGroup with Anonymous Functions

This repository demonstrates how to use a Go WaitGroup in conjunction with anonymous functions to synchronize multiple goroutines. It showcases how to launch goroutines with anonymous functions and wait for their completion using a `sync.WaitGroup`.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example demonstrates the use of WaitGroups with anonymous functions in Go.</li>
  <li>It covers launching multiple goroutines, each using an anonymous function to simulate a task, and synchronizing their completion using a `sync.WaitGroup`.</li>
  <li>The example shows how to pass arguments to anonymous functions within goroutines.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
)

func main() {

	// WaitGroup with Anonymous Functions
	// Demonstrates using a WaitGroup with an anonymous function

	var wg sync.WaitGroup
	n := 5

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d is processing\n", id)
		}(i)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers completed.")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `023_waitgroup` directory:

```bash
cd go_sample_examples/023_waitgroup/002_waitGroup_with_anonymous_functions
```

4. Run the Go program:

```bash
go run 002_waitGroup_with_anonymous_functions.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Worker 3 is processing
Worker 2 is processing
Worker 5 is processing
Worker 4 is processing
Worker 1 is processing
All workers completed.
```