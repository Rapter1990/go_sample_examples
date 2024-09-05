package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")
	done <- true // Signal that the work is done
}

func worker1(id int, done chan bool) {
	fmt.Printf("Worker %d: Working...\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d: Done\n", id)
	done <- true
}

func worker2(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Printf("Worker %d: Working...\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d: Done\n", id)
}

func worker3(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d: started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d: finished job %d\n", id, j)
		results <- j * 2
	}
}

func worker4(done chan<- bool) {
	fmt.Println("Worker: Starting work")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker: Done with work")
	done <- true
}

func worker5(jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		fmt.Printf("Processing job %d\n", job)
	}
	done <- true
}

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic Channel Synchronization -> synchronize the completion of a goroutine

	done := make(chan bool)

	go worker(done)

	// Wait for the worker to finish
	<-done
	fmt.Println("Worker has finished.")

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Using Buffered Channels for Synchronization
	// A buffered channel can allow sending a limited number of signals without blocking

	done = make(chan bool, 2)

	for i := 1; i <= 2; i++ {
		go worker1(i, done)
	}

	// Wait for both workers to finish
	<-done
	<-done
	fmt.Println("All workers have finished.")

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// WaitGroup for Synchronization
	// The sync.WaitGroup is a common and more idiomatic way to synchronize multiple goroutines

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker2(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers have finished.")

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Channel Synchronization for Fan-Out / Fan-In Patterns
	// multiple goroutines send results to a single channel

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker3(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// One-Way Channels for Synchronization
	// a channel is only used for sending or receiving, which can help with synchronization

	done = make(chan bool)

	go worker4(done)

	// Wait for the worker to signal completion
	<-done
	fmt.Println("Main: Worker has finished")

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	//  Closing a Channel to Signal Completion

	//  Close a channel to signal to multiple receivers that no more data will be sent

	jobs = make(chan int, 5)
	done = make(chan bool)

	go worker5(jobs, done)

	// Send some jobs to the worker
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for the worker to finish
	<-done
	fmt.Println("All jobs processed.")
	fmt.Println("-----------------------------------------------------------------------------------")

}
