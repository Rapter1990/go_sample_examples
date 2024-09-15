# Go Sample Example - WaitGroup

This repository demonstrates how to use WaitGroups in Go to synchronize multiple goroutines. A WaitGroup helps wait for a collection of goroutines to complete their tasks before the main program continues execution or terminates.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the basic usage of WaitGroups in Go.</li>
  <li>It shows how to start multiple goroutines and wait for their completion using the `sync.WaitGroup`.</li>
  <li>The example includes a simple worker function that simulates a task and demonstrates how to use WaitGroups to synchronize the completion of multiple workers.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// Simulate some work with sleep
	// time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// we use a WaitGroup to wait for a set of goroutines to complete their execution before the program exits

	// Basic WaitGroup Example

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
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
cd go_sample_examples/023_waitgroup/001_basic_waitgroup
```

4. Run the Go program:

```bash
go run 001_basic_waitgroup.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Worker 3 starting
Worker 3 done         
Worker 1 starting     
Worker 1 done         
Worker 2 starting     
Worker 2 done         
All workers completed.
```