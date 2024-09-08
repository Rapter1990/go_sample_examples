package main

import (
	"fmt"
	"sort"
)

type Person3 struct {
	Name string
	Age  int
}

func main() {

	// Sort a slice of structs with a custom comparator, which provides greater flexibility

	people := []Person3{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	// Custom sort by Name length
	sort.Slice(people, func(i, j int) bool {
		return len(people[i].Name) < len(people[j].Name)
	})

	fmt.Println("Sorted by name length:", people)
}
