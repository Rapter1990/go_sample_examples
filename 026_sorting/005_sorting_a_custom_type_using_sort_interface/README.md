# Go Sample Example - Sorting a Custom Type Using sort.Interface

This repository demonstrates how to sort a custom type in Go using the `sort.Interface`. It provides an example of defining a custom type `Person2` and implementing the `sort.Interface` methods (`Len`, `Swap`, and `Less`) to sort by a specific field, in this case, `Age`.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers sorting a custom type by implementing the `sort.Interface` in Go.</li>
  <li>The custom type `Person2` is defined with `Name` and `Age` fields, and the sort order is determined by the `Age` field.</li>
  <li>This is useful for sorting custom types in Go based on one or more fields.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"sort"
)

// Define the Person2 struct with Name and Age fields
type Person2 struct {
	Name string
	Age  int
}

// Define ByAge type, which is a slice of Person2
type ByAge []Person2

// Implementing the sort.Interface for ByAge
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {

	// Slice of Person2 structs to sort
	people := []Person2{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	// Sorting the slice by age using the sort.Interface
	sort.Sort(ByAge(people))

	// Printing the sorted slice
	fmt.Println("Sorted by age using sort.Interface:", people)
}
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `005_sorting_a_custom_type_using_sort_interface` directory:

   ```bash
   cd go_sample_examples/026_sorting/005_sorting_a_custom_type_using_sort_interface
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, the output will display the sorted slice of `Person2` structs by age:

```bash
Sorted by age using sort.Interface: [{Charlie 20} {Alice 25} {Bob 30}]
```