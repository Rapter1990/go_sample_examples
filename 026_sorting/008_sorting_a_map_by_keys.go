package main

import (
	"fmt"
	"sort"
)

func main() {

	// Show how to sort a map by its keys

	ages := map[string]int{
		"Alice":   25,
		"Charlie": 20,
		"Bob":     30,
	}

	keys := make([]string, 0, len(ages))
	for k := range ages {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Println("Sorted by keys:")
	for _, k := range keys {
		fmt.Println(k, ages[k])
	}
}
