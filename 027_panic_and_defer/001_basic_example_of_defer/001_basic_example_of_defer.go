package main

import (
	"fmt"
)

func main() {

	// The defer statement is used to delay the execution of a function until the surrounding function returns

	defer fmt.Println("This is deferred")
	fmt.Println("This happens first")
}
