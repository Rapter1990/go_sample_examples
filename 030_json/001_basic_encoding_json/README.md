# Go Sample Example - Basic JSON Encoding

This repository demonstrates how to encode Go structs into JSON format using Go's standard library package `encoding/json`. It showcases the basic steps required to marshal Go data structures into JSON format for serialization.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic JSON encoding using Go structs and the `json.Marshal` function.</li>
  <li>It demonstrates how to convert a Go struct into JSON format, handling potential errors during the encoding process.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {

	// This example shows how to convert a Go struct into JSON format using json.Marshal

	person := Person{
		Name:    "John",
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

3. Navigate to the `001_basic_encoding_json` directory:

   ```bash
   cd go_sample_examples/030_json/001_basic_encoding_json
   ```

4. Run the Go program:

   ```bash
   go run 001_basic_encoding_json.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```json
{"Name":"John","Age":30,"Country":"USA"}
```