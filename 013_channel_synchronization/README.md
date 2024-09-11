# Go Sample Example - Channel Synchronization

This example demonstrates various ways to use channels and goroutines for synchronization in Go. It covers basic channel synchronization, buffered channels, `sync.WaitGroup`, fan-out/fan-in patterns, one-way channels, and signaling completion with closed channels.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers different methods of using Go channels for synchronizing goroutines.</li>
  <li>It demonstrates how to use basic channel synchronization, buffered channels, wait groups, fan-out/fan-in patterns, one-way channels, and closing channels to signal completion.</li>
</ul>

## 💻 Code Example

```go
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

	// Basic Channel Synchronization
	done := make(chan bool)
	go worker(done)
	<-done
	fmt.Println("Worker has finished.")

	fmt.Println("-----------------------------------------------------------------------------------")

	// Using Buffered Channels
	done = make(chan bool, 2)
	for i := 1; i <= 2; i++ {
		go worker1(i, done)
	}
	<-done
	<-done
	fmt.Println("All workers have finished.")

	fmt.Println("-----------------------------------------------------------------------------------")

	// WaitGroup for Synchronization
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker2(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers have finished.")

	fmt.Println("-----------------------------------------------------------------------------------")

	// Channel Synchronization for Fan-Out/Fan-In Patterns
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

	// One-Way Channels
	done = make(chan bool)
	go worker4(done)
	<-done
	fmt.Println("Main: Worker has finished")

	fmt.Println("-----------------------------------------------------------------------------------")

	// Closing a Channel to Signal Completion
	jobs = make(chan int, 5)
	done = make(chan bool)
	go worker5(jobs, done)
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)
	<-done
	fmt.Println("All jobs processed.")
	fmt.Println("-----------------------------------------------------------------------------------")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `013_channel_synchronization` directory:

```bash
   cd go_sample_examples/013_channel_synchronization
```

4. Run the Go program:

```bash
   go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```
-----------------------------------------------------------------------------------
Working...
Done
Worker has finished.
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Worker 2: Working...
Worker 1: Working...
Worker 1: Done
Worker 2: Done
All workers have finished.
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Worker 3: Working...
Worker 1: Working...
Worker 2: Working...
Worker 2: Done
Worker 3: Done
Worker 1: Done
All workers have finished.
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Worker 3: started job 1
Worker 1: started job 2
Worker 2: started job 3
Worker 1: finished job 2
Worker 1: started job 4
Worker 3: finished job 1
Worker 3: started job 5
Worker 2: finished job 3
Result: 4
Result: 2
Result: 6
Worker 3: finished job 5
Result: 10
Worker 1: finished job 4
Result: 8
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Worker: Starting work
Worker: Done with work
Main: Worker has finished
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Processing job 1
Processing job 2
Processing job 3
All jobs processed.
-----------------------------------------------------------------------------------
```
