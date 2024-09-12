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
