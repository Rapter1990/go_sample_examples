# Go Sample Example - Omitting Empty Fields in JSON

This repository demonstrates how to omit empty fields from JSON output in Go using the `omitempty` struct tag. It showcases how to use Go's `encoding/json` package to exclude fields that have zero values during serialization.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers how to use the `omitempty` tag to exclude empty or zero-value fields from JSON output.</li>
  <li>It demonstrates defining a struct where certain fields are omitted from the JSON output if they have default values (such as an empty string or zero).</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person6 struct {
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
	Country string `json:"country,omitempty"`
}

func main() {

	// This example shows how to omit empty fields from the JSON output using the omitempty struct tag

	person := Person6{
		Name: "John Doe",
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

3. Navigate to the `006_omitting_empty_fields` directory:

   ```bash
   cd go_sample_examples/030_json/006_omitting_empty_fields
   ```

4. Run the Go program:

   ```bash
   go run 006_omitting_empty_fields.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```json
{"name":"John Doe"}
```