package main

import (
	"fmt"
	"os"
)

func main() {

	// This example demonstrates using defer to ensure resources like files are closed properly, even if an error occurs

	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("File opened successfully")
	// Processing file content...
}
