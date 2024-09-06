package main

import (
	"fmt"
	"sync"
	"time"
)

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
