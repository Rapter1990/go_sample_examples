# Go Sample Example - Basic JSON Decoding

This repository demonstrates how to decode JSON strings into Go structs using Go's standard library package `encoding/json`. It showcases how to unmarshal JSON data into Go structures and perform operations on the decoded data.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic JSON decoding using Go structs and the `json.Unmarshal` function.</li>
  <li>It demonstrates how to convert a JSON string into a Go struct, handle errors during decoding, and marshal the struct back into a formatted JSON string.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person1 struct {
	Name    string
	Age     int
	Country string
}

func main() {

	// This example demonstrates how to convert a JSON string into a Go struct using json.Unmarshal

	jsonString := `{"Name":"John","Age":30,"Country":"USA"}`

	var person Person1

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

3. Navigate to the `002_basic_decoding_json` directory:

   ```bash
   cd go_sample_examples/030_json/002_basic_decoding_json
   ```

4. Run the Go program:

   ```bash
   go run 002_basic_decoding_json.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```json
{
  "Name": "John",
  "Age": 30,
  "Country": "USA"
}
```