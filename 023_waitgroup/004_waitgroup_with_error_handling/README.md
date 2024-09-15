# Go Sample Example - WaitGroup with Error Handling

This repository demonstrates how to use Go's `sync.WaitGroup` to synchronize goroutines and handle errors that may occur during their execution. The example showcases how to capture and process errors from multiple goroutines using a channel.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example demonstrates using a `sync.WaitGroup` with error handling in Go.</li>
  <li>It includes a worker function that simulates tasks, where one worker intentionally encounters an error.</li>
  <li>The example shows how to propagate errors back to the main goroutine using a channel and handle them appropriately after all workers finish execution.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
)

func worker1(id int, wg *sync.WaitGroup, errors chan error) {
	defer wg.Done()

	// Simulate an error for worker 2
	if id == 2 {
		errors <- fmt.Errorf("worker %d encountered an error", id)
		return
	}

	fmt.Printf("Worker %d completed successfully\n", id)
}

func main() {

	// WaitGroup with Error Handling
	// we demonstrate using a WaitGroup with error handling in the workers

	var wg sync.WaitGroup
	errors := make(chan error, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker1(i, &wg, errors)
	}

	// Wait for all workers to finish
	wg.Wait()
	close(errors)

	// Check for errors
	for err := range errors {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

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
cd go_sample_examples/023_waitgroup/004_waitgroup_with_error_handling
```

4. Run the Go program:

```bash
go run 004_waitgroup_with_error_handling.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Worker 3 completed successfully
Worker 1 completed successfully     
Error: worker 2 encountered an error
All workers completed.
```