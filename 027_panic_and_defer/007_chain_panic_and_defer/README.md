# Go Sample Example - Chain Panic and Defer

This repository demonstrates the use of chained `panic` and `defer` in Go. It showcases how to recover from a panic, handle the recovery, and then re-trigger a new panic, allowing partial recovery from errors while still stopping execution as needed.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the chaining of `panic` and `recover` with `defer` in Go.</li>
  <li>The `recover` function allows capturing the panic and taking necessary actions (e.g., logging or cleanup).</li>
  <li>In this case, after recovering from the initial panic, a new panic is triggered, demonstrating how recovery can be used to handle errors before re-panicking.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
)

func main() {

	// You can chain panic and recover to allow partial recovery and continue execution after a panic

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
			panic("panic again")
		}
	}()
	funcThatPanics()
	fmt.Println("This will not be printed")
}

func funcThatPanics() {
	panic("Original panic")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/007_chain_panic_and_defer` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/007_chain_panic_and_defer
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
Recovered in main: Original panic
panic: Original panic [recovered]
        panic: panic again

goroutine 1 [running]:
main.main.func1()
        /path/to/your-repository/go_sample_examples/027_panic_and_defer/007_c
hain_panic_and_defer/007_chain_panic_and_defer.go:14 +0x78
panic({0xc82180?, 0xcc0478?})
        /path/to/your-repository/Go/src/runtime/panic.go:785 +0x132
main.funcThatPanics(...)
        /path/to/your-repository/go_sample_examples/027_panic_and_defer/007_c
hain_panic_and_defer/007_chain_panic_and_defer.go:22
main.main()
        /path/to/your-repository/go_sample_examples/027_panic_and_defer/007_c
hain_panic_and_defer/007_chain_panic_and_defer.go:17 +0x3f
```