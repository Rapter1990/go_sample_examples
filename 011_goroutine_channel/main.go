package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		time.Sleep(700 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify that the Goroutine is done

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func sum(a, b int, c chan int) {
	time.Sleep(2 * time.Second)
	c <- a + b // Send the sum to the channel
}

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

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic Goroutine Example

	// Starting a Goroutine
	go printNumbers()
	go printLetters()

	// Giving time for Goroutines to finish
	time.Sleep(4 * time.Second)
	fmt.Println("Main function finished")

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Using Anonymous Functions as Goroutines

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("Inside anonymous Goroutine:", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Main function finished")

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Synchronizing Goroutines with WaitGroup

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add a Goroutine to the WaitGroup
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all Goroutines to finish
	fmt.Println("All workers finished")

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Goroutines with Channels

	c := make(chan int) // Channels can be used to communicate between Goroutines safely.

	go sum(3, 4, c)

	result := <-c // Receive the result from the channel
	fmt.Println("Sum:", result)

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Buffered Channels with Goroutines
	// Buffered channels allow sending multiple values without blocking

	ch := make(chan int, 3)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("Values sent to channel")
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Received:", val)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Select Statement with Goroutines
	// The select statement allows waiting on multiple channel operations

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

	fmt.Println("All messages received")

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Goroutines with Mutex
	// A sync.Mutex can be used to protect shared data from concurrent access.

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
	fmt.Println("-----------------------------------------------------------------------------------")

}
