package main

import (
	"encoding/json"
	"fmt"
)

type Person6 struct {
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
	Country string `json:"country,omitempty"`
}

func main() {

	// This example shows how to omit empty fields from the JSON output using the omitempty struct tag

	person := Person6{
		Name: "John Doe",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
