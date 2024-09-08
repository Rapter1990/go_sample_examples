package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// Show how to sort a slice of structs by a custom field, like sorting by Age

	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted by age:", people)
}
