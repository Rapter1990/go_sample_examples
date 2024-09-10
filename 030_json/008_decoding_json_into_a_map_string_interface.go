package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// This example demonstrates how to decode JSON data into a map[string]interface{}.
	// This is useful for working with JSON data when the structure isn't known in advance

	jsonString := `{"Name":"John","Age":30,"Country":"USA"}`

	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println("Error decoding JSON into map:", err)
		return
	}

	fmt.Println("Decoded map:", result)

	// Accessing fields
	fmt.Println("Name:", result["Name"])
	fmt.Println("Age:", result["Age"])
	fmt.Println("Country:", result["Country"])

}
