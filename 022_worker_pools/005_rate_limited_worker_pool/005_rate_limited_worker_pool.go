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
				-	The time.Sleep(500 * time.Millisecond) introduces a delay of 500 milliseconds before processing the next job.
				-	This delay is meant to simulate the effect of rate limiting, where each worker can only process one job every 500 milliseconds.

			Job Processing:
				-	Since there are multiple workers (3 in this case), and each job takes 500 milliseconds to process, the overall rate at which jobs are completed will be influenced by the number of workers and the processing time.

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
