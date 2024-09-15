# Go Sample Example - WaitGroup with Multiple Waits

This repository demonstrates how to use Go's `sync.WaitGroup` to synchronize different sets of goroutines at various stages of execution. The example showcases how to coordinate goroutines across multiple stages using the same WaitGroup for synchronization.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example demonstrates using a single `sync.WaitGroup` to manage synchronization between multiple stages of goroutine execution.</li>
  <li>It covers the use of a WaitGroup in two separate phases, waiting for a set of goroutines to complete before moving on to the next stage.</li>
  <li>The goroutines simulate tasks using sleep timers to mimic real processing time.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func stageOne(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Stage one: Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Stage one: Worker %d done\n", id)
}

func stageTwo(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Stage two: Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Stage two: Worker %d done\n", id)
}

func main() {

	// WaitGroup with Multiple Waits
	// we use a WaitGroup to wait for different sets of goroutines at different stages of execution.

	var wg sync.WaitGroup

	// Stage one
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go stageOne(i, &wg)
	}
	wg.Wait() // Wait for all workers in stage one to complete
	fmt.Println("Stage one completed.")

	// Stage two
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go stageTwo(i, &wg)
	}
	wg.Wait() // Wait for all workers in stage two to complete
	fmt.Println("Stage two completed.")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `023_waitgroup` directory:

```bash
cd go_sample_examples/023_waitgroup/003_waitgroup_with_multiple_waits
```

4. Run the Go program:

```bash
go run 003_waitgroup_with_multiple_waits.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Stage one: Worker 3 starting
Stage one: Worker 1 starting
Stage one: Worker 2 starting
Stage one: Worker 2 done
Stage one: Worker 1 done    
Stage one: Worker 3 done    
Stage one completed.        
Stage two: Worker 3 starting
Stage two: Worker 1 starting
Stage two: Worker 2 starting
Stage two: Worker 2 done
Stage two: Worker 3 done
Stage two: Worker 1 done
Stage two completed.
```