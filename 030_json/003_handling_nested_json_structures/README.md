# Go Sample Example - Handling Nested JSON Structures

This repository demonstrates how to handle nested JSON objects in Go using structs within structs. It shows how to unmarshal nested JSON data into Go structures and perform operations on the decoded data.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to decode JSON objects with nested structures using Go structs and the `json.Unmarshal` function.</li>
  <li>It demonstrates how to define structs that contain other structs, unmarshal JSON into these structs, and handle errors during the decoding process.</li>
</ul>

## 💻 Code Example

```go
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `003_handling_nested_json_structures` directory:

   ```bash
   cd go_sample_examples/030_json/003_handling_nested_json_structures
   ```

4. Run the Go program:

   ```bash
   go run 003_handling_nested_json_structures.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```json
{
  "Name": "Alice",
  "Age": 28,
  "Address": {
    "City": "Los Angeles",
    "State": "CA",
    "Country": "USA"
  }
}
```