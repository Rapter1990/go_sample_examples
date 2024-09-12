package main

import (
	"fmt"
)

func main() {

	// You can chain panic and recover to allow partial recovery and continue execution after a panic

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
			panic("panic again")
		}
	}()
	funcThatPanics()
	fmt.Println("This will not be printed")
}

func funcThatPanics() {
	panic("Original panic")
}
