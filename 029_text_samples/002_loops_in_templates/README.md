# Go Sample Example - Loops in Templates

This repository demonstrates how to use loops in Go's `text/template` package. It shows how to iterate over a collection of data and render each item in a template, providing dynamic content generation based on input data.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic loop constructs in Go templates using the `range` keyword.</li>
  <li>It includes a demonstration of looping over a slice of structs to display tasks and their statuses.</li>
  <li>Showcases how to render dynamic data in a structured format within templates.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Task struct {
	Title  string
	Status string
}

type Person1 struct {
	Name  string
	Tasks []Task
}

func main() {

	// This template uses range to loop over a slice of Task structs and list out each task with its status.

	person := Person1{
		Name: "Bob",
		Tasks: []Task{
			{"Task 1", "Completed"},
			{"Task 2", "Pending"},
			{"Task 3", "In Progress"},
		},
	}

	tpl := `{{.Name}}'s Tasks:
{{range .Tasks}}- {{.Title}}: {{.Status}}
{{end}}`

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

3. Navigate to the `029_text_samples/002_loops_in_templates` directory:

```bash
   cd go_sample_examples/029_text_samples/002_loops_in_templates
```

4. Run the Go program:

```bash
   go run 002_loops_in_templates.go
```

### 📦 Output

When you run the program, you should see the following output:

```
Bob's Tasks:
- Task 1: Completed
- Task 2: Pending
- Task 3: In Progress
```