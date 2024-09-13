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
