# Go Sample Example - Sorting with Custom Comparator Function

This repository demonstrates how to sort a slice of structs using a custom comparator function in Go. It showcases how to use Go's `sort.Slice` function to sort a collection based on a user-defined condition, providing greater flexibility for sorting complex data structures.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers sorting a slice of structs using a custom comparator in Go.</li>
  <li>It uses the `sort.Slice` function to sort a slice of `Person3` structs by the length of their `Name` field.</li>
  <li>This approach can be adapted to sort slices based on any criteria.</li>
</ul>

## 💻 Code Example

```golang
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
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `007_sorting_with_custom_comparator_function` directory:

   ```bash
   cd go_sample_examples/026_sorting/007_sorting_with_custom_comparator_function
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, the output will display the `Person3` structs sorted by the length of their names:

```bash
Sorted by name length: [{Bob 30} {Alice 25} {Charlie 20}]
```