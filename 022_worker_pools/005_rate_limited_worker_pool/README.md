# Go Sample Example - Rate-Limited Worker Pool

This repository demonstrates a rate-limited worker pool in Go. The example simulates rate-limiting in a worker pool, where each worker can only process one job every 500 milliseconds. It showcases how to implement rate limiting for concurrent job processing.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the implementation of a rate-limited worker pool in Go.</li>
  <li>It demonstrates how to simulate rate limiting by introducing a delay in worker processing.</li>
  <li>Each worker in the pool processes jobs with a controlled delay, mimicking real-world rate-limiting scenarios.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"sync"
	"time"
)

func rateLimitedWorker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(500 * time.Millisecond) // Simulate work and rate limiting
		results <- j * 2
	}
}

func main() {

	// Rate-Limited Worker Pool
	// This example limits the rate at which jobs are processed by the workers

	/*
		How Rate Limiting Works Here:
			Rate Limiting Simulation:
				- The time.Sleep(500 * time.Millisecond) introduces a delay of 500 milliseconds before processing the next job.
				- This delay simulates rate limiting, where each worker can only process one job every 500 milliseconds.

			Job Processing:
				- With multiple workers (3 in this case), each job takes 500 milliseconds to process.
				- The overall rate of job completion depends on the number of workers and the processing time.
	*/

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go rateLimitedWorker(w, jobs, results, &wg)
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

3. Navigate to the `005_rate_limited_worker_pool` directory:

   ```bash
   cd go_sample_examples/022_worker_pools/005_rate_limited_worker_pool
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
Worker 3 processing job 2
Worker 1 processing job 1
Worker 2 processing job 3
Worker 3 processing job 4
Worker 1 processing job 5
Result: 4
Result: 2 
Result: 6 
Result: 10
Result: 8
```