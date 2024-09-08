package main

import (
	"fmt"
	"sort"
)

type Person2 struct {
	Name string
	Age  int
}

type ByAge []Person2

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {

	// Show how to implement the sort.Interface to sort a custom type

	people := []Person2{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	sort.Sort(ByAge(people))
	fmt.Println("Sorted by age using sort.Interface:", people)
}
