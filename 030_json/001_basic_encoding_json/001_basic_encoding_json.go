package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {

	// This example shows how to convert a Go struct into JSON format using json.Marshal

	person := Person{
		Name:    "John",
		Age:     30,
		Country: "USA",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
