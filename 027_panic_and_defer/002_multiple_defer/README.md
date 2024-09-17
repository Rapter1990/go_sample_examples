# Go Sample Example - Panic and Defer (Multiple `defer` Statements)

This repository demonstrates the usage of multiple `defer` statements in Go, showcasing how they are executed in a Last In, First Out (LIFO) order. The `defer` keyword is commonly used to schedule tasks such as resource cleanup that should occur after the surrounding function completes.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the usage of multiple `defer` statements in Go.</li>
  <li>It shows how the `defer` statements are executed in reverse order (LIFO) after the main function completes.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
)

func main() {

	// multiple defer statements are used, they are executed in Last In, First Out (LIFO) order
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")

	fmt.Println("Executing main function")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/002_multiple_defer` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/002_multiple_defer
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
Executing main function
Third defer
Second defer
First defer
```