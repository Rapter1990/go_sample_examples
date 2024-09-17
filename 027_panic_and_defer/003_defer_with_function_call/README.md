# Go Sample Example - Panic and Defer (Defer with File Operations)

This repository demonstrates how to use `defer` statements in Go for resource cleanup, such as closing files. It shows how to handle file operations safely by ensuring that resources are released even if an error occurs during program execution.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how `defer` is used in Go to close files and clean up resources automatically.</li>
  <li>It includes file creation, writing to the file, and ensuring proper resource cleanup using `defer` to close the file.</li>
  <li>It also demonstrates basic error handling with `panic` in case of file operation failures.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// use defer to call a function that cleans up resources, such as closing a file
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello, Go!")
	if err != nil {
		panic(err)
	}
	fmt.Println("File written successfully")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/003_defer_with_function_call` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/003_defer_with_function_call
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
File written successfully
```

Additionally, a new file `example.txt` will be created with the content:

```txt
Hello, Go!
```