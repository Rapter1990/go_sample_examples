# Go Sample Example - Basic Sorting with Integers

This repository demonstrates how to sort a slice of integers in Go using the built-in `sort` package. It provides a basic example of how to work with sorting functionality in Go.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers sorting a slice of integers in ascending order using the `sort.Ints` function.</li>
  <li>Sorting is an essential task in programming, and Go provides convenient tools for it in the `sort` package.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Slice of integers to sort
	numbers := []int{5, 2, 6, 3, 1, 4}
	
	// Sorting the slice in ascending order
	sort.Ints(numbers)
	
	// Printing the sorted slice
	fmt.Println("Sorted integers:", numbers)
}
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `001_basic_sorting_with_integers` directory:

   ```bash
   cd go_sample_examples/026_sorting/001_basic_sorting_with_integers
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, the output will display the sorted integers:

```bash
Sorted integers: [1 2 3 4 5 6]
```