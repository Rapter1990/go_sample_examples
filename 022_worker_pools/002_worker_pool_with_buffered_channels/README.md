# Go Sample Example - Worker Pool with Buffered Channels

This repository demonstrates how to implement a worker pool in Go using buffered channels. Buffered channels allow workers to process jobs more efficiently by reducing blocking when sending jobs to workers.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the use of worker pools with buffered channels in Go to improve task distribution and minimize blocking.</li>
  <li>It showcases how buffered channels enable jobs to be queued before workers process them, improving throughput and reducing idle time.</li>
  <li>Each worker processes jobs concurrently, multiplying the job by 2 and returning the result.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
)

func bufferedWorker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		results <- j * 2
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {

	// Worker Pool with Buffered Channels
	// How to use buffered channels to improve performance by reducing blocking when sending jobs to workers.

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go bufferedWorker(w, jobs, results, &wg)
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

3. Navigate to the `002_worker_pool_with_buffered_channels` directory:

   ```bash
   cd go_sample_examples/022_worker_pools/002_worker_pool_with_buffered_channels
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
Worker 3 started job 1
Worker 3 finished job 1
Worker 3 started job 4 
Worker 3 finished job 4
Worker 3 started job 5 
Worker 3 finished job 5
Worker 2 started job 3 
Worker 2 finished job 3
Worker 1 started job 2 
Worker 1 finished job 2
Result: 2              
Result: 8              
Result: 10             
Result: 6              
Result: 4
```