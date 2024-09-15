# Go Sample Example - Worker Pools

This repository demonstrates how to implement and use a basic worker pool in Go. Worker pools allow for concurrent task processing by distributing jobs across multiple workers, making the task management more efficient.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the basic concept of worker pools in Go, including how to assign tasks (jobs) to multiple workers concurrently.</li>
  <li>It showcases a simple worker pool where each worker processes a task from a shared channel.</li>
  <li>Workers simulate task processing by doubling the value of the job and returning the result.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function to process jobs
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		results <- j * 2
	}
}

func main() {

	// In Go, worker pools are groups of workers dedicated to processing tasks.
	// Jobs are sent to a channel and handled by these workers, allowing tasks to be distributed and workload balanced efficiently.

	// Basic Worker Pool
	// Basic worker pool where multiple workers process tasks from a shared channel.

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `001_basic_worker_pool` directory:

   ```bash
   cd go_sample_examples/022_worker_pools/001_basic_worker_pool
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
Worker 3 processing job 1
Worker 1 processing job 2
Worker 2 processing job 3
Worker 2 processing job 4
Worker 3 processing job 5
Result: 6
Result: 2 
Result: 4 
Result: 10
Result: 8
```