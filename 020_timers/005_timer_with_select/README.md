# Go Sample Example - Timer with `select`

This repository demonstrates how to use timers in Go with the `select` statement to handle multiple concurrent events. It showcases how to manage timers alongside other asynchronous events, providing a useful mechanism to handle concurrent operations and timeouts.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to use timers within a <code>select</code> statement to wait for multiple events simultaneously in Go.</li>
  <li>It includes a demonstration of handling multiple timers, firing based on time, and managing timeouts using <code>time.After</code>.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// Timer with Select ( Use select to handle timers alongside other concurrent events )
	// You can use timers within a select statement to wait for multiple events simultaneously,
	// including the timer firing.

	timer1 := time.NewTimer(2 * time.Second) // create timer in 2 seconds
	timer2 := time.NewTimer(4 * time.Second) // create timer in 4 seconds

	select {
	case <-timer1.C:
		fmt.Println("Timer 1 fired")
	case <-timer2.C:
		fmt.Println("Timer 2 fired")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! No timer fired in 3 seconds.")
	}
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `005_timer_with_select` directory:

   ```bash
   cd go_sample_examples/020_timers/005_timer_with_select
   ```

4. Run the Go program:

   ```bash
   go run 005_timer_with_select.go
   ```

### 📦 Output

When you run the program, you should see one of the following outputs based on the timing of events:

```bash
Timer 1 fired
```