# Go Sample Example - Nested Templates

This repository demonstrates how to define and use nested templates in Go's `text/template` package. It showcases how to break down complex templates into smaller, reusable components.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the usage of nested templates in Go to improve readability and reuse of template code.</li>
  <li>It demonstrates how to define a nested template for a person's address and include it within a larger template that renders a person's details.</li>
  <li>Nested templates can be useful for separating logic and managing complex data structures within templates.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Address struct {
	City    string
	Country string
}

type Person4 struct {
	Name    string
	Age     int
	Address Address
}

func main() {

	// This example shows how to define and use nested templates within a larger template.
	// The address template is defined and then included using the template keyword

	tpl := `{{define "address"}}{{.City}}, {{.Country}}{{end}}

Name: {{.Name}}
Age: {{.Age}}
Address: {{template "address" .Address}}`

	t := template.Must(template.New("person").Parse(tpl))

	person := Person4{
		Name: "Dave",
		Age:  40,
		Address: Address{
			City:    "Berlin",
			Country: "Germany",
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

3. Navigate to the `029_text_samples/004_nested_templates` directory:

```bash
   cd go_sample_examples/029_text_samples/004_nested_templates
```

4. Run the Go program:

```bash
   go run 004_nested_templates.go
```

### 📦 Output

When you run the program, you should see the following output:

```
Name: Dave
Age: 40
Address: Berlin, Germany
```