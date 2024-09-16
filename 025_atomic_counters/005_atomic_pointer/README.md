# Go Sample Example - Atomic Pointer

This repository demonstrates how to atomically load and store a pointer value in Go using the `sync/atomic` package. It showcases safe concurrent updates to pointer values, ensuring data integrity across goroutines.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the use of atomic operations to load and store a pointer value in a thread-safe manner.</li>
  <li>It demonstrates the use of `atomic.Value` for storing and retrieving pointer values concurrently between goroutines.</li>
  <li>The example includes updating the value in a separate goroutine and safely retrieving it using atomic operations.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	// Atomic Pointer
	// demonstrate how to atomically load and store a pointer value

	var data atomic.Value

	// Store an initial value
	data.Store("Initial Value")

	// Load and print the current value
	fmt.Println("Current Value:", data.Load())

	// Start a goroutine to change the value
	go func() {
		data.Store("New Value")
	}()

	// Load and print the new value
	fmt.Println("New Value:", data.Load())
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the **025_atomic_counters/005_atomic_pointer** directory:

```bash
cd go_sample_examples/025_atomic_counters/005_atomic_pointer
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see an output similar to:

```bash
Current Value: Initial Value
New Value: New Value
```