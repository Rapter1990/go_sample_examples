package main

import (
	"fmt"
)

func main() {

	// The recover function is used to regain control of a panicking goroutine.
	// It can be used within a deferred function to catch a panic and prevent the program from crashing

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Starting main")
	panic("A severe error occurred")
	fmt.Println("This will not be printed")

	/*

		Starting main
		Recovered from panic: A severe error occurred

	*/

}
