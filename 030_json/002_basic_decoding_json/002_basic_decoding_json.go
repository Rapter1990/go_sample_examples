package main

import (
	"encoding/json"
	"fmt"
)

type Person1 struct {
	Name    string
	Age     int
	Country string
}

func main() {

	// This example demonstrates how to convert a JSON string into a Go struct using json.Unmarshal

	jsonString := `{"Name":"John","Age":30,"Country":"USA"}`

	var person Person1

	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Marshal the struct back to JSON format
	jsonOutput, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(jsonOutput))
}
