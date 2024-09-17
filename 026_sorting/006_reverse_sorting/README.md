# Go Sample Example - Reverse Sorting

This repository demonstrates how to sort a slice of integers in reverse order using Go's `sort` package. It showcases how to apply the `sort.Reverse` function to an integer slice to sort it in descending order.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers sorting a slice of integers in reverse order in Go.</li>
  <li>It utilizes the `sort.Reverse` function along with `sort.IntSlice` to achieve reverse sorting of integers.</li>
  <li>This is useful for sorting any numerical data in descending order.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"sort"
)

func main() {

	// Demonstrates how to sort in reverse order

	numbers := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println("Reverse sorted integers:", numbers)
}
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `006_reverse_sorting` directory:

   ```bash
   cd go_sample_examples/026_sorting/006_reverse_sorting
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, the output will display the integers sorted in reverse order:

```bash
Reverse sorted integers: [6 5 4 3 2 1]
```