package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
}

type Person3 struct {
	Name    string
	Age     int
	Address Address
}

func main() {

	// This example illustrates how to handle nested JSON objects by defining structs within structs.

	jsonString := `{"Name":"Alice","Age":28,"Address":{"City":"Los Angeles","State":"CA","Country":"USA"}}`

	var person Person3

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
