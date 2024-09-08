package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortSlice(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(slice)
	fmt.Println("Sorted slice:", slice)
}

func main() {

	// Use goroutines to sort multiple slices concurrently

	var wg sync.WaitGroup

	slices := [][]int{
		{3, 2, 1},
		{6, 5, 4},
		{9, 8, 7},
	}

	for _, slice := range slices {
		wg.Add(1)
		go sortSlice(slice, &wg)
	}

	wg.Wait()
	fmt.Println("All slices sorted")
}
