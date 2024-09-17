# Go Sample Example - Concurrent Sorting with Goroutines

This repository demonstrates how to perform concurrent sorting in Go using goroutines. It showcases how to use goroutines and `sync.WaitGroup` to sort multiple slices concurrently.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers concurrent sorting of slices in Go using goroutines.</li>
  <li>It includes sorting multiple slices concurrently using the `sort.Ints` function and a `sync.WaitGroup` to synchronize the goroutines.</li>
  <li>This example demonstrates the power of concurrency in Go for improving performance in multi-threaded environments.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortSlice(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(slice)
	fmt.Println("Sorted slice:", slice)
}

func main() {

	// Use goroutines to sort multiple slices concurrently

	var wg sync.WaitGroup

	slices := [][]int{
		{3, 2, 1},
		{6, 5, 4},
		{9, 8, 7},
	}

	for _, slice := range slices {
		wg.Add(1)
		go sortSlice(slice, &wg)
	}

	wg.Wait()
	fmt.Println("All slices sorted")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `010_concurrent_sorting_with_goroutines` directory:

   ```bash
   cd go_sample_examples/026_sorting/010_concurrent_sorting_with_goroutines
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see output like the following:

```bash
Sorted slice: [7 8 9]
Sorted slice: [1 2 3]
Sorted slice: [4 5 6]
All slices sorted
```