# Go Sample Example - Template Functions

This repository demonstrates how to use custom template functions in Go's `text/template` package. It shows how to register and use functions within templates to transform and manipulate data dynamically.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers the usage of `FuncMap` in Go templates to register and apply custom functions like `strings.ToUpper` and `strings.ToLower`.</li>
  <li>It demonstrates how to pass data into a template and modify it using built-in functions.</li>
  <li>Focuses on rendering customized template outputs based on the input data and registered functions.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Person2 struct {
	Name    string
	Age     int
	Country string
}

func main() {

	// This example demonstrates how to use custom template functions.
	// The FuncMap is used to register functions like strings.ToUpper and strings.ToLower, which are then used in the template.

	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
	}

	tpl := `Name (upper): {{.Name | ToUpper}}
Name (lower): {{.Name | ToLower}}
Country: {{.Country}}`

	t := template.Must(template.New("person").Funcs(funcMap).Parse(tpl))

	person := Person2{
		Name:    "Charlie",
		Age:     28,
		Country: "UK",
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

3. Navigate to the `029_text_samples/003_template_functions` directory:

```bash
   cd go_sample_examples/029_text_samples/003_template_functions
```

4. Run the Go program:

```bash
   go run 003_template_functions.go
```

### 📦 Output

When you run the program, you should see the following output:

```
Name (upper): CHARLIE
Name (lower): charlie
Country: UK
```