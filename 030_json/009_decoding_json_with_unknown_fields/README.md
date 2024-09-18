# Go Sample Example - Decoding JSON with Unknown Fields

This repository demonstrates how to decode JSON into a struct while capturing any additional unknown fields into a `map[string]interface{}`. It showcases how to extend the built-in `UnmarshalJSON` method to handle dynamic or extra fields in the JSON data.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example shows how to decode JSON into a struct while capturing additional, unknown fields that aren't part of the predefined struct.</li>
  <li>It demonstrates overriding the `UnmarshalJSON` method to capture any extra fields in a `map[string]interface{}` for later use.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person9 struct {
	Name    string
	Age     int
	Country string
	Extra   map[string]interface{} `json:"-"` // Extra fields will be stored here
}

func (p *Person9) UnmarshalJSON(data []byte) error {
	type Alias Person9
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Unmarshal the remaining fields into the Extra map
	if err := json.Unmarshal(data, &p.Extra); err != nil {
		return err
	}
	delete(p.Extra, "Name")
	delete(p.Extra, "Age")
	delete(p.Extra, "Country")

	return nil
}

func main() {

	// This example shows how to decode JSON into a struct while capturing any additional unknown fields into a map.
	// The UnmarshalJSON method is overridden to achieve this

	jsonString := `{"Name":"John","Age":30,"Country":"USA","Nickname":"Johnny","Hobby":"Golf"}`

	var person Person9
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Decoded Struct with Extra Fields: %+v\n", person)
	fmt.Printf("Extra Fields: %+v\n", person.Extra)
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `009_decoding_json_with_unknown_fields` directory:

   ```bash
   cd go_sample_examples/030_json/009_decoding_json_with_unknown_fields
   ```

4. Run the Go program:

   ```bash
   go run 009_decoding_json_with_unknown_fields.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
Decoded Struct with Extra Fields: {Name:John Age:30 Country:USA Extra:map[Hobby:
Golf Nickname:Johnny]}
Extra Fields: map[Hobby:Golf Nickname:Johnny]
```