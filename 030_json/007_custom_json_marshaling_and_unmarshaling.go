package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Person7 struct {
	Name string
	Age  int
}

// Custom Marshaling for Person
func (p Person7) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}{
		Name: p.Name,
		Age:  strconv.Itoa(p.Age) + " years old",
	})
}

// Custom Unmarshaling for Person
func (p *Person7) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	p.Name = aux.Name
	age, err := strconv.Atoi(aux.Age)
	if err != nil {
		return err
	}
	p.Age = age
	return nil
}

func main() {

	// This example shows how to implement custom marshaling and unmarshaling methods by overriding MarshalJSON and UnmarshalJSON for the Person struct

	person := Person7{
		Name: "John",
		Age:  30,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Custom Marshaled JSON:", string(jsonData))

	jsonString := `{"name":"Alice","age":"28"}`
	var newPerson Person7

	err = json.Unmarshal([]byte(jsonString), &newPerson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Custom Unmarshaled Struct: %+v\n", newPerson)
}
