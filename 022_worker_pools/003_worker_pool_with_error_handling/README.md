# Go Sample Example - Worker Pool with Error Handling

This repository demonstrates how to implement a worker pool in Go with error handling. The example showcases how to distribute jobs to multiple workers while managing errors for specific jobs, allowing error reporting and handling in a concurrent environment.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the implementation of a worker pool where each worker can handle jobs and report errors if any issues occur during the job execution.</li>
  <li>It includes the use of Go channels for job distribution and result collection, as well as synchronization with `sync.WaitGroup` to ensure all workers complete before results are processed.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID    int
	Error error
}

func workerWithErrorHandling(id int, jobs <-chan Job, results chan<- Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j.ID)
		// Simulate error handling
		if j.ID%2 == 0 {
			j.Error = fmt.Errorf("error processing job %d", j.ID)
		}
		results <- j
	}
}

func main() {

	// Worker Pool with Error Handling
	// Workers can report errors, and the main function can handle them

	const numJobs = 5
	jobs := make(chan Job, numJobs)
	results := make(chan Job, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go workerWithErrorHandling(w, jobs, results, &wg)
	}

	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Job %d failed with error: %v\n", result.ID, result.Error)
		} else {
			fmt.Printf("Job %d completed successfully\n", result.ID)
		}
	}
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `worker_pool_with_error_handling` directory:

```bash
cd go_sample_examples/022_worker_pools/003_worker_pool_with_error_handling
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see the following output (job results will vary based on job ID):

```
Worker 3 processing job 1
Worker 3 processing job 4                      
Worker 3 processing job 5                      
Worker 1 processing job 2                      
Worker 2 processing job 3                      
Job 1 completed successfully                   
Job 4 failed with error: error processing job 4
Job 5 completed successfully                   
Job 2 failed with error: error processing job 2
Job 3 completed successfully
```