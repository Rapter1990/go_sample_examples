# Go Sample Example - Using `time.After`

This repository demonstrates how to use the `time.After` function in Go. It includes an example of creating a simple timer that waits for a specified duration and fires once, providing a straightforward alternative to `time.NewTimer`.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to use the <code>time.After</code> function in Go to wait for a specific duration.</li>
  <li>The <code>time.After</code> function returns a channel that will receive the current time after the specified duration has passed.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Using time.After
	// time.After is a simpler alternative to time.NewTimer for creating a timer that only needs to fire once.
	// It returns a channel that will receive the current time after the specified duration.

	fmt.Println("Waiting for 2 seconds...")

	// Wait for 2 seconds using time.After
	<-time.After(2 * time.Second)

	fmt.Println("2 seconds passed.")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `004_using_time_after` directory:

   ```bash
   cd go_sample_examples/020_timers/004_using_time_after
   ```

4. Run the Go program:

   ```bash
   go run 004_using_time_after.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
Waiting for 2 seconds...
2 seconds passed.
```