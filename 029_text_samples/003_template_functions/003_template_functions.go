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
