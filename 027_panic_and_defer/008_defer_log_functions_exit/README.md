# Go Sample Example - Defer Log Functions Exit

This repository demonstrates the use of `defer` in Go to log function exits. The `defer` keyword ensures that a function call is executed after the surrounding function completes, even if it exits unexpectedly due to a panic. This example showcases how `defer` can be used to log the exit of each function.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the usage of `defer` to log function exits in Go.</li>
  <li>The `defer` statement ensures that even when a panic occurs, the `exitLog` function is executed to log the function exit.</li>
  <li>The example demonstrates how `defer` works in normal execution and in the presence of panics.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
)

func main() {
	defer exitLog("main")
	fmt.Println("Main function running")
	performTask()
}

func performTask() {
	defer exitLog("performTask")
	fmt.Println("Task performed")
	panic("Panic in performTask")
}

func exitLog(funcName string) {
	fmt.Printf("Exiting %s\n", funcName)
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/008_defer_log_functions_exit` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/008_defer_log_functions_exit
   ```

4. Run the Go program:

   ```bash
   go run 008_defer_log_functions_exit.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
Main function running
Task performed
Exiting performTask
Exiting main
panic: Panic in performTask

goroutine 1 [running]:
main.performTask()
        /path/to/your-repository/IdeaProjects/go_sample_examples/027_panic_and_defer/008_d
efer_log_functions_exit/008_defer_log_functions_exit.go:16 +0x78
main.main()
        /path/to/your-repository/go_sample_examples/027_panic_and_defer/008_d
efer_log_functions_exit/008_defer_log_functions_exit.go:10 +0x6a

Process finished with the exit code 2

```