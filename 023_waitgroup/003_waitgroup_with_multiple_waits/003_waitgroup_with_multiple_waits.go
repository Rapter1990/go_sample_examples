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
