package main

import (
	"fmt"
	"sort"
)

func main() {

	// Show how to sort a map by its values

	ages := map[string]int{
		"Alice":   25,
		"Charlie": 20,
		"Bob":     30,
	}

	type kv struct {
		Key   string
		Value int
	}

	var sortedAges []kv
	for k, v := range ages {
		sortedAges = append(sortedAges, kv{k, v})
	}

	sort.Slice(sortedAges, func(i, j int) bool {
		return sortedAges[i].Value < sortedAges[j].Value
	})

	fmt.Println("Sorted by values:")
	for _, kv := range sortedAges {
		fmt.Println(kv.Key, kv.Value)
	}
}
