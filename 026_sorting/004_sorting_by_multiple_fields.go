package main

import (
	"fmt"
	"sort"
)

type Person1 struct {
	Name string
	Age  int
}

func main() {

	// Demonstrate how to sort by multiple fields. First by Age, and if the Age is the same, then by Name

	people := []Person1{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
		{"Dave", 25},
	}

	sort.Slice(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Name < people[j].Name
		}
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted by age and name:", people)
}
