package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(numbers)
	fmt.Println("Sorted integers:", numbers)
}
