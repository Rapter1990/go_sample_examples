# Go Sample Example - Sorting a Map by Keys

This repository demonstrates how to sort a map by its keys in Go. It showcases how to extract keys from a map, sort them using the `sort.Strings` function, and then iterate over the sorted keys to access the corresponding values.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers sorting a Go map by its keys.</li>
  <li>The map contains key-value pairs of names and ages, which are sorted alphabetically by the key (name).</li>
  <li>By extracting the map's keys, sorting them, and iterating over the sorted list, the values can be printed in order of the keys.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"sort"
)

func main() {

	// Show how to sort a map by its keys

	ages := map[string]int{
		"Alice":   25,
		"Charlie": 20,
		"Bob":     30,
	}

	keys := make([]string, 0, len(ages))
	for k := range ages {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Println("Sorted by keys:")
	for _, k := range keys {
		fmt.Println(k, ages[k])
	}
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `008_sorting_a_map_by_keys` directory:

   ```bash
   cd go_sample_examples/026_sorting/008_sorting_a_map_by_keys
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, the output will display the map entries sorted by their keys:

```bash
Sorted by keys:
Alice 25
Bob 30
Charlie 20
```