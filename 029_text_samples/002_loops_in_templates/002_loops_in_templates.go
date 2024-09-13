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
