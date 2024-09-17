# Go Sample Example - Panic with Recover

This repository demonstrates the use of `panic` along with `recover` in Go. It showcases how a panicking program can regain control using the `recover` function within a deferred function, preventing the program from crashing.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the usage of `panic` and `recover` in Go.</li>
  <li>The `panic` function triggers a panic, which begins unwinding the stack, while the `recover` function is used to catch the panic and allow the program to continue executing.</li>
  <li>This technique can be used to handle unexpected errors and avoid abrupt termination of the program.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
)

func main() {

	// The recover function is used to regain control of a panicking goroutine.
	// It can be used within a deferred function to catch a panic and prevent the program from crashing

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Starting main")
	panic("A severe error occurred")
	fmt.Println("This will not be printed")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/005_panic_with_recover` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/005_panic_with_recover
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
Starting main
Recovered from panic: A severe error occurred
```