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
	// workers can report errors, and the main function can handle them

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
