# Go Sample Example - Complex Structs and Template Actions

This repository demonstrates the use of complex structs and template actions in Go. It focuses on rendering nested struct data using Go's `text/template` package, showcasing how to display deeply nested fields using the `with` action for a cleaner and more focused template rendering.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the use of structs and nested structs in Go.</li>
  <li>It showcases how to define templates that can render structured data with ease, including how to access nested fields in structs.</li>
  <li>It uses Go’s template package to format and print data, focusing specifically on the `with` action to handle deeply nested fields efficiently.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Contact struct {
	Phone string
	Email string
}

type Person6 struct {
	Name    string
	Age     int
	Contact Contact
}

func main() {

	// The with action is used here to focus the template on a specific field (Contact) of the Person struct,
	// making it easier to work with deeply nested data

	tpl := `Name: {{.Name}}
Age: {{.Age}}
Contact Info:
    {{with .Contact}}
    Phone: {{.Phone}}
    Email: {{.Email}}
    {{end}}`

	t := template.Must(template.New("person").Parse(tpl))

	person := Person6{
		Name: "Eva",
		Age:  32,
		Contact: Contact{
			Phone: "123-456-7890",
			Email: "eva@example.com",
		},
	}

	err := t.Execute(os.Stdout, person)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `006_complex_structs_and_template_actions` directory:

```bash
   cd go_sample_examples/029_text_samples/006_complex_structs_and_template_actions
```

4. Run the Go program:

```bash
   go run 006_complex_structs_and_template_actions.go
```

### 📦 Output

When you run the program, you should see the following output:

```
Name: Eva
Age: 32
Contact Info:
    Phone: 123-456-7890
    Email: eva@example.com
```