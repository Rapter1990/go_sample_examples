# Go Sample Example - Conditional Logic in Templates

This repository demonstrates the use of conditional logic in Go's `text/template` package. It shows how to handle conditions within templates to control the flow of content based on dynamic data, such as struct fields.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic conditional logic in Go templates.</li>
  <li>It includes a demonstration of an `if` block to check the presence of a field and display alternate content if the condition is not met.</li>
  <li>Shows how to render templates dynamically with struct data.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {

	// This template includes a conditional if block to check whether the Age field is present.
	// If it's not, an alternate message is shown

	person := Person{
		Name:    "Alice",
		Age:     25,
		Country: "Canada",
	}

	tpl := `{{if .Age}}My name is {{.Name}} and I am {{.Age}} years old.{{else}}My name is {{.Name}} and I prefer not to disclose my age.{{end}}
I live in {{.Country}}.`

	t := template.Must(template.New("person").Parse(tpl))

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

3. Navigate to the `029_text_samples/001_conditional_logic_in_templates` directory:

```bash
   cd go_sample_examples/029_text_samples/001_conditional_logic_in_templates
```

4. Run the Go program:

```bash
   go run 001_conditional_logic_in_templates.go
```

### 📦 Output

When you run the program, you should see the following output:

```
My name is Alice and I am 25 years old.
I live in Canada.
```

If the `Age` field were missing or set to `0`, the template would display:

```
My name is Alice and I prefer not to disclose my age.
I live in Canada.
```