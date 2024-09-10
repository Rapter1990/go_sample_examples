package main

import (
	"encoding/json"
	"fmt"
)

type Person4 struct {
	Name    string
	Age     int
	Country string
}

func main() {

	// This example shows how to decode a JSON array into a slice of structs

	jsonString := `[{"Name":"John","Age":30,"Country":"USA"},{"Name":"Alice","Age":28,"Country":"Canada"}]`

	var people []Person4

	err := json.Unmarshal([]byte(jsonString), &people)
	if err != nil {
		fmt.Println("Error decoding JSON array:", err)
		return
	}

	// Marshal the struct back to JSON format
	jsonOutput, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(jsonOutput))

}
