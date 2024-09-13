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
