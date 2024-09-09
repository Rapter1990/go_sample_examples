package main

import (
	"fmt"
	"os"
)

func main() {

	// use defer to call a function that cleans up resources, such as closing a file

	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello, Go!")
	if err != nil {
		panic(err)
	}
	fmt.Println("File written successfully")
}
