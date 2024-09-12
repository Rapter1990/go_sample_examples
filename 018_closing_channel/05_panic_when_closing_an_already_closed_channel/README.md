# Go Sample Example - Panic When Closing an Already Closed Channel

This repository demonstrates what happens when you attempt to close a channel that has already been closed in Go. It showcases how to handle such a situation gracefully using the `recover()` function.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example shows the behavior of closing an already closed channel in Go, which results in a panic.</li>
  <li>It demonstrates how to use `recover()` to handle the panic and continue execution without crashing the program.</li>
</ul>

## 💻 Code Example

```go
package main

import "fmt"

func main() {

	// Panic when Closing an Already Closed Channel

	// Closing an already closed channel will cause a panic.
	// In this example, the recover() function is used to handle the panic gracefully

	ch := make(chan int)
	close(ch) // Close the channel

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Attempting to close an already closed channel will cause a panic
	close(ch)

	fmt.Println("This line will not be executed")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `018_closing_channel/05_panic_when_closing_an_already_closed_channel` directory:

```bash
cd go_sample_examples/018_closing_channel/05_panic_when_closing_an_already_closed_channel
```

4. Run the Go program:

```bash
go run 05_panic_when_closing_an_already_closed_channel.go
```

### 📦 Output

When you run the program, you should see the following output:

```bash
Recovered from panic: close of closed channel
```
