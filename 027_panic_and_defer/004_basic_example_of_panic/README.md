# Go Sample Example - Panic (Basic Example of Panic)

This repository demonstrates a basic use of the `panic` function in Go. It showcases how a `panic` can abruptly stop the flow of the program and unwind the stack, terminating the program execution with an error.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the basic usage of the `panic` function in Go.</li>
  <li>The `panic` function is used to handle unexpected situations or serious errors that make further program execution impossible.</li>
  <li>Once a panic is triggered, the program begins unwinding the stack and prints an error message, terminating the process.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
)

func main() {

	// The panic function is used to stop the ordinary flow of control and start panicking,
	// which in turn begins unwinding the stack

	fmt.Println("Start of main")
	panic("Something went wrong")
	fmt.Println("This will not be printed")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/004_basic_example_of_panic` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/004_basic_example_of_panic
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
Start of main
panic: Something went wrong

goroutine 1 [running]:
main.main()
	/path/to/your-repository/go_sample_examples/027_panic_and_defer/004_basic_example_of_panic/004_basic_example_of_panic.go:13 +0x59
exit status 2
```