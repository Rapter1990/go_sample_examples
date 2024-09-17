# Go Sample Example - Closing Resources with Defer

This repository demonstrates how to use `defer` in Go to ensure that resources, such as files, are properly closed after use. It highlights the importance of deferring cleanup actions, such as closing files, to avoid resource leaks, even when an error occurs.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the usage of `defer` to manage resource cleanup in Go.</li>
  <li>The `defer` statement ensures that the file is closed properly after it is no longer needed, even if the program encounters an error.</li>
  <li>This example demonstrates a basic use case of opening and closing a file, with error handling for file operations.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// This example demonstrates using defer to ensure resources like files are closed properly, even if an error occurs

	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("File opened successfully")
	// Processing file content...
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `027_panic_and_defer/009_closing_resources_with_defer` directory:

   ```bash
   cd go_sample_examples/027_panic_and_defer/009_closing_resources_with_defer
   ```

4. Run the Go program:

   ```bash
   go run 009_closing_resources_with_defer.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
File opened successfully
```