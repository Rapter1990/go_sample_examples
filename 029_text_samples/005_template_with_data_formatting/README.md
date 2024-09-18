# Go Sample Example - Template with Date Formatting

This repository demonstrates how to use Go's `text/template` package to format data, specifically focusing on formatting `time.Time` values in templates. It showcases how to format dates in a human-readable way using Go's date formatting patterns.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers formatting time and date fields in Go templates.</li>
  <li>It demonstrates how to use Go's date formatting string, specifically the format string `"Monday, January 2, 2006"`, to render dates in a readable format.</li>
  <li>Using this technique, you can control the presentation of time data within templates.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

type Event struct {
	Name string
	Date time.Time
}

func main() {

	// This template formats a time.Time field in a specific way using Go's date formatting.
	// The format string "Monday, January 2, 2006" is a standard Go way of specifying date formats

	tpl := `Event: {{.Name}}
Date: {{.Date.Format "Monday, January 2, 2006"}}`

	t := template.Must(template.New("event").Parse(tpl))

	event := Event{
		Name: "Go Conference",
		Date: time.Now(),
	}

	err := t.Execute(os.Stdout, event)
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

3. Navigate to the `029_text_samples/005_template_with_data_formatting` directory:

```bash
cd go_sample_examples/029_text_samples/005_template_with_data_formatting
```

4. Run the Go program:

```bash
go run 005_template_with_data_formatting.go
```

### 📦 Output

When you run the program, you should see the following output (the date will vary depending on the current time):

```
Event: Go Conference
Date: Wednesday, September 18, 2024
```