# Go Sample Example - Working with JSON Arrays

This repository demonstrates how to work with JSON arrays in Go by decoding a JSON array into a slice of structs. It showcases how to handle and manipulate JSON arrays using Go's `encoding/json` package.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to decode JSON arrays into slices of structs in Go using the `json.Unmarshal` function.</li>
  <li>It demonstrates how to work with JSON data representing multiple objects and then marshal it back to a formatted JSON string.</li>
</ul>

## 💻 Code Example

```go
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `004_working_with_json_arrays` directory:

   ```bash
   cd go_sample_examples/030_json/004_working_with_json_arrays
   ```

4. Run the Go program:

   ```bash
   go run 004_working_with_json_arrays.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```json
[
  {
    "Name": "John",
    "Age": 30,
    "Country": "USA"
  },
  {
    "Name": "Alice",
    "Age": 28,
    "Country": "Canada"
  }
]
```