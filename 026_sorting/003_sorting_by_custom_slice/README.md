# Go Sample Example - Sorting by Custom Slice

This repository demonstrates how to sort a slice of structs in Go using the built-in `sort.Slice` function. It provides an example of custom sorting based on a specific field, like sorting by age.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers sorting a slice of structs in Go using custom criteria.</li>
  <li>It showcases sorting a slice of `Person` structs by their `Age` field in ascending order using the `sort.Slice` function.</li>
  <li>Sorting slices of structs with custom logic is a common task in Go, and this example illustrates how to implement it efficiently.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"sort"
)

// Define the Person struct with Name and Age fields
type Person struct {
	Name string
	Age  int
}

func main() {
	// Slice of Person structs to sort
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	// Sorting the slice of structs by Age field in ascending order
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	// Printing the sorted slice
	fmt.Println("Sorted by age:", people)
}
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `003_sorting_by_custom_slice` directory:

   ```bash
   cd go_sample_examples/026_sorting/003_sorting_by_custom_slice
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, the output will display the sorted slice of `Person` structs by their age:

```bash
Sorted by age: [{Charlie 20} {Alice 25} {Bob 30}]
```