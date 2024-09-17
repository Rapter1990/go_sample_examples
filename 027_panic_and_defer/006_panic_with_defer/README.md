# Go Sample Example - Panic with Defer

This repository demonstrates the use of `panic` along with `defer` in Go. It showcases how deferred functions are executed even in the event of a panic, making them useful for cleanup tasks or ensuring certain operations are completed.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the usage of `panic` and `defer` in Go.</li>
  <li>The `panic` function triggers a panic that halts normal execution, while the `defer` function ensures that certain code will still run after a panic, typically for cleanup or logging purposes.</li>
  <li>This technique can help ensure that important operations, such as resource deallocation, are still performed when an error occurs.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
)

func main() {

	// Deferred functions will still run even if a panic occurs. This is useful for cleanup tasks.

	defer fmt.Println("This will run even if panic occurs")
	panic("Panic in main function")
	fmt.Println("This will not be printed")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/006_panic_with_defer` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/006_panic_with_defer
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
This will run even if panic occurs
panic: Panic in main function

goroutine 1 [running]:
main.main()
        /path/to/your-repository/go_sample_examples/027_panic_and_defer/006_p
anic_with_defer/006_panic_with_defer.go:12 +0x59
exit status 2
```