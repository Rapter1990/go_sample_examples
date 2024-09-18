# Go Sample Example - Custom JSON Field Names

This repository demonstrates how to customize JSON field names in Go using struct tags. It showcases how to use Go's `encoding/json` package to serialize and deserialize struct fields with custom names.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to use struct tags to customize JSON field names during encoding and decoding.</li>
  <li>It demonstrates defining a struct with custom JSON field names using the `json` struct tag.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person5 struct {
	Name    string `json:"full_name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {

	// This example demonstrates how to customize JSON field names using struct tags (json:"field_name")

	person := Person5{
		Name:    "John Doe",
		Age:     30,
		Country: "USA",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `005_custom_json_field_names` directory:

   ```bash
   cd go_sample_examples/030_json/005_custom_json_field_names
   ```

4. Run the Go program:

   ```bash
   go run 005_custom_json_field_names.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```json
{"full_name":"John Doe","age":30,"country":"USA"}
```