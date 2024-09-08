package main

import (
	"fmt"
	"sort"
)

func main() {

	// Demonstrates how to sort in reverse order

	numbers := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println("Reverse sorted integers:", numbers)
}
