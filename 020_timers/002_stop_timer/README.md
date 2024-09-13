# Go Sample Example - Stopping a Timer

This repository demonstrates how to stop a timer in Go before it fires. It includes an example of creating a timer that can be stopped, preventing it from executing the intended action after a specified duration.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example shows how to stop a Go timer using the <code>Stop</code> method before it fires.</li>
  <li>If the timer is successfully stopped, the program will prevent the timer's action from executing, and the message indicating the timer has fired will not be printed.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Stopping a Timer
	// You can stop a timer before it fires using the Stop method.
	// If the timer is stopped before it fires, the program will not receive any value from the timer's channel

	// Create a timer that will fire after 3 seconds
	timer := time.NewTimer(3 * time.Second)

	// Start a goroutine to stop the timer after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		stop := timer.Stop()
		if stop {
			fmt.Println("Timer stopped before it fired.")
		}
	}()

	// Wait for the timer to fire
	<-timer.C
	fmt.Println("This will not be printed if the timer is stopped.")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `001_stopping_timer` directory:

   ```bash
   cd go_sample_examples/020_timers/001_stopping_timer
   ```

4. Run the Go program:

   ```bash
   go run 001_stopping_timer.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```bash
Timer stopped before it fired.
```

If the timer is not stopped in time, the program will print:

```bash
Timer stopped before it fired.
fatal error: all goroutines are asleep - deadlock!
```