# Go Sample Example - Custom JSON Marshaling and Unmarshaling

This repository demonstrates how to implement custom marshaling and unmarshaling for structs in Go. By overriding the `MarshalJSON` and `UnmarshalJSON` methods, you can customize how structs are converted to and from JSON format.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to override default JSON marshaling and unmarshaling for a struct in Go.</li>
  <li>It demonstrates custom formatting for fields during marshaling and parsing specific formats during unmarshaling.</li>
</ul>

## 💻 Code Example

```go
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `007_custom_json_marshaling_and_unmarshaling` directory:

   ```bash
   cd go_sample_examples/030_json/007_custom_json_marshaling_and_unmarshaling
   ```

4. Run the Go program:

   ```bash
   go run 007_custom_json_marshaling_and_unmarshaling.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
Custom Marshaled JSON: {"name":"John","age":"30 years old"}
Custom Unmarshaled Struct: {Name:Alice Age:28}
```