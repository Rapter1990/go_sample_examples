# Go Sample Example - Sorting a Map by Values

This repository demonstrates how to sort a map by its values in Go. It showcases how to convert a map into a slice of key-value pairs, sort the slice by the values, and then iterate over the sorted list to print the results.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers sorting a Go map by its values.</li>
  <li>The map contains key-value pairs of names and ages, which are sorted by their values (ages).</li>
  <li>It demonstrates how to use `sort.Slice` to sort the slice of key-value pairs by their values in ascending order.</li>
</ul>

## 💻 Code Example

```golang
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `009_sorting_a_map_by_values` directory:

   ```bash
   cd go_sample_examples/026_sorting/009_sorting_a_map_by_values
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, the output will display the map entries sorted by their values:

```bash
Sorted by values:
Charlie 20
Alice 25
Bob 30
```