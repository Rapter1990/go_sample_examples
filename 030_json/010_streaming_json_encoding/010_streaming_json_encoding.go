package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person10 struct {
	Name    string
	Age     int
	Country string
}

func main() {

	people := []Person10{
		{"John", 30, "USA"},
		{"Alice", 28, "Canada"},
		{"Bob", 25, "UK"},
	}

	// Create a new JSON encoder that writes to standard output
	encoder := json.NewEncoder(os.Stdout)

	// Optionally set the encoder to use indentation for better readability
	encoder.SetIndent("", "  ")

	// Encode each person individually
	for _, person := range people {
		err := encoder.Encode(person)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
	}

}
