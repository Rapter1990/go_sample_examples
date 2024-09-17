# Go Sample Example - Sorting by Multiple Fields

This repository demonstrates how to sort a slice of structs in Go by multiple fields using the built-in `sort.Slice` function. It provides an example of sorting by one field, and if the values are the same, sorting by a secondary field.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers sorting a slice of structs in Go by multiple fields.</li>
  <li>It sorts a slice of `Person1` structs first by the `Age` field, and if two people have the same age, it sorts them by the `Name` field.</li>
  <li>Sorting by multiple fields is a common task in Go, especially when you need to order records based on multiple criteria.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"sort"
)

// Define the Person1 struct with Name and Age fields
type Person1 struct {
	Name string
	Age  int
}

func main() {

	// Slice of Person1 structs to sort
	people := []Person1{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
		{"Dave", 25},
	}

	// Sorting the slice first by Age, then by Name if Age is the same
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Name < people[j].Name
		}
		return people[i].Age < people[j].Age
	})

	// Printing the sorted slice
	fmt.Println("Sorted by age and name:", people)
}
```

### 🏃 How to Run

1. Ensure you have Go installed. If not, download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `004_sorting_by_multiple_fields` directory:

   ```bash
   cd go_sample_examples/026_sorting/004_sorting_by_multiple_fields
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, the output will display the sorted slice of `Person1` structs by age and name:

```bash
Sorted by age and name: [{Bob 25} {Dave 25} {Alice 30} {Charlie 30}]
```