# Go Sample Example - Dynamic Worker Pool

This repository demonstrates a dynamic worker pool in Go, where the number of workers can be adjusted dynamically based on the workload. It showcases how to implement worker pools and efficiently process tasks in parallel using goroutines and channels.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the implementation of a dynamic worker pool in Go.</li>
  <li>It shows how to spawn workers, dynamically adjust the number of workers based on workload, and handle job processing efficiently.</li>
  <li>The code simulates work using `time.Sleep` and demonstrates how results are returned from workers.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Function for dynamic workers that process jobs and send results
func dynamicWorker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		results <- j * 2
	}
}

func main() {

	// Dynamic Worker Pool
	// the number of workers can be adjusted dynamically based on the workload

	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	numWorkers := 3

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go dynamicWorker(w, jobs, results, &wg)
	}

	// Adjust number of workers based on workload
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Increasing number of workers...")
		for w := numWorkers + 1; w <= numWorkers+2; w++ {
			wg.Add(1)
			go dynamicWorker(w, jobs, results, &wg)
		}
	}()

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

3. Navigate to the `004_dynamic_worker_pool` directory:

```bash
cd go_sample_examples/022_worker_pools/004_dynamic_worker_pool
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see output similar to the following:

```
Worker 3 processing job 3
Worker 2 processing job 2
Worker 1 processing job 1
Worker 1 processing job 4
Worker 3 processing job 5
Worker 2 processing job 6
Increasing number of workers...
Worker 5 processing job 7
Worker 4 processing job 8
Worker 2 processing job 9
Worker 1 processing job 10
Result: 2
Result: 6
Result: 4
Result: 12
Result: 8
Result: 10
Result: 16
Result: 14
Result: 18
Result: 20
```