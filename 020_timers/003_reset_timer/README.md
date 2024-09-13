Here is the `README.md` for the Go example of resetting a timer:

# Go Sample Example - Resetting a Timer

This repository demonstrates how to reset a timer in Go. It includes an example of using the `Reset` method to modify the duration of an existing timer, allowing it to restart with a new duration even after it has been started.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example shows how to use the <code>Reset</code> method to reset a timer to a new duration in Go.</li>
  <li>If the timer had already fired, it will restart with the new duration and fire after the updated time.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Resetting a Timer
	// The Reset method is used to reset a timer to a new duration.
	// If the timer had already fired, it will restart with the new duration

	// Create a timer that will fire after 2 seconds
	timer := time.NewTimer(2 * time.Second)

	// Start a goroutine to reset the timer after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		timer.Reset(4 * time.Second)
		fmt.Println("Timer reset to 4 seconds.")
	}()

	// Wait for the timer to fire
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

3. Navigate to the `003_reset_timer` directory:

   ```bash
   cd go_sample_examples/020_timers/003_reset_timer
   ```

4. Run the Go program:

   ```bash
   go run 003_reset_timer.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
Timer reset to 4 seconds.
Timer fired!
```