# Go Sample Example - Goroutines & Channels

This repository demonstrates how to use Goroutines, Channels, WaitGroups, Mutex, and the Select statement in Go. It showcases concurrent programming by utilizing these constructs to handle tasks in parallel and communicate between them safely.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic to advanced usage of Goroutines and Channels in Go.</li>
  <li>It demonstrates how to start Goroutines, synchronize them with WaitGroups, use channels for communication, and implement mutex locks to prevent race conditions.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Basic Goroutine function to print numbers
func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}

// Basic Goroutine function to print letters
func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		time.Sleep(700 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}

// Worker function using WaitGroup to synchronize
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// Function to sum two numbers and send result via a channel
func sum(a, b int, c chan int) {
	time.Sleep(2 * time.Second)
	c <- a + b
}

// SafeCounter struct using Mutex for concurrent access
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	sc.value++
	sc.mu.Unlock()
}

func (sc *SafeCounter) GetValue() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.value
}

func main() {

	// Running basic Goroutines
	go printNumbers()
	go printLetters()

	time.Sleep(4 * time.Second)
	fmt.Println("Main function finished")

	// Using Anonymous Function Goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("Inside anonymous Goroutine:", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Main function finished")

	// Synchronizing Goroutines with WaitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished")

	// Goroutines with Channels
	c := make(chan int)
	go sum(3, 4, c)
	result := <-c
	fmt.Println("Sum:", result)

	// Buffered Channels
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Received:", val)
	}

	// Select statement with multiple Goroutines
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Hello from Goroutine 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Hello from Goroutine 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

	// Goroutines with Mutex
	sc := SafeCounter{}
	var wa sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wa.Add(1)
		go func() {
			defer wa.Done()
			sc.Increment()
		}()
	}

	wa.Wait()
	fmt.Println("Final Counter Value:", sc.GetValue())
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `011_goroutine_channel` directory:

```bash
   cd go_sample_examples/011_goroutine_channel
```

4. Run the Go program:

```bash
   go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```
-----------------------------------------------------------------------------------
1
A
2
B
3
4
C
5
D
E
Main function finished
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Inside anonymous Goroutine: 1
Inside anonymous Goroutine: 2
Inside anonymous Goroutine: 3
Main function finished
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Worker 1 starting
Worker 3 starting
Worker 2 starting
Worker 2 done
Worker 3 done
Worker 1 done
All workers finished
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Sum: 7
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Values sent to channel
Received: 1
Received: 2
Received: 3
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Hello from Goroutine 1
Hello from Goroutine 2
All messages received
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
Final Counter Value: 1000
-----------------------------------------------------------------------------------
```
