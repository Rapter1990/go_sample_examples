# Go Sample Example - Buffering

This repository demonstrates various types of buffering in Go, including buffered channels, bytes.Buffer, and a custom buffer implementation. It showcases how to use buffering mechanisms to manage data in concurrent programming and handle I/O operations efficiently.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the usage of buffered channels in Go, allowing non-blocking operations until the buffer is full.</li>
  <li>It also demonstrates how to use <code>bytes.Buffer</code> for efficient string and byte manipulations.</li>
  <li>A custom buffer implementation using a mutex lock is shown to manage concurrent access safely.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

type MyBuffer struct {
	data []int
	mu   sync.Mutex
}

func (b *MyBuffer) Add(value int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.data = append(b.data, value)
}

func (b *MyBuffer) Get() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.data) == 0 {
		return -1 // Indicating buffer is empty
	}
	value := b.data[0]
	b.data = b.data[1:]
	return value
}

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Buffered Channels

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Example with Blocking on Buffered Channel

	ch = make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		fmt.Println("Sending 3 (will block until a value is received)")
		ch <- 3
		fmt.Println("Sent 3")
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Using select with Buffered Channels

	ch = make(chan int, 2)

	go func() {
		ch <- 1
		time.Sleep(1 * time.Second)
		ch <- 2
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout, no message received")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Buffering with bytes.Buffer

	fmt.Println("-----------------------------------------------------------------------------------")

	var buffer bytes.Buffer

	buffer.WriteString("Hello, ")
	buffer.WriteString("World!")

	fmt.Println(buffer.String())

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Example with buffer.Write and buffer.WriteByte

	var buffer1 bytes.Buffer

	buffer1.Write([]byte("This is a buffer. "))
	buffer1.WriteByte('A')
	buffer1.WriteString("dditional text.")

	fmt.Println(buffer1.String())

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Using bytes.Buffer with fmt.Fprintf

	var buffer2 bytes.Buffer

	fmt.Fprintf(&buffer2, "This is a formatted number: %d", 42)

	fmt.Println(buffer2.String())

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Custom Buffer Implementation

	buffer3 := &MyBuffer{}

	buffer3.Add(1)
	buffer3.Add(2)
	buffer3.Add(3)

	fmt.Println("Buffer content:")
	fmt.Println(buffer3.Get())
	fmt.Println(buffer3.Get())
	fmt.Println(buffer3.Get())
	fmt.Println(buffer3.Get())

	fmt.Println("-----------------------------------------------------------------------------------")

}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `012_buffering` directory:

   ```bash
   cd go_sample_examples/012_buffering
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
-----------------------------------------------------------------------------------
1
2
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Sending 3 (will block until a value is received)
Sent 3
1
2
3

--------------------------------------------------------------------------------
---
Received: 1
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Hello, World!
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
This is a buffer. Additional text.
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
This is a formatted number: 42
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Buffer content:
1
2
3
-1
-----------------------------------------------------------------------------------
```
