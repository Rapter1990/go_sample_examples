package main

import (
	"fmt"
)

func main() {

	// The panic function is used to stop the ordinary flow of control and start panicking,
	// which in turn begins unwinding the stack

	fmt.Println("Start of main")
	panic("Something went wrong")
	fmt.Println("This will not be printed")
}
