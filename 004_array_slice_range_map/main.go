package main

import (
	"fmt"
)

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Arrays
	var arr [5]int // Declares an array of 5 integers
	arr[0] = 1     // Assigning a value to the first element
	arr[1] = 2     // Assigning a value to the second element
	fmt.Println(arr)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Array initialization with values
	arr2 := [5]int{10, 20, 30, 40, 50}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Array with inferred length
	arr3 := [...]string{"apple", "banana", "cherry"}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Slices
	var slice []int          // Declares a slice of integers (initially nil)
	slice = append(slice, 1) // Appending elements to the slice
	slice = append(slice, 2, 3, 4)

	// Slice initialization with values
	slice2 := []string{"Go", "Python", "JavaScript"}

	// Slice from an array
	slice3 := arr2[1:4] // Creates a slice from arr2 with elements 20, 30, 40

	// Modifying the slice modifies the underlying array
	slice3[0] = 25 // it can also modify array[1]

	// Slice with make (length 3, capacity 5)
	slice4 := make([]int, 3, 5)
	slice4[0] = 100
	slice4[1] = 200
	slice4[2] = 300

	fmt.Println("-----------------------------------------------------------------------------------")

	// Range with array
	fmt.Println("Array elements:")
	for i, v := range arr2 {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Range with arr3
	fmt.Println("\narr3 elements:")
	for i, v := range arr3 {
		fmt.Printf("Index %d: %s\n", i, v)
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Range with slice
	fmt.Println("\nSlice elements:")
	for i, v := range slice2 {
		fmt.Printf("Index %d: %s\n", i, v)
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Range to sum elements of an array
	sum := 0
	for _, v := range arr2 {
		sum += v
	}
	fmt.Println("\nSum of array elements:", sum)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Range to modify slice elements
	for i := range slice4 {
		slice4[i] *= 2
	}
	fmt.Println("\nModified slice4 elements:", slice4)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Map examples
	// Declare and initialize a map
	m := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}

	// Add a new key-value pair to the map
	m["date"] = 40

	// Update an existing key-value pair
	m["banana"] = 25

	// Delete a key-value pair
	delete(m, "cherry")

	fmt.Println("-----------------------------------------------------------------------------------")

	// Iterate over the map using range
	fmt.Println("\nMap elements:")
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Accessing a value for a specific key
	if value, ok := m["apple"]; ok {
		fmt.Println("\nValue for 'apple':", value)
	} else {
		fmt.Println("\n'apple' not found in the map")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

}
