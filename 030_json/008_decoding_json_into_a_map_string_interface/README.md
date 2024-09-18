# Go Sample Example - Decoding JSON into a Map

This repository demonstrates how to decode JSON data into a `map[string]interface{}` in Go. This approach is useful when the structure of the JSON data is not known in advance or when working with dynamic JSON data.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example shows how to handle JSON data when the structure is unknown or dynamic.</li>
  <li>It demonstrates decoding JSON into a `map[string]interface{}` and accessing the fields from the resulting map.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// This example demonstrates how to decode JSON data into a map[string]interface{}.
	// This is useful for working with JSON data when the structure isn't known in advance

	jsonString := `{"Name":"John","Age":30,"Country":"USA"}`

	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println("Error decoding JSON into map:", err)
		return
	}

	fmt.Println("Decoded map:", result)

	// Accessing fields
	fmt.Println("Name:", result["Name"])
	fmt.Println("Age:", result["Age"])
	fmt.Println("Country:", result["Country"])

}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `008_decoding_json_into_a_map_string_interface` directory:

   ```bash
   cd go_sample_examples/030_json/008_decoding_json_into_a_map_string_interface
   ```

4. Run the Go program:

   ```bash
   go run 008_decoding_json_into_a_map_string_interface.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
Decoded map: map[Age:30 Country:USA Name:John]
Name: John
Age: 30
Country: USA
```