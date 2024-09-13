# Go Sample Example - Simple Timer

This example demonstrates the usage of a simple timer in Go. The program shows how to create a timer that fires after a specified duration and performs an action after the timer expires.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the basic usage of Go's <code>time.Timer</code> to schedule an event after a delay.</li>
  <li>The program waits for the timer to fire after 2 seconds, and once the timer fires, it prints a message to the console.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Basic Example: Simple Timer
	// Create a basic timer that fires after 2 seconds

	// Create a timer that will fire after 2 seconds
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for the timer to fire...")

	// Block until the timer's channel sends a value
	<-timer.C

	fmt.Println("Timer fired!")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `001_simple_timer` directory:

   ```bash
   cd go_sample_examples/020_timers/001_simple_timer
   ```

4. Run the Go program:

   ```bash
   go run 001_simple_timer.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
Waiting for the timer to fire...
Timer fired!
```