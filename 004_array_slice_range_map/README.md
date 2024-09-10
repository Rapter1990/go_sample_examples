# Go Sample Example - Array, Slice, Range, Map

This repository contains a Go (Golang) program that demonstrates the use of arrays, slices, ranges, and maps. The example illustrates how to declare, initialize, and manipulate these data structures in Go.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the basics of arrays, including how to declare and initialize them, and how to work with arrays of different lengths.</li>
  <li>It introduces slices, including slice creation, initialization, and modifications, as well as the use of the `make` function for slice management.</li>
  <li>The example demonstrates the use of the `range` keyword to iterate over arrays and slices, perform operations like summing elements and modifying slice elements.</li>
  <li>It also covers maps, including how to declare, initialize, add, update, delete key-value pairs, and iterate over maps using `range`.</li>
</ul>

## 💻 Code Example

```golang
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:
   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```
3. Navigate to the `004_array_slice_range_map` directory:
   ```bash
   cd go_sample_examples/004_array_slice_range_map
   ```
4. Run the Go program:
   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
-----------------------------------------------------------------------------------
[1 2 0 0 0]
-----------------------------------------------------------------------------------
-----------------------------------------------------------------------------------
[apple banana cherry]
-----------------------------------------------------------------------------------
[1 2 3 4]
-----------------------------------------------------------------------------------
Array elements:
Index 0: 10
Index 1: 20
Index 2: 30
Index 3: 40
Index 4: 50
-----------------------------------------------------------------------------------

arr3 elements:
Index 0: apple
Index 1: banana
Index 2: cherry
-----------------------------------------------------------------------------------

Slice elements:
Index 0: Go
Index 1: Python
Index 2: JavaScript
-----------------------------------------------------------------------------------

Sum of array elements: 150
-----------------------------------------------------------------------------------

Modified slice4 elements: [200 400 600]
-----------------------------------------------------------------------------------

Map elements:
Key: apple, Value: 10
Key: banana, Value: 25
Key: date, Value: 40
-----------------------------------------------------------------------------------

Value for 'apple': 10
-----------------------------------------------------------------------------------
```