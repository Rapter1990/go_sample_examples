package main

import (
	"encoding/json"
	"fmt"
)

type Person5 struct {
	Name    string `json:"full_name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {

	// This example demonstrates how to customize JSON field names using struct tags (json:"field_name")

	person := Person5{
		Name:    "John Doe",
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
