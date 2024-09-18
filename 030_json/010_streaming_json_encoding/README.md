# Go Sample Example - Streaming JSON Encoding

This repository demonstrates how to use streaming JSON encoding in Go to write JSON data directly to an output stream. It showcases encoding a slice of structs into JSON format and optionally formatting the output with indentation for readability.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example shows how to use the `json.Encoder` type to stream JSON encoding directly to an output stream, such as standard output.</li>
  <li>It demonstrates how to set indentation for better readability and encode each struct in a slice individually.</li>
</ul>

## 💻 Code Example

```go
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `010_streaming_json_encoding` directory:

   ```bash
   cd go_sample_examples/030_json/010_streaming_json_encoding
   ```

4. Run the Go program:

   ```bash
   go run 010_streaming_json_encoding.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```json
{
  "Name": "John",
  "Age": 30,
  "Country": "USA"
}
{
  "Name": "Alice",
  "Age": 28,
  "Country": "Canada"
}
{
  "Name": "Bob",
  "Age": 25,
  "Country": "UK"
}
```