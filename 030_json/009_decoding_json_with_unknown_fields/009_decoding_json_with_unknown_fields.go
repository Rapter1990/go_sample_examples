package main

import (
	"encoding/json"
	"fmt"
)

type Person9 struct {
	Name    string
	Age     int
	Country string
	Extra   map[string]interface{} `json:"-"`
}

func (p *Person9) UnmarshalJSON(data []byte) error {
	type Alias Person9
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Unmarshal the remaining fields into the Extra map
	if err := json.Unmarshal(data, &p.Extra); err != nil {
		return err
	}
	delete(p.Extra, "Name")
	delete(p.Extra, "Age")
	delete(p.Extra, "Country")

	return nil
}

func main() {

	// This example shows how to decode JSON into a struct while capturing any additional unknown fields into a map.
	// The UnmarshalJSON method is overridden to achieve this

	jsonString := `{"Name":"John","Age":30,"Country":"USA","Nickname":"Johnny","Hobby":"Golf"}`

	var person Person9
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Decoded Struct with Extra Fields: %+v\n", person)
	fmt.Printf("Extra Fields: %+v\n", person.Extra)
}
