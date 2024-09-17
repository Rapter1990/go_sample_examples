# Go Sample Example - Panic and Defer

This repository demonstrates the basic usage of `defer` in Go, showcasing how to delay the execution of a function until the surrounding function returns. This is a common feature in Go used for resource cleanup, logging, and more.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers a simple use case of the `defer` statement in Go.</li>
  <li>It includes an example where a function call is deferred, showing how it runs after the surrounding function finishes.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
)

func main() {

	// The defer statement is used to delay the execution of a function until the surrounding function returns
	defer fmt.Println("This is deferred")
	
	fmt.Println("This happens first")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/001_basic_example_of_defer` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/001_basic_example_of_defer
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
This happens first
This is deferred
```