package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"banana", "apple", "cherry"}
	sort.Strings(strings)
	fmt.Println("Sorted strings:", strings)
}
