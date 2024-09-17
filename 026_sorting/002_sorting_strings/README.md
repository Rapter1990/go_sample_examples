# Go Sample Example - Sorting Strings

This repository demonstrates how to sort a slice of strings in Go using the built-in `sort` package. It provides an example of working with string sorting functionality in Go.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers sorting a slice of strings in ascending alphabetical order using the `sort.Strings` function.</li>
  <li>Sorting strings is a common task in many applications, and Go provides a simple way to do this using the `sort` package.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Slice of strings to sort
	strings := []string{"banana", "apple", "cherry"}
	
	// Sorting the slice in ascending order
	sort.Strings(strings)
	
	// Printing the sorted slice
	fmt.Println("Sorted strings:", strings)
}
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `002_sorting_strings` directory:

   ```bash
   cd go_sample_examples/026_sorting/002_sorting_strings
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, the output will display the sorted strings:

```bash
Sorted strings: [apple banana cherry]
```