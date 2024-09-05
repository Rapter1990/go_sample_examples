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

	// Create a buffered channel with a capacity of 2
	ch := make(chan int, 2)

	// Send two values to the channel without blocking
	ch <- 1
	ch <- 2

	// Print the buffered values
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// If you try to receive more values than what the channel holds, it blocks
	// Uncommenting the line below will cause a deadlock
	// fmt.Println(<-ch)

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Example with Blocking on Buffered Channel

	ch = make(chan int, 2)

	// This will block after the buffer is full
	go func() {
		ch <- 1
		ch <- 2
		fmt.Println("Sending 3 (will block until a value is received)")
		ch <- 3
		fmt.Println("Sent 3")
	}()

	// Receive values from the channel
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

	// Write strings to the buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString("World!")

	// Convert the buffer to a string and print it
	fmt.Println(buffer.String())

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Example with buffer.Write and buffer.WriteByte

	var buffer1 bytes.Buffer

	// Write bytes to the buffer
	buffer1.Write([]byte("This is a buffer. "))
	buffer1.WriteByte('A')
	buffer1.WriteString("dditional text.")

	// Print the buffer content
	fmt.Println(buffer1.String())

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Using bytes.Buffer with fmt.Fprintf
	// fmt.Fprintf writes formatted output to the buffer, which is useful for building complex strings.

	var buffer2 bytes.Buffer

	// Use Fprintf to format and write to the buffer
	fmt.Fprintf(&buffer2, "This is a formatted number: %d", 42)

	// Print the buffer content
	fmt.Println(buffer2.String())

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Custom Buffer Implementation

	buffer3 := &MyBuffer{}

	buffer3.Add(1)
	buffer3.Add(2)
	buffer3.Add(3)

	fmt.Println("Buffer content:")
	fmt.Println(buffer3.Get()) // Output: 1
	fmt.Println(buffer3.Get()) // Output: 2
	fmt.Println(buffer3.Get()) // Output: 3
	fmt.Println(buffer3.Get()) // Output: -1 (empty buffer)

	fmt.Println("-----------------------------------------------------------------------------------")

}
