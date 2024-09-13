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
