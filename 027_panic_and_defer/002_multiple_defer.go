package main

import (
	"fmt"
)

func main() {

	// multiple defer statements are used, they are executed in Last In, First Out (LIFO) order

	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	fmt.Println("Executing main function")
}
