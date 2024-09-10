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
