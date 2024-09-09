package main

import (
	"fmt"
)

func main() {

	// Deferred functions will still run even if a panic occurs. This is useful for cleanup tasks

	defer fmt.Println("This will run even if panic occurs")
	panic("Panic in main function")
	fmt.Println("This will not be printed")

	/*
		This will run even if panic occurs
		panic: Panic in main function

	*/

}
